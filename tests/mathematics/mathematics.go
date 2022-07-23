package mathematics

func Mean(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	mean := total / float64(len(numbers))

	return mean
}
