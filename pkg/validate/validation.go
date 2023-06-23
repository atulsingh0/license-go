package validate

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/atulsingh0/license-go/models"
)

type Slic struct {
	*models.SignedLicense
}

type Rlic struct {
	*models.RawLicense
}

// func (sl *Slic) string() ([]byte, error) {
// 	license, err := json.Marshal(sl)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return license, nil
// }

func (sl *Slic) unsignedLicense() *Rlic {

	return &Rlic{&models.RawLicense{
		Customer:   sl.Customer,
		ValidFrom:  sl.ValidFrom,
		Expiry:     sl.Expiry,
		HardExpiry: sl.HardExpiry,
		Seats:      sl.Seats,
		HardSeats:  sl.HardSeats,
		Type:       sl.Type,
	}}
}

func (sl *Slic) Validate() error {

	var pubKey string = os.Getenv("PUB_KEY")

	if len(pubKey) == 0 {
		return fmt.Errorf("PubKey should not be 0 length")
	}

	fmt.Println("---", pubKey)

	publicKey, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return fmt.Errorf("[public key] - %v", err)
	}

	// Verification of the signature
	if len(sl.Signature) < 88 {
		return fmt.Errorf("invalid signature")
	}

	// Populate unsignedLicense with the license data
	unsignedLicense := sl.unsignedLicense()

	// Create a JSON version to validate the signature
	jsonLicense, err := json.Marshal(unsignedLicense)
	if err != nil {
		return fmt.Errorf("[jsonLicense] - %v", err)
	}
	// base64 decode the signature
	signature, err := base64.StdEncoding.DecodeString(sl.Signature)
	if err != nil {
		return fmt.Errorf("[signature] - %v", err)
	}

	// Verify the signature using our public key
	if !ed25519.Verify(ed25519.PublicKey(publicKey), jsonLicense, signature) {
		return fmt.Errorf("[ed25519] Signature verification failed")
	}
	return nil
}
