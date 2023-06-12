package generate

import (
	"errors"
	"fmt"

	"github.com/atulsingh0/license-go/src/models"
)

func InputValidation(inp *models.RawLicense) error {

	var err error = nil

	if inp.Version == "" {
		err = errors.Join(err, fmt.Errorf("%s", "version should not be empty"))
	}
	if inp.Customer == "" {
		err = errors.Join(err, fmt.Errorf("%s", "customer should not be empty"))
	}
	if inp.ValidFrom == "" {
		err = errors.Join(err, fmt.Errorf("%s", "valid-from should not be empty"))
	}
	if inp.Expiry == "" {
		err = errors.Join(err, fmt.Errorf("%s", "expiry-date should not be empty"))
	}
	if inp.HardExpiry == "" {
		err = errors.Join(err, fmt.Errorf("%s", "hard-expiry-date should not be empty"))
	}
	if inp.Seats == 0 {
		err = errors.Join(err, fmt.Errorf("%s", "seats should not be 0"))
	}
	if inp.HardSeats == 0 {
		err = errors.Join(err, fmt.Errorf("%s", "hard-seats should not be 0"))
	}
	if inp.Type == "" {
		err = errors.Join(err, fmt.Errorf("%s", "type should not be empty"))
	}
	return err
}
