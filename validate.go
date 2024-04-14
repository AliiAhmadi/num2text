/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
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
	iv := []string{"-", ".", ""}
	if existInArray(num, iv) {
		return false, fmt.Errorf("number can not be %s or %s or %s", iv[0], iv[1], iv[2])
	}

	return true, nil
}
