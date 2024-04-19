package main

import "fmt"

// Person 構造体の定義
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // 構造体のインスタンス化
    person1 := Person{Name: "Alice", Age: 30, City: "New York"}
    person2 := Person{Name: "Bob", Age: 25, City: "Los Angeles"}

    // 構造体のフィールドへのアクセス
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)
    fmt.Println("City:", person1.City)

    // 構造体のフィールドの更新
    person1.Age = 35

    // 構造体のフィールドをループして表示
    fmt.Println("\nAll persons:")
    persons := []Person{person1, person2}
    for _, p := range persons {
        fmt.Printf("Name: %s, Age: %d, City: %s\n", p.Name, p.Age, p.City)
    }
}
