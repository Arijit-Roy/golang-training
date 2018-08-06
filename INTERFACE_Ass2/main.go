package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no file to read")
		os.Exit(1)
	}
	fileName := os.Args[1]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

}
