package main

import (
	"fmt"

	"github.com/umeshpathak1234/sampleSvc/api"

	"github.com/google/logger"
)

func main() {

	fmt.Println("Hello World !")
	logger.Info("I'm about to do something!")
	api.Info()
}
