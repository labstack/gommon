package random

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Len(t, String(32), 32)
	r := New()
	want := r.String(8, Numeric)
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), want)
}

func BenchmarkRandomString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String(32, Alphanumeric)
	}
}
