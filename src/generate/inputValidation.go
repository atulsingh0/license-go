package generate

import (
	"errors"
	"fmt"
)

func (rl *Rlic) InputValidation() error {

	var err error = nil

	if rl.Version == "" {
		err = errors.Join(err, fmt.Errorf("%s", "version should not be empty"))
	}
	if rl.Customer == "" {
		err = errors.Join(err, fmt.Errorf("%s", "customer should not be empty"))
	}
	if rl.ValidFrom == "" {
		err = errors.Join(err, fmt.Errorf("%s", "valid-from should not be empty"))
	}
	if rl.Expiry == "" {
		err = errors.Join(err, fmt.Errorf("%s", "expiry-date should not be empty"))
	}
	if rl.HardExpiry == "" {
		err = errors.Join(err, fmt.Errorf("%s", "hard-expiry-date should not be empty"))
	}
	if rl.Seats == 0 {
		err = errors.Join(err, fmt.Errorf("%s", "seats should not be 0"))
	}
	if rl.HardSeats == 0 {
		err = errors.Join(err, fmt.Errorf("%s", "hard-seats should not be 0"))
	}
	if rl.Type == "" {
		err = errors.Join(err, fmt.Errorf("%s", "type should not be empty"))
	}
	return err
}
