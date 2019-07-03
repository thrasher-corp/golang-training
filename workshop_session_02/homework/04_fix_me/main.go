// This is a broken file please fix to allow it to build.
package main

import "fmt"

func main() {
	if isTrue() {
		fmt.Println("This should be displayed")
	}

	if isFalse() {
		fmt.Println("This should NOT be displayed")
	}
}

func isTrue() bool {
	return false
}

func isFalse() bool {
	return true
}
