package main

import (
	"embed"
	"fmt"
	"log"
)

// enterprise application in Go
// ---------------------------
var name = "Joseph"

//go:embed hello.txt
var data string

//go:embed public
var public embed.FS

func main() {
	fmt.Println(data)

	// Can also use fs.ReadFile
	d, err := public.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(d))

	fmt.Println(name)
}
