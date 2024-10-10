package main

import (
	"fmt"

	"github.com/AmirH21/puppy"
)

func main() {
	fmt.Println(puppy.Woof())
	fmt.Println(puppy.Woofs())
	fmt.Println(puppy.Dog())
	fmt.Println(puppy.Dogs())

	for {
		fmt.Println("Press ctrl + c")
	}

}
