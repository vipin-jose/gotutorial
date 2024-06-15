package main

import "fmt"

func main() {
	// Generics in Go
	// Generics are a way to write functions and data structures that can work with any type.
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(intSlice))
	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumSlice(float32Slice))
	var float64Slice = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumSlice(float64Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	sum := T(0)
	for _, v := range slice {
		sum += v
	}
	return sum
}

// use of any type
// Any can be used here since it is valid for any type of slice
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumSliceInt(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumSliceFloat32(slice []float32) float32 {
	sum := float32(0)
	for _, v := range slice {
		sum += v
	}
	return sum
}
