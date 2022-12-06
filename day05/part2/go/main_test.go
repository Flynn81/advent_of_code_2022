package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopOfEachStack(t *testing.T) {
	assert.Equal(t, "MCD", topOfEachStack("../../test-input"))
}
