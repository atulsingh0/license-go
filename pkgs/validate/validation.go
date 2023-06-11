package validate

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// TODO: Add logic to read the public key from config

type SignedLicense struct {
	Version    string `json:"version"`
	Customer   string `json:"customer"`
	ValidFrom  string `json:validfrom`
	Expiry     string `json:"expiry-date"`
	HardExpiry string `json:"hard-expiry-date"`
	Seats      int    `json:"seats"`
	HardSeats  int    `json:"hard-seats"`
	Type       string `json:"type"`
	Signature  string `json:"signature"`
}

type License struct {
	Version    string `json:"version"`
	Customer   string `json:"customer"`
	ValidFrom  string `json:validfrom`
	Expiry     string `json:"expiry-date"`
	HardExpiry string `json:"hard-expiry-date"`
	Seats      int    `json:"seats"`
	HardSeats  int    `json:"hard-seats"`
	Type       string `json:"type"`
}

func ParseRawLicense(rawLicense []byte) (SignedLicense, error) {
	var sl SignedLicense = SignedLicense{}
	if err := json.Unmarshal([]byte(rawLicense), &sl); err != nil {
		return sl, fmt.Errorf("[json] - %v", err)
	}
	return sl, nil
}

func (sl SignedLicense) UnsignedLicense() License {
	return License{sl.Version, sl.Expiry, sl.HardExpiry, sl.Seats, sl.HardSeats, sl.Type}
}

func (sl SignedLicense) Validate() error {

	publicKey, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return fmt.Errorf("[public key] - %v", err)
	}

	// Verification of the signature
	if len(sl.Signature) < 88 {
		return fmt.Errorf("invalid signature")
	}

	// Populate unsignedLicense with the license data
	unsignedLicense := sl.UnsignedLicense()

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
