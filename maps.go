package main

import (
    "fmt"
)

func main() {
    menu := map[string]float64{
        "oreo": 10,
        "milk": 20,
        "bread": 30,
    }
    fmt.Println(menu)
    fmt.Println(menu["milk"])

    for k,v := range menu {
        fmt.Println(k, "-", v)
    }

    phonebook := map[int]string {
        123: "rashaad",
        234: "saleem",
    }
    fmt.Println(phonebook[123])
    phonebook[123] = "rash"
    fmt.Println(phonebook)
}