/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"regexp"
	"strings"
)

var (
	cleanStrRX = regexp.MustCompile(`\s+`)
)

func cleanStr(num string) string {
	num = strings.TrimSpace(num)
	num = cleanStrRX.ReplaceAllString(num, " ")

	return num
}

func cleanNum(num string, dec bool) string {
	if dec {
		num = strings.TrimRight(num, "0")
	} else {
		num = strings.TrimLeft(num, "0")
	}

	if num == "" {
		num = "0"
	}

	return num
}

//
