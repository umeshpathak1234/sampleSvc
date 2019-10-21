package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y float64) float64 {
	return x + y

}

func multiple(a, b string) (string, string) {
	return a, b

}

func main() {
	fmt.Println("the squareroot of 4 is", math.Sqrt(4))
	fmt.Println("a number from 1-100", rand.Intn(100))
	num1, num2 := 5.6, 9.5
	w1, w2 := "Hi", "There"
	fmt.Println(add(num1, num2))
	fmt.Println(multiple(w1, w2))
}
