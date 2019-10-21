package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("****************** EXERCISE 1 *****************")
	exercise1()

	fmt.Println("****************** EXERCISE 3 *****************")
	fmt.Println("for exercise 3")
	exercise3()

	fmt.Println("****************** EXERCISE 4 *****************")
	fmt.Println("for exercise 4")
	exercise4()
	fmt.Println("****************** EXERCISE 5 *****************")
	fmt.Println("for exercise 5")
	exercise5()
	fmt.Println("****************** EXERCISE 6 AND 7 *****************")
	fmt.Println("for exercise 6 And 7")
	exercise6n7()
	fmt.Println("****************** EXERCISE 8 *****************")
	fmt.Println("for exercise 8")
	exercise8()
	fmt.Println("****************** EXERCISE 9 *****************")
	fmt.Println("for exercise 9")
	exercise9()
	fmt.Println("****************** EXERCISE 10 *****************")
	fmt.Println("for exercise 10")
	exercise10()

}
func exercise1() {

	var m int = 1
	for i := 1; i <= 1000; i++ {
		fmt.Print(i, "\t")

		for m%10 == 0 {
			fmt.Println("")
			break
		}
		m++
	}
	fmt.Println("")
}
func exercise3() {
	var k int = 1
	for i := 2000; i <= time.Now().Local().Year(); i++ {
		fmt.Print(i, "\t")
		for k%5 == 0 {
			fmt.Println("")
			break
		}
		k++
	}
	fmt.Println("")
}
func exercise4() {
	var l int = 1
	var j int = 2000
	for {
		if j == (time.Now().Year() + 1) {
			break
		}
		fmt.Print(j, "\t")
		j++
		for l%5 == 0 {
			fmt.Println("")
			break
		}
		l++
	}
	fmt.Println("")

}
func exercise5() {

	var j int = 0
	for i := 10; i <= 100; i++ {

		fmt.Print(i%4, "\t")

		j++
		for j%5 == 0 {
			fmt.Println("")
			break
		}
	}
	fmt.Println("")
}
func exercise6n7() {
	s := "Golangr"
	if len(s) == 6 {
		fmt.Println("The length of the string is 6")
	} else if string(s[2]) == "l" {
		fmt.Println("The first alphabet is G")
	} else {
		fmt.Println("this is not golang word")
	}
	fmt.Println("")
}
func exercise8() {
	switch {
	case true:
		fmt.Println("True Statement")
	case false:
		fmt.Println("False Statement")
	}
	fmt.Println("")
}
func exercise9() {
	var favSport string = "football"
	switch favSport {
	case "football":
		fmt.Println("you are right")
	case "baseball":
		fmt.Println("Not that much")
	}
	fmt.Println("")
}
func exercise10() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
