package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// check if user gave a filename
	if len(os.Args) < 2 {
		fmt.Println("You didn't type a .txt filename when calling the program!")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)

	// handle potential error on opening file
	if err != nil {
		fmt.Println("Encountered the error – ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
