/*
  [TEST COMMENTS]
*/

package main

import "fmt"

func main() {
	const name = "Sudharsan Ravikumar"
	const age = 24

	fmt.Printf("%s is %d years old \n", name, age)
	fmt.Printf("%s is %d years old. \t and the type is %T and %T\n", name, age, name, age)
}
