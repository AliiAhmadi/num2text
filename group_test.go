/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupFunction(t *testing.T) {
	tests := []struct {
		name string
		inp  string
		res  []string
	}{
		{
			name: "",
			inp:  "123456",
			res:  []string{"123", "456"},
		},
		{
			name: "",
			inp:  "321",
			res:  []string{"321"},
		},
		// TODO - Error
		{
			name: "",
			inp:  "-23.456",
			res:  []string{"-23", ".45", "6"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s - test case %d", test.name, i), func(t *testing.T) {
			r := group(test.inp)
			if !reflect.DeepEqual(r, test.res) {
				t.Errorf("group(%s) = %v - expected %v", test.inp, r, test.res)
			}
		})
	}
}
