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
	x := 15
	d := &x
	fmt.Println(d)
}

// multidimension slice
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := [][]int{{1, 2, 3}, {4, 5, 6}}
// 	for _, element := range a {
// 		for j, e:= range element {
// 		fmt.Println(j, e)
// 	}
// 	}
// }

//{
//syntax of decleration of maps
//var a = make(map[string]int)

//person name(key) : age of the person (value)
//vamsi : 26
//pavan :27
//roshan :25
// a["vamsi"] = 26
// a["pavan"] = 27
// a["roshan"] = 25
// fmt.Println(a)
// fmt.Println(a["pavan"])
// // store the value of key int and value string
// // what is the syntax of map

// var b = make(map[int]string)
// b[26] ="vamsi"
// b[27] = "pavan"
// b[25] = "roshan"
// fmt.Println(b)
// // delete a element from map
// delete(a, "vamsi")
// fmt.Println(a)

// fmt.Println(b[27])
// value, isExist := a["vamsi"]
// fmt.Println(value, isExist)
// fmt.Println("88888888888.  new.  8888888888")
// for index, element := range a{
// fmt.Println(index, element)
// }

// }
