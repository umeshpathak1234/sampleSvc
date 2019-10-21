package main

import (
	"fmt"
)

func main() {
	exercise1()
	exercise2n3()
	exercise4()
	exercise5()
}
func exercise1() {
	fmt.Println("**************  Exercise 1  **********")
	var b [5]int
	b[0] = 1
	b[1] = 2
	b[2] = 3
	b[3] = 4
	b[4] = 5
	for i, value := range b {
		fmt.Println(i, value)

	}
	//fmt.Println(b[0:5])
	fmt.Printf("%T\n", b)

}
func exercise2n3() {
	fmt.Println("**************  Exercise 2  **********")
	c := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, values := range c {
		fmt.Println(i, values)
	}
	fmt.Println("**************  Exercise 3  **********")
	fmt.Println(c[0:5], "\n", c[5:10], "\n", c[2:7], "\n", c[1:6])

}
func exercise4() {
	fmt.Println("**************  Exercise 4  **********")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
func exercise5() {
	fmt.Println("**************  Exercise 5  **********")
	var x []int = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	var y []int
	y = append(x[:3], x[6:]...)
	fmt.Println(y)
	// var a []int = make([]int, 3)
	// var z []int = make([]int, 4)

	// a = append(x[0:3])
	// z = append(x[6:])
	// fmt.Println(a)
	// fmt.Println(z)
	// fmt.Println(append(a, z...))

}
