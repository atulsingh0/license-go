package models

type SignedLicense struct {
	Id         string `json:"id"`
	Customer   string `json:"customer"`
	ValidFrom  string `json:"valid-from"`
	Expiry     string `json:"expiry-date"`
	HardExpiry string `json:"hard-expiry-date"`
	Seats      int    `json:"seats"`
	HardSeats  int    `json:"hard-seats"`
	Type       string `json:"type"`
	Signature  string `json:"signature"`
}
