package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" json:"phone"`
	IsActive bool   `json:"active" xml:"active"`
	Password string `json:"-" xml:"-"` // use "-" so it never gets shown
}

var payload = `{"name":"John Smith","age":42, "phone":"", "active":true}`

func main() {
	var u User

	enc := json.NewDecoder(strings.NewReader(payload))
	if err := enc.Decode(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
}
