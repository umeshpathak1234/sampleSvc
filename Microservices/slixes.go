package main

import "fmt"

// func main() {
// 	d := []string{"r", "o", "a", "d"}
// 	var e []string
// 	//e := make([]string, len(d))
// 	// e := d[2:]
// 	// e[0] = "t"
// 	// e[1] = "M"
// 	// d = append(d, "gh")
// 	fmt.Println("e i s", d)
// 	fmt.Println("d is ", e)
// 	copy1 := copy(e, d)

// 	fmt.Println(copy1)
// }

func main() {

	// Creating slices
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	slc2 := []int{1, 2}
	fmt.Println("test")
	fmt.Println(copy(slc1, slc2))
	slc3 := make([]int, 5)

	// Before copying
	fmt.Println("Slice_1:", slc1)
	fmt.Println("Slice_2:", slc2)

	// Copying the slices
	copy1 := copy(slc3, slc1)
	fmt.Println("\nSlice:", slc3)
	fmt.Println("Total number of elements copied:", copy1)

}
