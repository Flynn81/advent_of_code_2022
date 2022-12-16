package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHowManyVisible(t *testing.T) {
	assert.Equal(t, 8, howManyVisible("../../test-input"))
	//assert.Equal(t, 24, howManyVisible("../../test-input2"))
}
