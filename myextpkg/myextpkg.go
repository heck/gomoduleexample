package myextpkg

import (
	"fmt"

	"github.com/heck/gomoduleexample/internal/myintpkg"
)

// Run func
func Run() {
	myintpkg.MyInternalFunc()
	fmt.Printf("Hello, external example!\n")
}
