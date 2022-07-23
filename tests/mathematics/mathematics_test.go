package mathematics

import "testing"

const defaultAssertion = "%v passed"
const defaultError = "Expected %v. Got %v."

func TestMean(t *testing.T) {
	expectedValue := 7.275

	value := Mean(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(defaultError, expectedValue, value)
	}
}
