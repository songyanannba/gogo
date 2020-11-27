package main

import (
	"fmt"
	"io"
	"log"
)

func trace(s string) {
	fmt.Println("entering", s)
}

func untrace(s string) {
	fmt.Println("leaving", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	trace("in b")
	a()
}

func fun1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	//b()

	fun1("GO")

}
