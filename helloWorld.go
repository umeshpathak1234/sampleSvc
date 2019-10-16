package main

import (
	"fmt"

	"github.com/umeshpathak1234/sampleSvc/api"

	"github.com/google/logger"

	check "github.com/umeshpathak1234/sampleSvc/test"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	fmt.Println("Hello World !")
	logger.Info("I'm about to do something!")
	api.Info()
	check.Testing()
	exerciseOne()
	exerciseTwo()
	exerciseThree()
	//exercise1()
	//exercise2()

	//exercise3()

}

// func exercise1() {
// 	//**********Exercise 1***********//
// 	fmt.Println("\nExercise 1 starts here ")
// 	var x int = 42
// 	var y string = "James Bond"
// 	var z bool = true
// 	fmt.Println("x is", x, "\ny is ", y, "\nz is", z)
// 	fmt.Println("\n", x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }
// func exercise2() {
// 	//**********Exercise 2***********//
// 	fmt.Println("\nExercise 2 starts here ")
// 	fmt.Println("\n", x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }
// func exercise3() {
// 	//**********Exercise 3***********//
// 	fmt.Println("\nExercise 3 starts here ")
// 	s := fmt.Sprintf("The Integer is %v , String is %v and Boolean is %v.", x, y, z)
// 	fmt.Println(s)
// }
