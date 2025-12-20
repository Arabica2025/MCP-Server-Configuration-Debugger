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
	y := 20
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
}
