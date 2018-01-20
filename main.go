package main

import (
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

	printDmarcFeedbacks(dfs)
}
