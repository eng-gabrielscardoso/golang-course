package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	// No Go, quando for exportado, se há a convenção:
	// Identificador maiúsculo: público
	// Identificador minúsculo: privado
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{
		1,
		"Notebook",
		4.500,
		[]string{
			"Promotion",
			"Eletronic",
		},
	}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	var p2 product

	jsonString := `{"id": 2, "name": "Pencil", "tags": ["Paper"]}`

	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tags[0])
}
