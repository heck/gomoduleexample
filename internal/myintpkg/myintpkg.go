package myintpkg

import "fmt"

// Run func
func Run() {
	fmt.Printf("Hello, internal example!\n")
}

// MyInternalFunc func
func MyInternalFunc() {
	fmt.Println("Hello, internal function called from external function")
}
