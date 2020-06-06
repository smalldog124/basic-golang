package example1_test

import (
	"basic/unit-test/example1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetName_shold_be_lek(t *testing.T) {
	expected := "lek"

	actual := example1.GetName()

	if actual != expected {
		t.Errorf("fail")
	}
}

func Test_FizzBuzz_input_1_should_be_1(t *testing.T) {
	expected := "1"
	num := 1

	actual := example1.FizzBuzz(num)

	assert.Equal(t, expected, actual)
}
