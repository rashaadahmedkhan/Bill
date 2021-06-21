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
func (b *bill) format() string {
    fs := "Bill breakdown: \n"
    var total float64 = 0

    //list items
    for k, v := range b.items {
        fs += fmt.Sprintf("%-25v Rs %v \n", k+":", v)
        total += v
    }

    //tip
    fs += fmt.Sprintf("%-25v Rs %0.2f", "Tip:", b.tip,)

    //total
    fs += fmt.Sprintf("\n%-25v Rs %0.2f", "Total:", total+b.tip)

    return fs
}

//to update tip
func (b *bill) updateTip(tip float64) {
    b.tip = tip
}

//add items to the bill
func (b *bill) addItem(name string, price float64) {
    b.items[name] = price
}