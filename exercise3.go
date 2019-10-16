package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

func exerciseThree() {
	fmt.Println("\nExercise 3 starts here ")
	s := fmt.Sprintf("The Integer is %v , String is %v and Boolean is %v.", a, b, c)
	fmt.Println(s)

}
