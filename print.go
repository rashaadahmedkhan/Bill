package main

import "fmt"

func main() {

    name := "Rashaad"
    age := 21
    fmt.Println("My name is", name, "and I am", age, "years old.")

    //% and a letter is called format specifier
    //%v for variables
    //%q for string
    //%T to get type of the variable
    //%0.1f to limit the number of decimal points
    fmt.Printf("My name is %v and I am %v years old", name, age)

    var str = fmt.Sprintf("My name is %v and I am %v years old", name, age)
    fmt.Println("The same string is:", str)
}