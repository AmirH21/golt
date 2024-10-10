package main

import (
	"fmt"
	"math/rand"

	"github.com/AmirH21/puppy"
)

func main() {
	fmt.Println(puppy.Woof())
	fmt.Println(puppy.Woofs())
	fmt.Println(puppy.Dog())
	fmt.Println(puppy.Dogs())

	for i := 1; i <= 42; i++ {

		fmt.Printf("This is iteration number %v\n", i)
		ranum := rand.Intn(5)

		switch ranum {
		case 0:
			fmt.Println("It's 0")
		case 1:
			fmt.Println("It's 1")
		case 2:
			fmt.Println("It's 2")
		case 3:
			fmt.Println("It's 3")
		case 4:
			fmt.Println("It's 4")
		}
		fmt.Println("")
	}

}
