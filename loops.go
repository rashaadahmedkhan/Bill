package main

import "fmt"

func main() {

    for i:= 0; i < 5; i++ {
        fmt.Println("Value of i is:", i)
    }

    names := []string{"a", "b", "c", "d"}
    for index, value := range names {
        fmt.Printf("The value at index %v is %v \n", index, value)
    }
    for _, value := range names {
        fmt.Printf("The value is %v \n", value)
    }
}