package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	validate "github.com/atulsingh0/license-go/pkg/validate"
)

func main() {
	var (
		rawLicense string
		filename   string
		export     bool
	)

	flag.StringVar(&rawLicense, `l`, "", "License JSON")
	flag.StringVar(&filename, `f`, "", "Output file (optional; stdout if not specified)")
	flag.BoolVar(&export, `e`, false, "Output with BASH export format")
	flag.Parse()

	// Only enforce the license as a required variable
	// because we can output to stdout if no filename is specified
	if rawLicense == "" {
		flag.Usage()
		log.Fatal("License is required")
	}

	// Convert the raw license into a struct
	// and verify that it is valid
	signedLicense, err := validate.ParseRawLicense([]byte(rawLicense))
	if err != nil {
		log.Fatal(err)
	}

	// Parsing and creating unSigned license from Signed License
	unsignedLicense := signedLicense.UnsignedLicense()

	// Create an object for the frontend license
	// because it has some special fields.
	// The frontend expects this format:
	// {
	// 	"template-id": "enterprise0",
	// 	"type": "{{ .type }}",
	// 	"expiry-date": "{{ .expiration.soft }}",
	// 	"hard-expiry-date": "{{ .expiration.hard }}",
	// 	"seats": {{ .seat_limit.soft }},
	// 	"hard-seats": {{ .seat_limit.hard }}
	// }
	fel := unsignedLicense.GenFrontendLicense()
	jsonFEL, err := json.Marshal(fel)
	if err != nil {
		log.Fatal("[jsonFEL] ", err)
	}

	err = signedLicense.Validate()
	if err == nil {
		fmt.Println("License is Valid")
	} else {
		log.Fatal("[validate Error] - ", err.Error())
	}

	var output string
	if export {
		output = fmt.Sprintf("export CIRCLE_LICENSE_STRING='%s'\n", jsonFEL)
	} else {
		output = string(jsonFEL)
	}

	// Output the license to stdout or to the specified file
	if filename != "" {
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal("[file] ", err)
		}
		defer f.Close()

		// We need to convert the struct to a []byte,
		// so we convert it to a string first
		f.Write([]byte(output))
	} else {
		// We use the "%+v" format to print the struct
		// with the field names
		fmt.Print(output)
	}
}
