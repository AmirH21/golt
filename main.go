package main

import (
	"fmt"
	"math/rand"

	"github.com/AmirH21/puppy"
)

var x, y = rand.Intn(10), rand.Intn(10)

func init() {
	fmt.Printf("X is %v\nY is %v\n", x, y)
}

func main() {
	fmt.Println(puppy.Woof())
	fmt.Println(puppy.Woofs())
	fmt.Println(puppy.Dog())
	fmt.Println(puppy.Dogs())

	/*
		switch {
		case x < 4 && y < 4:
			fmt.Println("They are both less than 4")

		case x > 6 && y > 6:
			fmt.Println("They are both greater than 6")

		case x >= 4 && x <= 6:
			fmt.Println("X is more than or equal to 4 and less than or equal to 6")

		case y != 5:
			fmt.Println("Y is not 5")

		default:
			fmt.Println("None of them happened")
		}
	*/

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

}
