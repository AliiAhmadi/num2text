/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"strings"
)

// group number to group of 3 digits
func group(num string) []string {
	g := make([]string, 0)
	chs := strings.Split(num, "")

	str := ""
	c := 0
	for _, ch := range chs {
		if c == 3 {
			c = 1
			g = append(g, str)
			str = ""
			str += ch
		} else {
			c++
			str += ch
		}
	}

	if str != "" {
		g = append(g, str)
	}

	return g
}
