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

func numberValidator(num string) (bool, error) {
	// number can not be just a '-', '.' or ''
	iv := []string{"-", ".", ""}
	if existInArray(num, iv) {
		return false, fmt.Errorf("number can not be %+v", iv)
	}

	// check all characters of input number to be in valid chars
	validchs := strings.Split("-.0123456789", "")
	for _, v := range strings.Split(num, "") {
		if !existInArray(v, validchs) {
			return false, fmt.Errorf("%s contain invalid characters", num)
		}
	}

	// number can not have more that one '-' or '.'
	if strings.Count(num, "-") > 1 || strings.Count(num, ".") > 1 {
		return false, fmt.Errorf("number can not have more that one '-' or '.'")
	}

	if strings.Count(num, "-") == 1 && strings.Index(num, "-") != 0 {
		return false, fmt.Errorf("'-' character should be in 0 index")
	}

	return true, nil
}
