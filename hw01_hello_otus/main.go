package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	const hello = "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(hello))
}
