package stringutils

import (
	"testing"
)

func TestFirst(t *testing.T) {

	tests := []struct {
		in         string
		targetSize int
		padStr     string
		want       string
	}{
		{"", 5, "-", "-----"},
		{"a", 5, "-", "----a"},
	}
	for _, tt := range tests {
		if got := LeftPad(tt.in, tt.targetSize, tt.padStr); got != tt.want {
			t.Errorf(" LeftPad( %s %d %s) = %s ; want = %s", tt.in, tt.targetSize, tt.padStr, got, tt.want)

		}
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		in      string
		wrapper string
		want    string
	}{
		{"ab", "x", "xabx"},
		{"ab", "\"", "\"ab\""},
		{"ab", "'", "'ab'"},
		{"'abcd'", "'", "''abcd''"},
	}

	for _, tt := range tests {
		if got := WrapStr(tt.in, tt.wrapper); got != tt.want {
			t.Errorf("WrapStr ( '%s','%s') = '%s'; want='%s'", tt.in, tt.wrapper, got, tt.want)
		}
	}
}
