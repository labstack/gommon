package bytes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesParseEB(t *testing.T) {
	// EB
	b, err := Parse("8EB")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b)
	}
	b, err = Parse("8E")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b)
	}

	// EB with spaces
	b, err = Parse("8 EB")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b)
	}
	b, err = Parse("8 E")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b)
	}
}
