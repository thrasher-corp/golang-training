package main

import "fmt"

// (    )
// {    }
// [    ]
// ,
// .
// ;
// :

func main() {
	HelloPerson("ryan", 5, 183, 93)
	fmt.Printf("hello my name is %s is %d tall", "Guy", 1234)
}

func HelloPerson(name string, age int, height float64, weight float64) {
	fmt.Println("Hello:", name)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("HEYYYYYYY!")
}
