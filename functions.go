package main

import (
	"fmt"
	"math"
)

func sayHi(n string) {
	fmt.Printf("Hi %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
    for _, v := range n {
        f(v)
    }
}

func circleArea(r float64) float64 {
    return math.Pi*r*r
}

func main() {
	sayHi("Rash")
	sayHi("Saleem")
	sayBye("Bhai")
	cycleNames([]string{"solid", "liquid", "gas"}, sayHi)
	a1 := circleArea(10)
	fmt.Println(a1)
}