package random

import (
	"math/rand"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Len(t, String(32), 32)
	r := New()
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), r.String(8, Numeric))
}

func TestRandomString(t *testing.T) {
	r := New()
	// overwrite initialized Source by New()
	rand.Seed(1)

	got := r.String(32, Alphanumeric)
	str := "onrHFkSuohOf8IvDIjS5HUHTGvw8q5mt"

	if got == str {
		t.Errorf("want random string; got: %s\n", got)
	}
}

func BenchmarkRandomString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String(32, Alphanumeric)
	}
}
