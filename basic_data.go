package main

import "fmt"

func main() {
    // 配列
    arr := [3]int{1, 2, 3}

    // スライス
    slice := []int{4, 5, 6}

    // マップ
    m := map[string]int{"a": 1, "b": 2}

    // コードを追加する
    fmt.Println(arr)
    fmt.Println(slice)
    fmt.Println(m)
}

