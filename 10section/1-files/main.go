package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePath := "./10section/1-files/text.txt"
	data := "Welcome to the Go programming language! Lots of love for Go"

	// Write file
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done writing file.")

	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	/* Create file as File
	file2, err := os.Create("file-via-create.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	_, err = file2.WriteString("Welcome Kotlin developers")
	if err != nil {
		log.Fatal(err)
	}
	*/

	// Open File
	fileName := "./10section/1-files/file-via-create.txt"
	newFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(newFile)
	lineNum := 1
	for scanner.Scan() {
		fmt.Println(lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	/* Append to file
	anotherFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer anotherFile.Close()

	_, _ = anotherFile.WriteString(fmt.Sprintf("- Kotlin\n"))
	_, _ = anotherFile.WriteString(fmt.Sprintf("- Ruby\n"))
	_, _ = anotherFile.WriteString(fmt.Sprintf("- Ada\n"))
	_, _ = anotherFile.WriteString(fmt.Sprintf("- Rust\n"))
	*/
}
