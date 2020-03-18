package main

import (
	"fmt"
)

func main() {
	//exercise1()
	//exercise2n3()
	//exercise4()
	//exercise5()
	//exercise6()
	//exercise7()
	exercise8910()
}
func exercise6() {
	var states []string = []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`,
		` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`,
		` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`,
		` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`,
		` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println("The length of the slice is", len(states))
	fmt.Println("The capacity of slice is", cap(states))
	fmt.Println("Index \tName of States ")
	for i := 0; i < len(states); i++ {
		fmt.Println(i, "\t", states[i])
	}
	// fmt.Println("test ")
	// for i, value := range states {
	// 	fmt.Println(i, value)
	// }
}

func exercise7() {
	var name1 []string = []string{"James", "Bond", "Shaken, not stirred"}
	var name2 []string = []string{"Miss", "Moneypenny", "Helloooooo, James."}
	var mdname [][]string = [][]string{name1, name2}
	fmt.Println(mdname)

	for i, ns := range mdname {
		fmt.Println("record \t", i)
		fmt.Println(ns)
		for j, newslice := range ns {
			fmt.Println("index \t", j, newslice)
		}
	}

}
func exercise8910() {
	lastname := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	//adding a record in a map
	lastname["check"] = []string{`one, two`, `three`, `four`}
	//deleting a record from a map
	delete(lastname, "no_dr")
	//fmt.Println(lastname["no_dr"])
	for i, data := range lastname {
		fmt.Println("Lastname is \t", i)
		for j, mdata := range data {
			fmt.Println(" \t", j, mdata)
		}

	}
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
