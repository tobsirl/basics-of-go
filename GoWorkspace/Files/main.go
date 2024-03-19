package main

import (
	"fmt"
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
		newContent	:= fmt.Sprintf("Original: %v\n Double Original: %v%v", c, c, c)
		println("File content: ", c)
		fileutils.WriteToFile(filePath + ".output.txt", newContent)
	}
}