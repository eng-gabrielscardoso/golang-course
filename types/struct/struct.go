package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// Método: função com receiver
func (p product) priceWithDiscount() float64 {
	return p.price - (p.price * p.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "Pencil",
		price:    1.79,
		discount: 0.05,
	}

	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{
		name:     "Laptop",
		price:    1799,
		discount: 0.1,
	}

	fmt.Println(product2, product2.priceWithDiscount())
}
