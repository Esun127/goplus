package main


import "fmt"
import "os"


type point struct {
	x, y int
}


var pf = fmt.Printf

var ps = fmt.Sprintf

func main() {
	p := point{1, 2}
	pf("%v\n", p)

	pf("%+v\n", p)

	pf("%#v\n", p)

	pf("%T\n", p)

	pf("%t\n", true)

	pf("%d\n", 123)

	pf("%b\n", 123)

	pf("%c\n", 123)

	pf("%x\n",123)

	pf("%f\n", 78.9)

	pf("%e\n", 123400000.0)
	pf("%E\n", 123400000.0)

	pf("%s\n", "\"string\"")

	pf("%q\n", "\"string\"")

	pf("%x\n", "hex this")

	pf("%p\n", &p)


	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	pf("|%6s|%6s|\n", "foo", "b")

	pf("|%-6s|%-6s|\n","foo", "b")


	s := ps("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}



