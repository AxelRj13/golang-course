package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Open(os.Args[1]) // os.Args[1] => get argument on index 1
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file) // copy file content to the terminal
	// the File type already implement Reader interface
}
