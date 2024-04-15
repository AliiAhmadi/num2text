/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"testing"
)

func TestNum2TextFunction(t *testing.T) {
	tests := []struct {
		name string
		inp  string
		out  string
	}{
		{
			name: "one digit",
			inp:  "1",
			out:  `یک`,
		},
		{
			name: "one digit",
			inp:  "2",
			out:  "دو",
		},
		{
			name: "one digit",
			inp:  "3",
			out:  "سه",
		},
		{
			name: "one digit",
			inp:  "4",
			out:  "چهار",
		},
		{
			name: "one digit",
			inp:  "5",
			out:  "پنج",
		},
		{
			name: "one digit",
			inp:  "6",
			out:  "شش",
		},
		{
			name: "one digit",
			inp:  "7",
			out:  "هفت",
		},
		{
			name: "one digit",
			inp:  "8",
			out:  "هشت",
		},
		{
			name: "one digit",
			inp:  "9",
			out:  "نه",
		},
		{
			name: "one digit",
			inp:  "10",
			out:  "ده",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s - test case %d", test.name, i), func(t *testing.T) {
			r, err := Convert(test.inp)
			if err != nil {
				t.Fatal(err)
			}

			if r != test.out {
				t.Errorf("Convert(%s) = %s - expected '%s'", test.inp, r, test.out)
			}
		})
	}
}
