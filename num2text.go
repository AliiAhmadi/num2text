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
			hdigit := cleanNum(strings.Split(strconv.Itoa(tmp/100), ".")[0], false)
			hd, _ := strconv.Atoi(hdigit)
			groupt += coll.hundreds[hd]

			v = fmt.Sprintf("%v", tmp%100)
			if v != "0" {
				groupt += coll.v
			}
		}

		if tmp, _ := strconv.Atoi(v); tmp >= 20 || tmp == 10 {
			tndigit := cleanNum(strings.Split(fmt.Sprintf("%v", tmp/10), ".")[0], false)
			td, _ := strconv.Atoi(tndigit)
			groupt += coll.tens[td]

			v = fmt.Sprintf("%v", tmp%10)
			if v != "0" {
				groupt += coll.v
			}
		}

		if tmp, _ := strconv.Atoi(v); tmp >= 1 && tmp <= 19 {
			groupt += coll.ones[tmp]
		}

		res += groupt

		if !jump {
			res += " " + coll.thousands[gcount]
			if gcount != 0 {
				tmpgs := ig[:i]
				for j, g := range tmpgs {
					ff, err := strconv.Atoi(g)
					if err != nil || ff == 0 {
						tmpgs = append(tmpgs[:j], tmpgs[j+1:]...)
					}
				}
				if len(tmpgs) != 1 {
					res += coll.v
				}
			}
		}
	}

	if dec != "0" {
		if ing == "0" {
			res += " "
		} else {
			res += coll.v
		}

		decg := group(dec)
		dcount := len(decg)
		dgc := dcount

		for _, gr := range decg {
			gr = cleanNum(gr, false)

			groupt := ""
			dcount--

			jump := (gr == "0")

			if tmp, _ := strconv.Atoi(gr); tmp >= 100 {
				hdigits := cleanNum(strings.Split(strconv.Itoa(tmp/100), ".")[0], false)
				ff, _ := strconv.Atoi(hdigits)
				groupt += coll.hundreds[ff]
				gr = fmt.Sprintf("%v", tmp%100)
				if gr != "0" {
					groupt += coll.v
				}
			}

			if tmp, _ := strconv.Atoi(gr); tmp >= 20 || tmp == 10 {
				tndigits := cleanNum(strings.Split(strconv.Itoa(tmp/10), ".")[0], false)
				td, _ := strconv.Atoi(tndigits)
				groupt += coll.tens[td]
				gr = fmt.Sprintf("%v", tmp%10)
				if gr != "0" {
					groupt += coll.v
				}
			}

			if tmp, _ := strconv.Atoi(gr); tmp >= 1 && tmp <= 19 {
				groupt += coll.ones[tmp]
			}

			res += groupt

			if !jump {
				res += " " + coll.thousands[gcount]
				if gcount != 0 {
					res += coll.v
				}
			}
		}

		z := ""
		if dgc == 1 {
			z = ""
		}

		dgl := len(decg[0])
		if dgl == 1 {
			res += coll.tens[1] + z
		} else if dgl == 2 {
			res += coll.hundreds[1] + z
		} else if dgl == 3 {
			dgc++
		}

		res += coll.thousands[dgc-1] + coll.m
	}

	return cleanStr(res), nil
}
