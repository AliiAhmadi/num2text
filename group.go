/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"slices"
	"strings"
)

// group number to group of 3 digits
func group(num string) []string {
	g := make([]string, 0)
	chs := strings.Split(num, "")

	str := ""
	c := 0

	for i := len(chs) - 1; i >= 0; i-- {
		if c == 3 {
			c = 1
			g = append(g, reverse(str))
			str = ""
			str += chs[i]
		} else {
			c++
			str += chs[i]
		}
	}

	if str != "" {
		g = append(g, reverse(str))
	}
	slices.Reverse(g)

	return g
}
