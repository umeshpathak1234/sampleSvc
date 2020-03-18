package main

import "fmt"

func main() {
	//var test = []int{100, 200, 300, 400, 500}
	var a int = 10
	var ab int = 20
	var ip *int
	ip = &ab
	fmt.Println(ab)
	fmt.Println(&ab)
	fmt.Println(ip)
	fmt.Println(*ip)
	for a < 20 {
		if a < 15 {
			a++
			continue
		}

		//	fmt.Printf("values of a is %d\n", a)
		a++

	}
	//	miltudimensionArrau()

	//fmt.Println("the average is ", avgerageofaarray(test, 5))
}

func miltudimensionArrau() {

	var a = [5][2]int{{0, 1}, {1, 2}, {3, 4}, {5, 6}, {7, 8}}
	var i, j int
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
func avgerageofaarray(array1 []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += array1[i]
	}
	avg = float32(sum / size)
	return avg

}
