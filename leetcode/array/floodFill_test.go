package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_floodFill(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}
