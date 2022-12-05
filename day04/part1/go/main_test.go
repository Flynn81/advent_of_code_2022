package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumPriorities(t *testing.T) {


	// assert.Equal(t, 38, processRucksack(""))
	// assert.Equal(t, 42, processRucksack(""))
	// assert.Equal(t, 22, processRucksack(""))
	// assert.Equal(t, 20, processRucksack(""))
	// assert.Equal(t, 19, processRucksack(""))

	assert.Equal(t, 2, howManyRangesFullyContain("../../test-input"))
}

// func TestCalculateRoundValue(t *testing.T) {
// 	assert.Equal(t, calculateRoundValue(rockValue, paperValue), wonValue+paperValue)
// 	assert.Equal(t, calculateRoundValue(rockValue, scissorsValue), lostValue+scissorsValue)
// 	assert.Equal(t, calculateRoundValue(rockValue, rockValue), drawValue+rockValue)
// 	assert.Equal(t, calculateRoundValue(paperValue, rockValue), lostValue+rockValue)
// 	assert.Equal(t, calculateRoundValue(paperValue, paperValue), drawValue+paperValue)
// 	assert.Equal(t, calculateRoundValue(paperValue, scissorsValue), wonValue+scissorsValue)
// 	assert.Equal(t, calculateRoundValue(scissorsValue, rockValue), wonValue+rockValue)
// 	assert.Equal(t, calculateRoundValue(scissorsValue, paperValue), lostValue+paperValue)
// 	assert.Equal(t, calculateRoundValue(scissorsValue, scissorsValue), drawValue+scissorsValue)
// }
//
// func TestDecrypt(t *testing.T) {
// 	assert.Equal(t, decryptElf("A"), rockValue)
// 	assert.Equal(t, decryptElf("B"), paperValue)
// 	assert.Equal(t, decryptElf("C"), scissorsValue)
// }
//
// func TestCalculateMyScore(t *testing.T) {
// 	assert.Equal(t, calculateMyScore("../../test-input"), 12)
// }
