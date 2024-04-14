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
