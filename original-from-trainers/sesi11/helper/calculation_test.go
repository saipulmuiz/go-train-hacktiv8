package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	actual := Multiply(10, 20)
	expected := 200
	msg := fmt.Sprintf("failed. expected %d got %d", expected, actual)

	assert.GreaterOrEqual(t, actual, expected)
	assert.Equal(t, expected, actual, msg)

	// if actual != expected {
	// 	msg := fmt.Sprintf("failed. expected %d got %d", expected, actual)
	// 	t.Fatal(msg)
	// }
}

func TestSum0(t *testing.T) {
	actual := Sum()
	expected := 0

	assert.Equal(t, expected, actual)
}

func TestSum1(t *testing.T) {
	actual := Sum(1)
	expected := 1

	if actual != expected {
		t.Fatalf("failed. expected %d got %d", expected, actual)
	}
}

func TestSumMany(t *testing.T) {
	actual := Sum(1, 2, 3, 4, 5)
	expected := 15

	if actual != expected {
		t.Fatalf("failed. expected %d got %d", expected, actual)
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		request  int
		expected bool
	}{
		{
			request:  10,
			expected: false,
		},
		{
			request:  7,
			expected: true,
		},
		{
			request:  13,
			expected: true,
		},
		{
			request:  21,
			expected: false,
		},
		{
			request:  39,
			expected: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.expected, IsPrime(test.request))
		})
	}
}
