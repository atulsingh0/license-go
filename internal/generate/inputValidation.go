package generate

import (
	"errors"
	"fmt"
	"time"
)

func (rl *Rlic) InputValidation() error {

	var err error
	if rl == nil {
		return errors.Join(err, fmt.Errorf("request should not be empty"))
	}
	if rl.Customer == "" {
		return errors.Join(err, fmt.Errorf("%s", "customer should not be empty"))
	}
	if rl.ValidFrom == "" {
		return errors.Join(err, fmt.Errorf("%s", "valid-from should not be empty"))
	}
	if rl.Expiry == "" {
		return errors.Join(err, fmt.Errorf("%s", "expiry-date should not be empty"))
	}
	if rl.HardExpiry == "" {
		return errors.Join(err, fmt.Errorf("%s", "hard-expiry-date should not be empty"))
	}
	if rl.Seats == 0 {
		return errors.Join(err, fmt.Errorf("%s", "seats should not be 0"))
	}
	if rl.HardSeats == 0 {
		return errors.Join(err, fmt.Errorf("%s", "hard-seats should not be 0"))
	}

	if rl.Seats > rl.HardSeats {
		return errors.Join(err, fmt.Errorf("%s", "seats should not be greater than hard-seats"))
	}

	if !((rl.Type != "") && ((rl.Type == "prod") || (rl.Type == "dev") || (rl.Type == "lab") || (rl.Type == "test"))) {
		return errors.Join(err, fmt.Errorf("%s", "type should be - lab, dev, test or prod"))
	}

	validFrom, err1 := time.Parse("2006-01-02", rl.ValidFrom)
	if err1 != nil {
		return errors.Join(err, fmt.Errorf("%s", "valid-from should be in YYYY-MM-DD format"))
	}

	Expiry, err1 := time.Parse("2006-01-02", rl.Expiry)
	if err1 != nil {
		return errors.Join(err, fmt.Errorf("%s", "expiry-date should be in YYYY-MM-DD format"))
	}

	HardExpiry, err1 := time.Parse("2006-01-02", rl.HardExpiry)
	if err1 != nil {
		return errors.Join(err, fmt.Errorf("%s", "hard-expiry-date should be in YYYY-MM-DD format"))
	}

	if Expiry.Before(validFrom) {
		return errors.Join(err, fmt.Errorf("%s", "expiry-date should be greater than valid-from date"))
	}

	if HardExpiry.Before(Expiry) {
		return errors.Join(err, fmt.Errorf("%s", "hard-expiry-date should be greater than expiry-date"))
	}

	return nil
}
