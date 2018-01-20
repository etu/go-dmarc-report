package main

import (
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	var dfs []DmarcFeedback

	for _, file := range files {
		dfs = append(
			dfs,
			parseDmarcXML(getZipFileContents(file)),
		)
	}

	fmt.Println("All parsed reports from all files:")
	fmt.Println(dfs)
}
