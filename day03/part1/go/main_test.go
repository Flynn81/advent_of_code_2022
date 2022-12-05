package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumPriorities(t *testing.T) {

	assert.Equal(t, 16, processRucksack("vJrwpWtwJgWrhcsFMMfFFhFp"))

	assert.Equal(t, 157, sumPriorities("../../test-input"))
}
