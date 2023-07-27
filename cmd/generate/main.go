package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var (
		key         string // 32 byte key
		jsonLicense string // JSON formatted license
		textLicense string // unformatted license from deal desk
	)

	flag.StringVar(&key, `k`, "", "A key consisting of 32 bytes")
	flag.StringVar(&jsonLicense, `l`, "", "A license to sign")
	flag.StringVar(&textLicense, `t`, "", "raw text containing license information")
	flag.Parse()

	// Verify the key is 32 bytes long
	if len(key) != 32 {
		fmt.Println("Key must be 32 bytes")
		flag.Usage()
		os.Exit(1)
		return
	}

	l, err := getLicenseFromInput(jsonLicense, textLicense)
	if err != nil {
		panic(err)
	}

	licenseAsString, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	var actualLicense signedLicense
	if err := json.Unmarshal([]byte(licenseAsString), &actualLicense); err != nil {
		panic(err)
	}

	privateKey, publicKey := getKeys(key)

	// Output the public key
	fmt.Println("\nPublic Key: ", string(publicKey))

	// Sign a message with the private key
	sig := ed25519.Sign(privateKey, []byte(licenseAsString))

	// Base64 encode the signature
	signature := base64.StdEncoding.EncodeToString(sig)

	actualLicense.Signature = signature

	display(actualLicense)
}

func parseInput(input string) (inputLicenseData, error) {
	var parsedLicense inputLicenseData

	parsedLicense.Version = "v1"

	re := regexp.MustCompile(`(license type \(trial\/ paid\): )(\w+)`)
	regexResult := re.FindStringSubmatch(input)
	//Check if the second regex group was found.  It contains the License Type
	if len(regexResult) < 2 {
		return parsedLicense, errors.New("Unable to find license type")
	}
	if !(regexResult[2] == "paid" || regexResult[2] == "trial" || regexResult[2] == "") {
		return parsedLicense, errors.New("Invalid license type.  Expected 'paid' or 'trial'. found '" + parsedLicense.Type + "'")
	}
	parsedLicense.Type = regexResult[2]

	re = regexp.MustCompile(`(total seats: )(\d+)`)
	regexResult = re.FindStringSubmatch(input)
	//Check if the second regex group was found.  It contains the seat count
	if len(regexResult) < 2 {
		return parsedLicense, errors.New("Unable to find seat number")
	}
	seats, seatErr := strconv.Atoi(regexResult[2])
	if seatErr != nil {
		return parsedLicense, seatErr
	}
	if seats < 1 {
		return parsedLicense, errors.New("Unable to find valid seats")
	}
	parsedLicense.Seats = seats
	parsedLicense.HardSeats = seats

	re = regexp.MustCompile(`(expiration date \(yyyy\-mm\-dd\): )(\d{4}-\d{2}-\d{2})`)
	regexResult = re.FindStringSubmatch(input)
	//Check if the second regex group was found.  It contains the expiry date
	if len(regexResult) < 2 {
		return parsedLicense, errors.New("Unable to find expiry date")
	}
	expiryDate := regexResult[2] + "T23:59:59Z"
	parsedLicense.Expiry = expiryDate
	parsedLicense.HardExpiry = expiryDate

	return parsedLicense, nil
}
