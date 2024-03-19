package main

import (
	"os"

	fileutils "frontendmasters.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadTextFile(filePath)
	
	if err != nil {
		println("Error reading file")
	} else {
		newContent	:= c + "Appended text"
		println("File content: ", c)
		fileutils.WriteToFile(filePath + ".output.txt", newContent)
	}
}