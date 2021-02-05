package main

import "fmt"

func main() {
	var a = "initial"		// var declares one or more variables
	fmt.Println(a)

	var b, c int = 1, 2		// multiple declarations at once
	fmt.Println(b, c)

	var d = true			// will infer the type of initialized variables
	fmt.Println(d)

	var e int				// variables not initialized are zero valued
	fmt.Println(e)

	f := "apple"			// := syntax is shorthand for declaring and initializing a variable (var f string = "apple")
	fmt.Println(f)
}
