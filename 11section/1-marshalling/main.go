package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {
	jane := User{
		Name:     "Jane",
		Age:      20,
		Phone:    "123-456-789",
		IsActive: true,
	}

	byteSlice, err := json.Marshal(jane)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteSlice))
}
