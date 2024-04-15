/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"testing"
)

func TestReverseFunction(t *testing.T) {
	tests := [...]struct {
		inp string
		out string
	}{
		{
			inp: "salam",
			out: "malas",
		},
		{
			inp: "12",
			out: "21",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			r := reverse(test.inp)
			if r != test.out {
				t.Errorf("reverse(%s) = '%s' - expected '%s'", test.inp, r, test.out)
			}
		})
	}
}
