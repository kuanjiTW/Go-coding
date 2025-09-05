package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		Debug       bool      = false
		LogLevel    string    = "info"
		startUpTime time.Time = time.Now()
	)

	fmt.Println(Debug, LogLevel, startUpTime)
}
