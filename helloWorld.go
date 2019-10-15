package main

import (
	"fmt"

	"Users/god/Desktop/Golang/sampleSvc/api"

	"github.com/google/logger"
)

func main() {

	fmt.Println("Hello World !")
	logger.Info("I'm about to do something!")
	api.Info()
}
