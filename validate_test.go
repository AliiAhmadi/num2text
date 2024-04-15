/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"testing"
)

// TestExistInArray
func TestExistInArray(t *testing.T) {
	tests := [...]struct {
		name     string
		value    interface{}
		values   []interface{}
		expected bool
	}{
		{
			name:     "int slice - exist",
			value:    2,
			values:   []interface{}{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "string slice - exist",
			value:    "ali",
			values:   []interface{}{"a", "b", "c", "ali"},
			expected: true,
		},
		{
			name:     "uint slice - exist",
			value:    uint(12),
			values:   []interface{}{uint(12), uint(100), uint(129)},
			expected: true,
		},
		{
			name:     "uint8 slice - exist",
			value:    uint8(90),
			values:   []interface{}{uint8(20), uint8(30), uint8(40), uint8(50), uint8(60), uint8(70), uint8(80), uint8(90)},
			expected: true,
		},
		{
			name:     "byte slice - exist",
			value:    byte('w'),
			values:   []interface{}{byte('w'), byte('w'), byte('w'), byte('8'), byte('w'), byte('5'), byte('-'), byte('1'), byte('r'), byte('q')},
			expected: true,
		},
		{
			name:     "string slice - not exist",
			value:    "hello",
			values:   []interface{}{"a", "b", "c", "d"},
			expected: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s - test case %d", test.name, i), func(t *testing.T) {
			ok := existInArray(test.value, test.values)
			if ok != test.expected {
				t.Errorf("existInArray(%v, %v) = %v - expected %v", test.value, test.values, ok, test.expected)
			}
		})
	}
}

// TestNumberValidator
func TestNumberValidator(t *testing.T) {
	tests := [...]struct {
		name     string
		num      string
		expected bool
		err      string
	}{
		{
			name:     "valid number - negetive zero",
			num:      "-0",
			expected: true,
			err:      "",
		},
		{
			name:     "invalid character - without any integer",
			num:      "-",
			expected: false,
			err:      `number can not be [- . ]`,
		},
		{
			name:     "",
			num:      "-0.9",
			expected: true,
			err:      ``,
		},
		{
			name:     "contain more that one -",
			num:      "-123-456",
			expected: false,
			err:      `number can not have more that one '-' or '.'`,
		},
		{
			name:     "contain more .",
			num:      "-.33.3",
			expected: false,
			err:      `number can not have more that one '-' or '.'`,
		},
		{
			name:     "more . and -",
			num:      "-0988.2.-12",
			expected: false,
			err:      `number can not have more that one '-' or '.'`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isv, err := numberValidator(test.num)
			if err != nil {
				if err.Error() != test.err {
					t.Errorf("number_validator error: %v", err)
				}
			} else if test.expected != isv {
				t.Errorf("number_validator(%v) = %v - expected %v", test.num, isv, test.expected)
			}
		})
	}
}
