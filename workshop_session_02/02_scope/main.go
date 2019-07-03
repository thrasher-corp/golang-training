package main

import "fmt"

var global = "global scoped variable"

func main() {
	fmt.Println(global)
}

func hello() {
	var local = "local scoped variable"
	fmt.Println(local)
	fmt.Println(global)
}
