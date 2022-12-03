package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateRoundValue(t *testing.T) {
	assert.Equal(t, calculateRoundValue(rockValue, paperValue), wonValue+paperValue)
	assert.Equal(t, calculateRoundValue(rockValue, scissorsValue), lostValue+scissorsValue)
	assert.Equal(t, calculateRoundValue(rockValue, rockValue), drawValue+rockValue)
	assert.Equal(t, calculateRoundValue(paperValue, rockValue), lostValue+rockValue)
	assert.Equal(t, calculateRoundValue(paperValue, paperValue), drawValue+paperValue)
	assert.Equal(t, calculateRoundValue(paperValue, scissorsValue), wonValue+scissorsValue)
	assert.Equal(t, calculateRoundValue(scissorsValue, rockValue), wonValue+rockValue)
	assert.Equal(t, calculateRoundValue(scissorsValue, paperValue), lostValue+paperValue)
	assert.Equal(t, calculateRoundValue(scissorsValue, scissorsValue), drawValue+scissorsValue)
}

func TestDecrypt(t *testing.T) {
	assert.Equal(t, decryptElf("A"), rockValue)
	assert.Equal(t, decryptElf("B"), paperValue)
	assert.Equal(t, decryptElf("C"), scissorsValue)
}

func TestCalculateMyScore(t *testing.T) {
	assert.Equal(t, calculateMyScore("../../test-input"), 12)
}
