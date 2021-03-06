package main

import (
    "fmt"
    "strings"
)

func getInitials(n string) (string, string) {
    s := strings.ToUpper(n)
    names := strings.Split(s, " ")
    var initials []string
    for _, v := range names {
        initials = append(initials, v[:1])
    }
}

func main() {
    getInitials("hello world")
}