package generate

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/atulsingh0/license-go/src/models"
	"github.com/google/uuid"
)

type Slic struct {
	*models.SignedLicense
}

type Rlic struct {
	*models.RawLicense
}

func getKeys(key string) (ed25519.PrivateKey, string) {

	// Compute the full 64 byte <private key><public key> from the private key
	priv := ed25519.NewKeyFromSeed([]byte(key))

	// Get the public key and base64 encode it
	pub := base64.StdEncoding.EncodeToString(priv.Public().(ed25519.PublicKey))

	return priv, pub
}

func (rl *Rlic) string() ([]byte, error) {
	licenseAsString, err := json.Marshal(rl)
	if err != nil {
		return nil, err
	}

	return licenseAsString, nil
}

func (sl *Slic) string() ([]byte, error) {
	license, err := json.Marshal(sl)
	if err != nil {
		return nil, err
	}

	return license, nil
}

func (sl *Slic) marshal(licstring []byte) error {
	if err := json.Unmarshal(licstring, sl); err != nil {
		return err
	}
	return nil
}

func (rl *Rlic) Generate() (string, error) {

	var sl = Slic{}

	// generating public and private key based on the random key passed.
	privateKey, _ := getKeys(os.Getenv("KEY"))

	licstring, err := rl.string()
	if err != nil {
		return "", err
	}

	sig := ed25519.Sign(privateKey, licstring)
	signature := base64.StdEncoding.EncodeToString(sig)

	err = sl.marshal(licstring)
	if err != nil {
		fmt.Println("marshaling sl", err)
		return "", err
	}
	sl.Signature = signature

	// generating uuid
	sl.Id = uuid.NewString()

	lic, err := sl.string()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(lic))
	return string(lic), nil
}
