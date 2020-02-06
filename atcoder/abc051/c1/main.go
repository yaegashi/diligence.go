package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func (con *contest) main() error {
	var sx, sy, tx, ty int
	con.Scan(&sx, &sy, &tx, &ty)
	dx := tx - sx
	dy := ty - sy
	for i := 0; i < dx; i++ {
		con.Print("R")
	}
	for i := 0; i < dy; i++ {
		con.Print("U")
	}
	for i := 0; i < dx; i++ {
		con.Print("L")
	}
	for i := 0; i < dy; i++ {
		con.Print("D")
	}
	con.Print("D")
	for i := 0; i <= dx; i++ {
		con.Print("R")
	}
	for i := 0; i <= dy; i++ {
		con.Print("U")
	}
	con.Print("L")
	con.Print("U")
	for i := 0; i <= dx; i++ {
		con.Print("L")
	}
	for i := 0; i <= dy; i++ {
		con.Print("D")
	}
	con.Print("R")
	con.Println()
	return nil
}
