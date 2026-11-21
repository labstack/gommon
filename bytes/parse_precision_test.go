package bytes

import (
	"math"
	"testing"
)

func TestParseExactByteCounts(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  int64
	}{
		{"1.001KB", 1001},
		{"-1.001KB", -1001},
		{"9007199254740993B", 9007199254740993},
		{"9223372036854775807B", math.MaxInt64},
		{"-9223372036854775807B", -math.MaxInt64},
		{"-9223372036854775808B", math.MinInt64},
		{"0.999999999999999999KiB", 1023},
		{"-0.999999999999999999KiB", -1023},
		{"010B", 10},
		{"12.25KiB", 12544},
		{"1.2345KB", 1234},
		{"-1.2345KB", -1234},
	} {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Parse(tc.input)
			if err != nil || got != tc.want {
				t.Fatalf("Parse(%q) = %d, %v; want %d, nil", tc.input, got, err, tc.want)
			}
		})
	}
}

func TestParseRejectsOutOfRangeByteCounts(t *testing.T) {
	for _, input := range []string{"9223372036854775808B", "-9223372036854775809B", "8EiB", "10EB"} {
		t.Run(input, func(t *testing.T) {
			if got, err := Parse(input); err == nil {
				t.Fatalf("Parse(%q) = %d, nil; want an out-of-range error", input, got)
			}
		})
	}
}
