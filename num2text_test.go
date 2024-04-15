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
	tests := [...]struct {
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
			name: "two digit",
			inp:  "10",
			out:  "ده",
		},
		{
			name: "two digit",
			inp:  "11",
			out:  "یازده",
		},
		{
			name: "two digit",
			inp:  "12",
			out:  "دوازده",
		},
		{
			name: "two digit",
			inp:  "13",
			out:  "سیزده",
		},
		{
			name: "two digit",
			inp:  "14",
			out:  "چهارده",
		},
		{
			name: "two digit",
			inp:  "15",
			out:  "پانزده",
		},
		{
			name: "two digit",
			inp:  "16",
			out:  "شانزده",
		},
		{
			name: "two digit",
			inp:  "17",
			out:  "هفده",
		},
		{
			name: "two digit",
			inp:  "18",
			out:  "هجده",
		},
		{
			name: "two digit",
			inp:  "19",
			out:  "نوزده",
		},
		{
			name: "two digit",
			inp:  "20",
			out:  "بیست",
		},
		{
			name: "two digit",
			inp:  "21",
			out:  "بیست و یک",
		},
		{
			name: "two digit",
			inp:  "22",
			out:  "بیست و دو",
		},
		{
			name: "two digit",
			inp:  "23",
			out:  "بیست و سه",
		},
		{
			name: "two digit",
			inp:  "24",
			out:  "بیست و چهار",
		},
		{
			name: "one digit - negetive",
			inp:  "-2",
			out:  "منفی دو",
		},
		{
			name: "two digit - negetive",
			inp:  "-23",
			out:  "منفی بیست و سه",
		},
		{
			name: "two digit - negetive",
			inp:  "99",
			out:  "نود و نه",
		},
		{
			name: "",
			inp:  "100",
			out:  "صد",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s - test case %d", test.name, i), func(t *testing.T) {
			r, err := Convert(test.inp)
			if err != nil {
				t.Fatal(err)
			}

			if r != test.out {
				t.Errorf("Convert(%s) = '%s' - expected '%s'", test.inp, r, test.out)
			}
		})
	}
}
