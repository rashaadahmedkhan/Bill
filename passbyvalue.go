package main

import (
    "fmt"
)

func updateName(x string) string{
    x = "rash"
    return x
}

func updateMenu(y map[string]int) {
    y["coffee"] = 100
}

//type1 (non-pointer values): string, int, bool, float, array, struct
//type2 (pointer wrapper values): slice, map, function
func main() {
    name := "rashaad"
    name = updateName(name)
    fmt.Println("Memory address of name is:", &name)
    fmt.Println(name)

    m := &name

    menu := map[string] int {
        "oreo": 10,
        "milk": 20,
    }
    updateMenu(menu)
    fmt.Println(menu)
}