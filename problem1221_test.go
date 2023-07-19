package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_balancedStringSplit(t *testing.T) {
	assert.Equal(t, 4, balancedStringSplit("RLRRLLRLRL"))
	assert.Equal(t, 2, balancedStringSplit("RLRRRLLRLL"))
	assert.Equal(t, 1, balancedStringSplit("LLLLRRRR"))
}
