package main

import "fmt"

type item struct {
	itemId int
	qtd    int
	price  float64
}

type order struct {
	userId int
	items  []item
}

func (o order) finalPrice() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.qtd)
	}

	return total
}

func main() {
	order := order{
		userId: 1,
		items: []item{
			item{
				itemId: 1,
				qtd:    2,
				price:  12.10,
			},
			item{
				itemId: 2,
				qtd:    1,
				price:  23.49,
			},
			item{
				itemId: 11,
				qtd:    100,
				price:  3.13,
			},
		},
	}

	fmt.Printf("Total order: U$ %.2f\n", order.finalPrice())
}
