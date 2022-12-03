package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
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
  assert.Equal(t, decrypt("A"), rockValue)
  assert.Equal(t, decrypt("X"), rockValue)
  assert.Equal(t, decrypt("B"), paperValue)
  assert.Equal(t, decrypt("Y"), paperValue)
  assert.Equal(t, decrypt("C"), scissorsValue)
  assert.Equal(t, decrypt("Z"), scissorsValue)
}

func TestCalculateMyScore(t *testing.T) {
  assert.Equal(t, calculateMyScore("../../test-input"), 15)
}
