package main

import (
	"fmt"
	"time"
)

// go keyword description

func main() {
	go SayHelloRoutine()
}

// SayHelloRoutine hello!
func SayHelloRoutine() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("Routine says G'day")
	}
}
