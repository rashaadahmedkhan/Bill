package main

import (
    "fmt"
    "strings"
    "sort"
)

func main() {

    greetings := "hi hello there how are you on this fine day?"
    //these strings functions do not alter the original
    fmt.Println(strings.Contains(greetings, "you")) //returns true
    fmt.Println(strings.ReplaceAll(greetings, "day", "night")) //replace string
    fmt.Println(strings.ToUpper(greetings)) //converts entire string to upper case
    fmt.Println(strings.Index(greetings, "ll")) //index starts from 0 and it will print first appearance
    fmt.Println(strings.Split(greetings, " ")) //gives a slice of all words in the string

    ages := []int{20, 40, 30, 25, 10, 15}
    sort.Ints(ages)
    fmt.Println(ages)
    index := sort.SearchInts(ages, 30)
    fmt.Println(index) //after arranging it will give where it is
}