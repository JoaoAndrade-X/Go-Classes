package main

import "fmt"

func main() {

	// C-style loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while-style
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("-------------- infinite loop --------------")

	counter := 0
	for {
		fmt.Println("counter:", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	fmt.Println("-------------- skipping --------------")

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("-------------- array --------------")
	items := [3]string{"Go", "Python", "Java"}
	for _, value := range items {
		fmt.Println(value)
	}
}
