package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, World!")
	// int, int 16, int32, int64
	var intNum int = 32767
	fmt.Println(intNum)

	// float32, float64
	var floatNum float32 = 3.14
	fmt.Println(floatNum)

	var floatNum32 float32 = 3.14
	var intNum32 int32 = 2
	// Cannot use arithmetic operator on different types
	// var result = intNum32 + floaNum32
	var result = float32(intNum32) + floatNum32
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	var result1 = intNum1 / intNum2
	fmt.Println(result1)
	fmt.Println(float32(intNum1) / float32(intNum2))
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello, World!"
	var myString2 string = `Hello, 
	World2!`
	var myString3 string = myString + myString2
	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)

	// len() function does not give number of characters in a string
	// it gives number of bytes (like some special characters might take 2 bytes to store)
	fmt.Println(len(myString))

	// To get number of characters in a string, use utf8.RuneCountInString()
	fmt.Println(utf8.RuneCountInString(myString))

	// Rune is a charcter in Go
	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	// Default value of a int, float, Rune variable is 0
	// Default value of a string variable is ""
	// Default value of a bool variable is false

	// inferred type
	var str = "Hello, World!"
	fmt.Println(str)

	// you can drop the var keyword and use := to declare and initialize a variable
	// Go will infer the type of the variable
	myVar := "Hello, World!"
	fmt.Println(myVar)

	// Initialize multiple variables
	var v1, v2, v3 int = 1, 2, 3
	fmt.Println(v1, v2, v3)

	var1, var2, var3 := 1, 2, 3
	fmt.Println(var1, var2, var3)

	// Constants
	const myConst int = 42
	fmt.Println(myConst)

	printMessage("Drop me in")
	var resultOfDivision, remainder, err = intDivision(10, 3)
	if err != nil {
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Println("Result of division is: ", resultOfDivision)
	} else {
		fmt.Println(resultOfDivision, remainder)
		fmt.Printf("Result: %d, Remainder: %d\n", resultOfDivision, remainder)
	}
	// no break is required in Go
	switch {
	case err != nil:
		fmt.Println(err)
	case remainder == 0:
		fmt.Println("Result of division is: ", resultOfDivision)
	default:
		fmt.Println(resultOfDivision, remainder)
		fmt.Printf("Result: %d, Remainder: %d\n", resultOfDivision, remainder)
	}

	arrayTrials()
	mapTrials()
	stringTrials()
}

func printMessage(printValue string) {
	fmt.Println("Hello, World from printMessage function! " + printValue)
}

// return multiple values -> (num, num) {}
// return single value -> num {}
func intDivision(num1 int, num2 int) (int, int, error) {
	var error error
	if num2 == 0 {
		error = errors.New("Error! Cannot divide by 0")
		return 0, 0, error
	}
	var result int = num1 / num2
	var remainder int = num1 % num2
	return result, remainder, error
}

func arrayTrials() {
	fmt.Println("Array Trials")
	var myArray [3]int32 // = [3]int32{1, 2, 3}
	// Another way to initialize an array
	// myArray = [3]int32{1, 2, 3}
	// Yet another way to initialize an array
	// myArray = [...]int32{1, 2, 3}
	myArray[1] = 102
	fmt.Println(myArray[0])
	fmt.Println(myArray[1:3])
	// The 3 elemts are located in contiguous memory locations
	// int32 takes 4 bytes. So the memory locations are 4 bytes apart
	// To access the memory location of the slice, use &myArray[1]
	fmt.Println(&myArray[0])
	fmt.Println(&myArray[1])
	fmt.Println(&myArray[2])

	// Slices wrap around arrays to give more functionality
	// Slices are reference types
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{5, 6}
	intSlice = append(intSlice, intSlice2...) // ... is used to unpack the slice
	fmt.Println(intSlice)

	//Another way to create a slice (Length is 3 and capacity(optional) is 5)
	var intSlice3 = make([]int32, 3, 5)
	fmt.Println(intSlice3)
	fmt.Println("Looping through array")
	for index, value := range intSlice {
		fmt.Println(index, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// No while loop in Go
	// for {

}

func mapTrials() {
	fmt.Println("Map Trials")
	var myMap map[string]int = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Println(myMap)
	fmt.Println(myMap["one"])

	// Another way to create a map
	myMap2 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(myMap2)

	// Delete a key from a map
	delete(myMap, "one")
	fmt.Println(myMap)

	// Check if a key exists in a map
	value1, ok := myMap["one"]
	if ok {
		fmt.Println("Key exists", value1)
	} else {
		fmt.Println("Key does not exist", value1)
	}

	// Iterate over a map
	for key, value := range myMap2 {
		fmt.Println(key, value)
	}
}

func stringTrials() {
	fmt.Println("String Trials")
	var myString string = "Résumé"
	fmt.Println(myString)
	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))
	var indexedString = myString[0]
	fmt.Printf("%v, %T\n", indexedString, indexedString)
	fmt.Println(string(indexedString))
	for i, v := range myString {
		fmt.Println(i, string(v))
	}
	fmt.Printf("The length of the string is %d\n", len(myString))
	// Note that because of UTF-8 encoding, some characters might take more than 1 byte to store
	// So, indexing a string might not give the expected result
	// note that in above loop, Index 2 is missing in the output.
	// This is because index 1 and 2 are used to store the character "é"
	// ASCII uses 7 bits to represent a character
	// letter "a" is 97th ASCII character is represented as 01100001 in binary
	// To represent an extended set of character like an Emoji or chinese characters, we need more than 7 bits
	// UTF-32 uses 32 bits to represent a character. But it can waste a lot of memory for many characters
	// UTF-8 uses 8 bits to represent a character. It is the most popular encoding scheme
	// it uses variable number of bytes to represent a character

	// The right way to loop through a string is to use the rune type
	var myString2 = []rune("Résumé")
	var indexedString2 = myString2[1]
	fmt.Printf("%v, %T\n", indexedString2, indexedString2)
	for i, v := range myString2 {
		fmt.Println(i, string(v))
	}
	fmt.Printf("The length of the string is %d\n", len(myString2))

	// Strings are immutable in Go
	// To modify a string, convert it to a slice of runes
	// instead of using string concatenation, use strings.Builder
	var stringBuilder strings.Builder
	stringBuilder.WriteString(myString)
	stringBuilder.WriteString(" Builder")
	var catString = stringBuilder.String()
	fmt.Println(catString)
}
