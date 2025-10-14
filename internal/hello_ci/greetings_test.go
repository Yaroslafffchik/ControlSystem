package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	assert.Equal(t, "Hello, Yaroslav", Greet("en", "Yaroslav"))
	assert.Equal(t, "Hola, Yaroslav", Greet("es", "Yaroslav"))
	assert.Equal(t, "Привет, Yaroslav", Greet("ru", "Yaroslav"))
	assert.Equal(t, "Hello, Yaroslav", Greet("unknown", "Yaroslav"))
}
