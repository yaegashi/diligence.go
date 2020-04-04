package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type contest struct {
	in  io.Reader
	out io.Writer
}

func (con *contest) Scan(a ...interface{}) (int, error) {
	return fmt.Fscan(con.in, a...)
}
func (con *contest) Scanln(a ...interface{}) (int, error) {
	return fmt.Fscanln(con.in, a...)
}
func (con *contest) Scanf(f string, a ...interface{}) (int, error) {
	return fmt.Fscanf(con.in, f, a...)
}
func (con *contest) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(con.out, a...)
}
func (con *contest) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(con.out, a...)
}
func (con *contest) Printf(f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(con.out, f, a...)
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	con := &contest{in: in, out: out}
	con.main()
}

type activity struct {
	id, min, start int
}

type byMin []*activity

func (a byMin) Len() int      { return len(a) }
func (a byMin) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byMin) Less(i, j int) bool {
	if a[i].min != a[j].min {
		return a[i].min < a[j].min
	}
	return a[i].start < a[j].start
}

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var N int
		con.Scan(&N)
		activities := make([]*activity, N*2)
		assigns := make([]string, N)
		for i := 0; i < N; i++ {
			var S, E int
			con.Scan(&S, &E)
			activities[i*2+0] = &activity{i, S, 1}
			activities[i*2+1] = &activity{i, E, 0}
		}
		sort.Sort(byMin(activities))
		C := false
		J := false
		I := false
		for _, activity := range activities {
			if activity.start == 1 {
				if !C {
					assigns[activity.id] = "C"
					C = true
					continue
				}
				if !J {
					assigns[activity.id] = "J"
					J = true
					continue
				}
				I = true
				break
			}
			switch assigns[activity.id] {
			case "C":
				C = false
			case "J":
				J = false
			}
		}
		A := "IMPOSSIBLE"
		if !I {
			A = ""
			for _, a := range assigns {
				A += a
			}
		}
		con.Printf("Case #%d: %s\n", testcase, A)
	}
	return nil
}
