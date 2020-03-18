package main

import (
	"fmt"
	"runtime"
	"sync"
)

type nameSystem struct {
	lname, nickname string
}

var w sync.WaitGroup
var a = make(map[string]string)

var m = map[string]nameSystem{
	"rabi": nameSystem{"basnet", "teja"},
	"bro":  nameSystem{"lalitpur", "houston"},
}

func foo() {
	fmt.Println("This is foo")
	w.Done()
}
func bar() {
	fmt.Println("This is bar")
}
func main() {
	fmt.Println("Number of cpus", runtime.NumCPU())
	fmt.Println("Number of go routine", runtime.NumGoroutine())
	w.Add(1)
	go foo()
	bar()
	fmt.Println("Number of cpus", runtime.NumCPU())
	fmt.Println("Number of go routine", runtime.NumGoroutine())
	a["Umesh"] = "Sr developer"
	a["Pathak"] = "Umesh's Lastname"
	fmt.Println(m)
	fmt.Println(a["m"])
	w.Wait()
}
