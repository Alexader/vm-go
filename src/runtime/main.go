package main

import (
	"flag"
	"fmt"
)

// This is a assembler for translating text format assembly into binary instructions
func main() {
	file := flag.String("file", "./file.s", "the path to text assembly")
	flag.Parse()

	binary, err := loadProgram(*file)
	if err != nil {
		fmt.Printf("wrong format of byte code to run: %s\n", err.Error())
		return
	}
	stop := false
	for !stop {
		fmt.Println(binary)
		// fetch instructions and run them in here
		stop = true
	}
}

func loadProgram(filePath string) ([]byte, error) {
	return nil, nil
}
