package main

import (
	"fmt"

	"github.com/mickamy/gopanix"
)

func main() {
	defer gopanix.Handle(true)

	fmt.Println("🧪 This program will panic for testing gopanix...")
	panic("🔴 something went wrong!")
}
