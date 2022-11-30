package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharactersToProcess(t *testing.T) {
	assert.Equal(t, 1, sumSizes("../../test-input"))
}
//
// func TestFindMarkerStart(t *testing.T) {
// 	assert.Equal(t, 5, findMarkerStart("bvwbjplbgvbhsrlpgdmjqwftvncz"))
// 	assert.Equal(t, 6, findMarkerStart("nppdvjthqldpwncqszvftbrmjlhg"))
// 	assert.Equal(t, 10, findMarkerStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
// 	assert.Equal(t, 11, findMarkerStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
// }
