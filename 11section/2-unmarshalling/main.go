package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string  `json:"name" xml:"name"`
	Age      int     `json:"age" xml:"age"`
	Phone    string  `json:"phone" json:"phone"`
	IsActive bool    `json:"active" xml:"active"`
	Role     string  `json:"role" xml:"role"` // don't have it on payload
	Password string  `json:"-" xml:"-"`       // use "-" so it never gets shown
	Profile  profile `json:"profile" xml:"profile"`
}

type profile struct {
	URL string `json:"url"`
}

var payload = `{
"name": "Jane",
"age": 20,
"phone": "123-456-789",
"active": true,
"profile": {
"url": "https://www.jane.co.id"
}
}
`

func main() {

	var u User
	err := json.Unmarshal([]byte(payload), &u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u)

	bs, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
