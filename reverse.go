/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

func reverse(str string) string {
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}
