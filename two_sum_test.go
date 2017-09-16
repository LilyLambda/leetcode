package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	assert.Equal(t, twoSum([]int{}, 0), []int{})
}
