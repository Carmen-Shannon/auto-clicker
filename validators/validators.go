package validators

import (
	"errors"
	"strconv"
)

func ValidateIntegerInput(val string) error {
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return errors.New("enter a valid integer greater than 0")
	} else if parsed < 0 {
		return errors.New("negative integers are not valid")
	}

	return nil
}
