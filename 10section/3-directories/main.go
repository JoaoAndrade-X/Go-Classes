package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Mkdir or MkdirAll
	dir := "Downloads/static/images"
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}

	// Remove or RemoveAll
	if err := os.RemoveAll(filepath.Clean("Downloads")); err != nil {
		log.Fatal(err)
	}
}
