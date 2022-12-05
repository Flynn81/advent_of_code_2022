package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumPriorities(t *testing.T) {
	assert.Equal(t, 4, howManyRangesOverlap("../../test-input"))
}
