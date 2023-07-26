package klogzap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInArray(t *testing.T) {
	key1 := ExtraKey("key1")
	key2 := ExtraKey("key2")
	assert.True(t, inArray(key1, []ExtraKey{key1}))
	assert.False(t, inArray(key2, []ExtraKey{key1}))
}
