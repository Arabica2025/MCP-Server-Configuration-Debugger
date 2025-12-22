// package 만들기
package main

// fmt package: built-in package providing functions for formatted I/O (input/output).
// 1. print
// 2. format strings
import "fmt"

// create main function
func main() {
	//항상 package 언급
	fmt.Println("Hello GO Lang!")

	// 변수 선언 방식
	// var (variable) (data type) = ??
	// (variable) := ??

	/*
		var: variable (mutable)
		const: constant (immutable)
	*/
	var x int = 10
	y := 20 // we can only use := for variables,not for constants(final in Java)
	// also, we cannot use := when we explicitly specify the data type as we declare vars
	const name = "mcp"
	if x > 5 {
		fmt.Println("x is more than 5!")
	} else if y == 20 {
		fmt.Println("x is less than 5 and y is 20!")
	} else {
		fmt.Println("executable?")
	}

	/*
		For loop
		var i int = ?? -> not allowed in golang
		Go's for loop syntax is more concise and does not require
		the var keyword or explicit type declaration within the loop's initialization statement.
	*/
	for i := 1; i < x; i++ {
		fmt.Print(i)
	}

	// int array is mutable
	/*
		In Golang, a map is a built-in data structure that stores a collection of unordered key-value pairs.
		It is an associative array, similar to dictionaries in Python or hash tables in other languages.
		Maps provide an efficient way to store and retrieve data using unique keys to access their corresponding values.
	*/
	var nums = []int{1, 2, 3}
	var conf = map[string]string{"env": "dev"}

	fmt.Println(nums)
	fmt.Println(conf)
	fmt.Printf("integer array nums: %d", nums)
	/*
		first modification for git commit test
	*/
	fmt.Println("first mod for git commit test")
	fmt.Println("any difference?")

	/*
		Data type in Go: need to tell the compiler the data type when declaring vars
		Go can infer data types only when we assign a value
		1. String (most common)
		2. int (most common)


		several types of Integers in Go compared to Java
		Go			Java
		int8		byte
		int16		short
		int32		int
		int64		long
		* uint: (Golang) positive, whole numbers
	*/

	var nameString string
	var ageInt int
	const finalInt uint = 5
	// ask user for their name (input)
	//nameString = "Tom"

	// input Scanner (equivalent to Scanner object in Java)
	// fmt package

	fmt.Scan(&nameString) // <- save user's input value in "nameString" variable <- simply this does not allow taking in user's input
	// we need pointer that we can store input values (pointer to the memory address of the cell that the actual variable is stored)
	// memory address(pointer) expression: &(var name) <- put & sign before var name
	ageInt = 4
	fmt.Printf("User %v is %v years old now\n", nameString, ageInt)
	fmt.Printf("nameString type: %T, ageInt type: %T, finalInt type: %T\n", nameString, ageInt, finalInt)

}
