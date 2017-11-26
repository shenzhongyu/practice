package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {

	os.Setenv("Foo", "1")
	fmt.Println("Foo", os.Getenv("Foo"))
	fmt.Println("BAR", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
