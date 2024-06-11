package main

import "fmt"

// Pointers are used to store the memory address of another variable.
// It is used to store the address of a variable.
// It is used to pass the address of a variable to a function.
// It is used to access the address of a variable.
// It is used to access the value of a variable using the address.

func main() {
	// initializing a pointer gives a free 32 bit location address
	// if we dont initialize a pointer, it will give a null pointer exception
	var p *int32 = new(int32)
	var i int32 = 42
	fmt.Printf("Value p points to is : %v\n", *p)
	fmt.Printf("Address of i is : %v\n", &i)
	fmt.Printf("Address of p is : %v\n", p)
	// Set the value at memory location p
	*p = i
	fmt.Printf("Address of p is : %v\n", p)
	fmt.Printf("Value p points to is : %v\n", *p)
	fmt.Printf("Address of i is : %v\n", &i)

	// set the memory location of i to p
	p = &i
	fmt.Printf("Address of p is : %v\n", p)
	fmt.Printf("Value p points to is : %v\n", *p)
	fmt.Printf("Address of i is : %v\n", &i)

	// Note that slice copy uses the same memory location
	// So if we change the value of one slice, the other slice will also change
	// This is because the memory location is the same

	// Passing array by reference
	var thing = [5]float64{1, 2, 3, 4, 5}
	fmt.Println(thing)
	fmt.Println(square(&thing))
	fmt.Println(thing)
}

func square(thing2 *[5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	fmt.Println(thing2)
	return *thing2
}
