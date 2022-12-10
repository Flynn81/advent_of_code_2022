package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSizes(t *testing.T) {
	assert.Equal(t, 95437, sumSizes("../../test-input"))
}
