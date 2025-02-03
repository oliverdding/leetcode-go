package id1991

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMiddleIndex(t *testing.T) {
	assert.EqualValues(t, 3, findMiddleIndex([]int{2, 3, -1, 8, 4}), "default test case")
}

func Test_findMiddleIndex2(t *testing.T) {
	assert.EqualValues(t, -1, findMiddleIndex([]int{-1, 2}), "failed test case")
}
