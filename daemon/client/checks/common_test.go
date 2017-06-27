package checks

import (
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	// Omit standard log output when running tests to allow one to focus on
	// actual test results.
	log.SetOutput(ioutil.Discard)
}

func TestIsVgSizeOk(t *testing.T) {
	tests := []struct {
		line   string
		okSize int
		want   bool
	}{
		{"5.37 26.84 vg_slow", 5, true},
		{"5.37 26.84 vg_slow", 25, false},
	}
	for _, tt := range tests {
		if got := isVgSizeOk(tt.line, tt.okSize); got != tt.want {
			t.Errorf("isVgSizeOk(%q, %v) = %v, want %v", tt.line, tt.okSize, got, tt.want)
		}
	}
}
