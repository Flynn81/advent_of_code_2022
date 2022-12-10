package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizeToDelete(t *testing.T) {
	assert.Equal(t, 24933642, sizeToDelete("../../test-input"))
}
