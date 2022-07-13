package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==:", "Banana" == "Banana")
	fmt.Println("!=:", 3 != 2)
	fmt.Println("<:", 3 < 2)
	fmt.Println(">:", 3 > 2)
	fmt.Println("<=:", 3 <= 2)
	fmt.Println(">=:", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		nome string
	}

	p1 := Pessoa{"Gabriel"}
	p2 := Pessoa{"Fabiana"}

	fmt.Println("Structs:", p1 == p2)
}
