/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"strings"
)

func existInArray[T comparable](value T, values []T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}

func number_validator(num string) (bool, error) {
	// number can not be just a '-', '.' or ''
	iv := []string{"-", ".", ""}
	if existInArray(num, iv) {
		return false, fmt.Errorf("number can not be %s or %s or %s", iv[0], iv[1], iv[2])
	}

	// check all characters of input number to be in valid chars
	validchs := strings.Split("-.0123456789", "")
	for _, v := range strings.Split(num, "") {
		if !existInArray(v, validchs) {
			return false, fmt.Errorf("%s contain invalid characters", num)
		}
	}

	return true, nil
}
