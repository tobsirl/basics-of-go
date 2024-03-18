package main

import (
	"os"

	fileutils "frontendmasters.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	
	if err != nil {
		println("Error reading file")
	} else {
		println("File content: ", c)
	}
}