package main

import "fmt"

func main() {
    // マップの宣言と初期化
    person := map[string]int{
        "Alice": 30,
        "Bob":   25,
        "Charlie": 35,
    }

    // マップの要素へのアクセス
    fmt.Println("Alice's age:", person["Alice"])
    fmt.Println("Bob's age:", person["Bob"])

    // マップに新しい要素を追加
    person["David"] = 40

    // マップの要素をループして表示
    fmt.Println("\nAll persons:")
    for name, age := range person {
        fmt.Printf("%s's age: %d\n", name, age)
    }
}
