package id55

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump(t *testing.T) {
	assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}), "example 1")
	assert.Equal(t, true, canJump([]int{0}), "one element")
}
