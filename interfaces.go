package main

import "fmt"

// Animal インターフェースの定義
type Animal interface {
    Speak() string
}

// Dog 構造体の定義
type Dog struct{}

// Dog 構造体に Speak メソッドを実装
func (d Dog) Speak() string {
    return "Woof!"
}

// Cat 構造体の定義
type Cat struct{}

// Cat 構造体に Speak メソッドを実装
func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    // Animal インターフェースを実装する構造体を使用
    animals := []Animal{Dog{}, Cat{}}

    // 各動物の発声を表示
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
