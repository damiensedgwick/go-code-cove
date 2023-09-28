package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 2, Sum(1, 1))
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, 0, Subtract(1, 1))
}
