package main

import (
	"io"
	"testing"

	"github.com/yaegashi/contest.go/tester"
)

func TestContest(t *testing.T) {
	tester.Run(t, "*.in.txt", func(in io.Reader, out io.Writer) error {
		con := &contest{in: in, out: out}
		return con.main()
	})
}
