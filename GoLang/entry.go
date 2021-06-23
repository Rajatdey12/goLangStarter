package main

import (
	"fmt"
)

var i int = 6
var j int = 4

const z int = 8

func main() {

	// calling example 1
	example1()
	// calling example 2
	example2()

	// calling example 2a
	example2a()
	// calling example 3
	example3()

	// calling loop example 1
	loopsExample1()

	// Method delegation call
	result := add(i, z)
	println(result)

	//Mthod deletegation with multiple return
	result1, result2 := calculator(i, z)
	println(result1, ":::::::::::::", result2)

	//Method delegation with multiple return in optimized way
	resultop1, resultop2 := calculatorOptimised(i, z)
	print(resultop1, " ::::::::::::: ", resultop2)

}

// CONSTS ANF VARIABLES

func example1() {

	fmt.Println(i + j) //10
}

func example2() {
	// This example will change the value of the variable i from 6 to 8.
	// (because it is not declared here & modifying the package level variable)

	i = 8
	j = 9
	fmt.Println(i + j) //17
}

func example2a() {
	/* -- This example will define & declare its own variable in its local scope and won't
	change the variable in package level -- */

	var i = 6
	println(i + z) //14
}

func example3() {

	k := 0
	l := 5
	fmt.Println(k + l + z) //13
}

// LOOPS

func loopsExample1() {

	for lp := 0; lp < 5; lp++ {
		fmt.Printf("looping values : %d\n", lp)
		// fmt.Println("looping values : ", lp)
	}
}

// METHOD DELEGATION

func add(x, y int) int {
	var out = x + y
	return out
}

// RETURN MULTIPLE VALUES

func calculator(x, y int) (int, int) {
	var sum = x + y
	var minus = x - y
	return sum, minus
}

// OPTIMISED RETURN MULTIPLE VALUES

func calculatorOptimised(x, y int) (sum, minus int) {
	sum = x + y
	minus = x - y
	return
}