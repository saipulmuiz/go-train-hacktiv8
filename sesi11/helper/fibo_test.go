package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	tests := []struct {
		requested int
		expected  string
	}{
		{
			requested: 10,
			expected:  "1,1,2,3,5,8",
		},
		{
			requested: 0,
			expected:  "",
		},
		{
			requested: 21,
			expected:  "1,1,2,3,5,8,13,21",
		},

		{
			requested: 42,
			expected:  "1,1,2,3,5,8,13,21,34",
		},
		{
			requested: 55,
			expected:  "1,1,2,3,5,8,13,21,34,55",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			assert.Equal(t, test.expected, Fibonacci(test.requested))
		})
	}
}
