/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"testing"
)

func TestCleanStr(t *testing.T) {
	tests := []struct {
		name string
		inp  string
		exp  string
	}{
		{
			name: "",
			inp:  "\n\n\n",
			exp:  "",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s - test case %d", test.name, i), func(t *testing.T) {
			r := cleanStr(test.inp)
			if r != test.exp {
				t.Errorf("cleanStr(%s) = '%s' - expected '%s'", test.inp, r, test.exp)
			}
		})
	}
}

func TestCleanNumberFunction(t *testing.T) {
	tests := []struct {
		name string
		dec  bool
		inp  string
		exp  string
	}{}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s - test case %d", test.name, i), func(t *testing.T) {
			r := cleanNum(test.inp, test.dec)
			if r != test.exp {
				t.Errorf("cleanNum(%s, %v) = %v - expected %v", test.inp, test.dec, r, test.exp)
			}
		})
	}
}
