package main

import (
	"archive/zip"
	"io/ioutil"
	"log"
)

func getZipFileContents(fileName string) []byte {
	// Open zip archive for reading
	r, err := zip.OpenReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Open the first file in the archive and return the string
	rc, err := r.File[0].Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	// Read all contents from this file
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}

	return b
}
