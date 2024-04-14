/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

import (
	"fmt"
	"testing"
)

func TestExistInArray(t *testing.T) {
	tests := []struct {
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
