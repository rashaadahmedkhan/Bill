package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
    fmt.Print(prompt)
    input, err := r.ReadString('\n')

    return strings.TrimSpace(input), err
}

func createBill() bill {
    reader := bufio.NewReader(os.Stdin)
    name, _ := getInput("Create a new bill name: ", reader)

    b := newBill(name)
    fmt.Println("Created Bill for:", b.name)

    return b
}

func promptOptions(b bill) {
    reader := bufio.NewReader(os.Stdin)
    opt, _ := getInput("Choose option (a: add item, s: save bill, t: update tip): ", reader)
    switch opt {
        case "a":
            name, _ := getInput("Item name:", reader)
            price, _ := getInput("Item price:", reader)
            p, err := strconv.ParseFloat(price, 64)
            if err != nil {
                fmt.Println("Please enter a numeric value for the price.")
                promptOptions(b)
            }
            b.addItem(name, p)
            fmt.Println("Item added:", name, price)
            promptOptions(b)
        case "t":
            tip, _ := getInput("Enter tip amount: Rs ", reader)
            t, err := strconv.ParseFloat(tip, 64)
            if err != nil {
                fmt.Println("Please enter a numeric value for the tip.")
                promptOptions(b)
            }
            b.updateTip(t)
            fmt.Println("Tip added: Rs", tip)
            promptOptions(b)
        case "s":
            fmt.Println("File saved.", b)
        default:
            fmt.Println("Invalid option. Try Again.")
            promptOptions(b)
    }
}

func main() {
    mybill := createBill()
    promptOptions(mybill)
}