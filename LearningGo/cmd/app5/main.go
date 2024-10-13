package main

import "fmt"

type gasEngine struct {
	mpg     uint32
	gallons uint32
}

func (e gasEngine) milesleft() uint32 {
	return e.mpg * e.gallons
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 20, gallons: 50}
	fmt.Println(myEngine.mpg, myEngine.gallons)

	fmt.Println(myEngine.milesleft())

	fmt.Println("--------------------")

}
