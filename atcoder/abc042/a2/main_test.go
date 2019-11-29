package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"5 7 5", "YES"},
		{"7 7 5", "NO"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Logf("in  %q", test.in)
			in := bytes.NewBufferString(test.in)
			out := &bytes.Buffer{}
			solve(in, out)
			actual := strings.TrimSpace(out.String())
			expected := test.out
			t.Logf("out %q", actual)
			if actual != expected {
				t.Errorf("got %q, want %q", actual, expected)
			}
		})
	}
}
