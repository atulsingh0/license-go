package validate

import (
	"testing"
)

func TestValidate(t *testing.T) {
	rawLic := `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg=="}`

	signedLicense, _ := ParseRawLicense([]byte(rawLic))
	err := signedLicense.Validate()

	if err != nil {
		t.Errorf("License is not valid - %v", err)
	}

}

func TestSignLength(t *testing.T) {
	rawLic := `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg==extra"}`

	signedLicense, _ := ParseRawLicense([]byte(rawLic))
	err := signedLicense.Validate()

	if err == nil {
		t.Errorf("License is not valid - %v", err)
	}

}

func TestWrongSign(t *testing.T) {
	rawLic := `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"V5Ai54Np39s20u6yYBFHBTegmxto+UKHp8g+8DaXJ0nGrNcufZivb1xoDEdqTcsSP4d/gEgbBJ4pIau87r0XCQ=="}`

	signedLicense, _ := ParseRawLicense([]byte(rawLic))
	err := signedLicense.Validate()

	if err == nil {
		t.Errorf("License is not valid - %v", err)
	}

}

func TestBadBase64(t *testing.T) {
	rawLic := `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"V5Ai54Np39s20u6yYBFHBTegmxto+UKHp8g+/gEgbBJ4pIau87r0XCQ=="}`

	signedLicense, _ := ParseRawLicense([]byte(rawLic))
	err := signedLicense.Validate()

	if err == nil {
		t.Errorf("License is not valid - %v", err)
	}

}

func TestNullData(t *testing.T) {

	var inputData = []struct {
		testName string
		license  string
	}{
		{testName: "Key version is null", license: `{"expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg=="}`},
		{testName: "Key expiry-date is null", license: `{"version":"v1", "hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg=="}`},
		{testName: "Key hard-expiry-date is null", license: `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg=="}`},
		{testName: "Key seats is null", license: `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","hard-seats":1,"type":"trial","signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg=="}`},
		{testName: "Key hard-seats is null", license: `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"type":"trial","signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg=="}`},
		{testName: "Key type is null", license: `{"version":"v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"signature":"h2wNAs5SW4V9X1bqDGVHRZWBCgUkFcOkeK0dkJgescUBsdfIKpHJfyBcbwkKp1IT+RdQVEjMcqYKrVUi4oztDg=="}`},
	}

	for _, td := range inputData {

		t.Run(td.testName, func(t *testing.T) {
			signedLicense, _ := ParseRawLicense([]byte(td.license))
			err := signedLicense.Validate()
			if err == nil {
				t.Errorf("Test: %s - License is not valid - %v", td.testName, err)
			}
		})
	}
}

func TestParseRawLicense(t *testing.T) {
	// missing { braces
	rawLic := `"version": "v1","expiry-date":"2022-01-01T00:00:00Z","hard-expiry-date":"2022-01-01T00:00:00Z","seats":1,"hard-seats":1,"type":"trial","signature":"V5Ai54Np39s20u6yYBFHBTegmxto+UKHp8g+/gEgbBJ4pIau87r0XCQ=="}`
	_, err := ParseRawLicense([]byte(rawLic))

	if err == nil {
		t.Errorf("Function ParseRawLicense is not working as expected - %v", err)
	}
}
