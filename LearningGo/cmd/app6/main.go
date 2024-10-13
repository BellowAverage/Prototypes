package main

import "fmt"

func main() {

	var thing1 = [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Printf("%p", &thing1)

	square(&thing1)

	fmt.Printf("\nThe result is: %v\n", thing1)

}

func square(thing2 *[5]float64) {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	fmt.Printf("\n%p", thing2)
}
