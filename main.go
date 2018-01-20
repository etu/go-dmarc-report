package main

import (
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	for _, file := range files {
		fmt.Printf("Contents of %s:\n", file)
		fmt.Println(string(getZipFileContents(file)))
	}
}
