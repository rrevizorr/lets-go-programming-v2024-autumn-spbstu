package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var message string

func main() {
	fmt.Println(message)
}
