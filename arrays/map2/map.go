package main

import "fmt"

func main() {
	empAndSal := map[string]float64{
		"Gabriel Santos Cardoso":  165.000,
		"Fabiana Pantoja Barreto": 100.000,
		"Analua Barreto Cardoso":  1.00,
	}

	empAndSal["Artur Barreto Cardoso"] = 2.00

	delete(empAndSal, "Inexpected")

	for name, salary := range empAndSal {
		fmt.Printf("%v: U$ %.2f\n", name, salary)
	}
}
