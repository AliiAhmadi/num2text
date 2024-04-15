/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(num string) (string, error) {
	// strip whitespaces
	num = strings.TrimSpace(num)

	var err error
	num, err = replaceChars(num,
		[]string{"۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"},
		[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	)

	if err != nil {
		return "", fmt.Errorf("error in conversion %v", num)
	}

	ok, err := numberValidator(num)
	if err != nil {
		return "", fmt.Errorf("error in conversion %v", num)
	}

	if !ok {
		return "", fmt.Errorf("%v contain invalid characters", num)
	}

	isn := (num[0] == '-')
	grp := strings.Split(num, ".")

	var ing string
	if grp[0] != "" {
		ing = grp[0]
	} else {
		ing = "0"
	}

	var dec string
	if len(grp) > 1 {
		dec = grp[1]
	} else {
		dec = "0"
	}

	if isn {
		ing = ing[1:]
	}

	ing = cleanNum(ing, false)
	dec = cleanNum(dec, true)

	if len(dec) > 12 {
		dec = dec[:12]
	}

	tmpDec, err := strconv.Atoi(dec)
	if err != nil {
		return "", fmt.Errorf("error in conversion %v", num)
	}

	tmpIng, err := strconv.Atoi(ing)
	if err != nil {
		return "", fmt.Errorf("error in conversion %v", num)
	}

	if tmpDec+tmpIng == 0 {
		return coll.ones[0], nil
	}

	res := ""

	if isn {
		res += coll.n
	}

	ig := group(ing)
	gcount := len(ig)

	for i, v := range ig {
		v = cleanNum(v, false)
		groupt := ""
		gcount--

		jump := (v == "0")
		if tmp, _ := strconv.Atoi(v); tmp >= 100 {
			// hdigit := cleanNum()
		}
	}

	return "", nil
}
