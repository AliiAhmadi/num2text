/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"strings"
)

func replaceChars(str string, old []string, new []string) (string, error) {
	if len(old) != len(new) {
		return "", fmt.Errorf("number of elements in slices should be equal")
	}

	for i, v := range new {
		str = strings.Replace(str, old[i], v, -1)
	}

	return str, nil
}
