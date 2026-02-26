package main

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" json:"phone"`
	IsActive bool   `json:"active" xml:"active"`
	Password string `json:"-" xml:"-"` // use "-" so it never gets shown
}

func main() {
	u := User{
		Name:     "John Smith",
		Age:      42,
		IsActive: true,
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(u); err != nil {
		log.Fatal(err)
	}
}
