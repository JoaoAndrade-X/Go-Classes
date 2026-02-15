package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("Modified value: %v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("Val is nil!")
		return
	}

	*val = *val * 10
	fmt.Printf("Modified address: %v\n", val)
}

func main() {
	num := 10

	modifyValue(num)
	fmt.Println(num)

	modifyPointer(&num)
	fmt.Println(num)

	grade := 50
	gradePtr := &grade
	fmt.Printf("gradePtr grade: %+v\n", gradePtr)
	fmt.Printf("gradePtr: %+v\n", &gradePtr)
	fmt.Printf("gradePtr grade: %+v\n", *(&gradePtr))
}
