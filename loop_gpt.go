package main

import "fmt"

func main() {
    // for ループ
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // if - else 文
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is not greater than 5")
    }

    // switch - case 文
    y := 3
    switch y {
    case 1:
        fmt.Println("y is 1")
    case 2:
        fmt.Println("y is 2")
    default:
        fmt.Println("y is neither 1 nor 2")
    }
}
