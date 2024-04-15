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
			name: "group 1",
			inp:  "123456",
			res:  []string{"123", "456"},
		},
		{
			name: "group 2",
			inp:  "321",
			res:  []string{"321"},
		},
		{
			name: "group 3",
			inp:  "-23.456",
			res:  []string{"-23", ".45", "6"},
		},
		{
			name: "group 4",
			inp:  "1234",
			res:  []string{"123", "4"},
		},
		{
			name: "group 5",
			inp:  "000111---...",
			res:  []string{"000", "111", "---", "..."},
		},
		{
			name: "group 6",
			inp:  "0234567895",
			res:  []string{"023", "456", "789", "5"},
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
