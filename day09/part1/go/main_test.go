package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHowManyVisible(t *testing.T) {
	assert.Equal(t, 13, howManyUniqueVisited("../../test-input"))
}
