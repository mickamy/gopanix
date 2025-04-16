package main

import (
	"fmt"

	"github.com/mickamy/gopanix/gopanix"
)

func main() {
	defer gopanix.Handle()

	fmt.Println("🧪 This program will panic for testing gopanix...")
	panic("🔴 something went wrong!")
}
