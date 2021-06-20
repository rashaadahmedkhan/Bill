package main

import "fmt"

type bill struct {
    name string
    items map[string]float64
    tip float64
}

//new bill
func newBill(name string)bill {
    b := bill {
        name: name,
        items: map[string]float64{"Oreo": 10, "Milk": 20},
        tip: 0,
    }
    return b
}

//to format the bill
func (b bill) format() string {
    fs := "Bill breakdown: \n"
    var total float64 = 0

    //list items
    for k, v := range b.items {
        fs += fmt.Sprintf("%-25v Rs%v \n", k+":", v)
        total += v
    }

    //total
    fs += fmt.Sprintf("%-25v Rs%0.2f", "Total:", total)

    return fs
}