package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"full_name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies_list"`
}

func main() {
	x := Person{
		Name:    "Walter White",
		Age:     51,
		Hobbies: []string{"dealing meth", "potted plants"},
	}

	text, _ := json.Marshal(x)

	fmt.Println(string(text))
}
