package main

import "fmt"

func main() {
    // 配列
    // 固定長のデータ構造であり、要素数を指定して宣言します
    arr := [3]int{1, 2, 3}

    // 配列の要素を一つずつ取り出して表示します
    fmt.Println("配列:")
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // スライス
    // 可変長のデータ構造であり、要素数を指定せずに宣言します
    slice := []int{4, 5, 6}

    // スライスの要素を一つずつ取り出して表示します
    fmt.Println("\nスライス:")
    for _, val := range slice {
        fmt.Println(val)
    }

    // スライスに要素を追加します
    slice = append(slice, 7)

    // 追加された要素を含んだスライスを表示します
    fmt.Println("\n追加された要素を含んだスライス:")
    for _, val := range slice {
        fmt.Println(val)
    }
}
