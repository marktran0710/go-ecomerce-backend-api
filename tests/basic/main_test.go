package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(input)

	// if actual != output {
	// 	t.Errorf("AddOne(%d), input %d, actual =%d", input, output, actual)
	// }

	assert.Equal(t, actual, output, "AddOne(2) should be 3")
}
