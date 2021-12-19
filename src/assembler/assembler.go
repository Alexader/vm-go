package main

import (
	"flag"
	"fmt"
)

// This is a assembler for translating text format assembly into binary instructions
func main() {
	file := flag.String("file", "./file.gvm", "the path to text assembly")
	flag.Parse()
	fmt.Println(file)
}
