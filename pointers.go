package main

import "fmt"

func main() {
    // 変数の宣言
    var num int = 10

    // num のポインタを取得し、ポインタ変数に格納
    var ptr *int = &num

    // ポインタを介して変数の値を更新
    *ptr = 20

    // 変数の値を表示
    fmt.Println("Value of num:", num)

    // 関数にポインタを渡して、関数内で変数の値を更新
    modifyValue(ptr)

    // 変数の値を表示
    fmt.Println("Value of num after modification:", num)
}

// modifyValue 関数は、ポインタを介して変数の値を更新します
func modifyValue(ptr *int) {
    *ptr = 30
}
