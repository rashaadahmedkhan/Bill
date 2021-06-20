package main

import "fmt"

func main() {
    mybill := newBill("Saleem bhai's bill")
    fmt.Println(mybill.format())
}