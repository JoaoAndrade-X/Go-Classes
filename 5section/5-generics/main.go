package main

import "fmt"

type Number interface {
	int | float64 | float32
}

func Sum[T Number](numbers ...T) T {
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	v := Sum(10, 20, 3.3)
	fmt.Printf("%T\n", v)
}
