package example2_test

import (
	"basic/unit-test/example2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuss(t *testing.T) {
	testCases := map[string]struct {
		input    int
		expected string
	}{
		"Input 1 should be 1":         {1, "1"},
		"Input 2 should be 2":         {2, "2"},
		"Input 3 should be Fizz":      {3, "Fizz"},
		"Input 5 should be Buzz":      {5, "Buzz"},
		"Input 15 should be FizzBuzz": {15, "FizzBuzz"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := example2.FizzBuzz(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
