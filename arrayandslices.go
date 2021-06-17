package main

import "fmt"

func main() {
    //var ages [3]int = [3]int{1,2,3}
    var ages = [3]int{1,2,3}
    fmt.Println(ages, len(ages))

    names := [4]string{"a", "b", "c", "d"}
    names[1] = "e"
    fmt.Println(names, len(names))

    //slices
    var scores = []int{100, 200, 300}
    scores[2] = 500
    scores = append(scores, 800)
    fmt.Println(scores, len(scores))

    //slice range
    rangeOne := names[1:3] //from position 1-3
    rangeTwo := names[2:] //from position 2 till the end
    rangeThree := names[:3] //from beginning till third position, not including the third position
    fmt.Println(rangeOne, rangeTwo, rangeThree)
}