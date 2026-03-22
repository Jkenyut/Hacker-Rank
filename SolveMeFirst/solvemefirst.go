// Complete the function  to compute the sum of two integers.
// example
// Input: 2
//        3
// Output: 5
// Function Description

// Complete the  function with the following parameters:

// a: the first value
// b: the second value

// Returns
// - int : the sum of a and b
package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
