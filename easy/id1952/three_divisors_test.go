package id1952

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isThree(t *testing.T) {
	assert.Equal(t, false, isThree(1), "1")
	assert.Equal(t, false, isThree(2), "2")
	assert.Equal(t, true, isThree(4), "4")
	assert.Equal(t, true, isThree(9), "9")
	assert.Equal(t, false, isThree(16), "16")
}
