package main

import (
    "fmt"
    "strconv"
)

var keys []string = []string {
    "|0",
    "1",
    "0",
}

var vals []string = []string {
    "0||",
    "0|" ,
    "",
}

var symbol string

// algo
// 1 規則を上からチェックし、keyの文字列がsymbol内に存在しないか調べる
// 2 1つも見つからない場合、アルゴリズムの実行を停止する
// 3 1つ以上見つかった場合、記号文字列の中でも最も左端に近い部分にマッチしたvalを適用し、置換する
// 4 適用した規則がterminating ruleであった場合、アルゴリズムを停止する
// 5 ステップ1に戻り繰り返す

func find(key string) int {
    len_k := len(key)
    len_s := len(symbol)

    for i_s:=0; i_s < len_s; i_s++ {
        if symbol[i_s] == key[0] {
            if (len_s < i_s+len_k) {
                break
            }
            if symbol[i_s:i_s+len_k] == key {
                return i_s
            }
        }
    }
    return -1;
}

func main() {
    fmt.Printf("input digit num: ")
    var num int64
    fmt.Scan(&num)
    symbol = strconv.FormatInt(num, 2)

    loop:
        for i, key := range keys {
            // 各keyがsymbol内に存在しないか調べる
            res := find(key)

            if (res >= 0) {
                symbol = symbol[:res] + vals[i] + symbol[res+len(key):]
                fmt.Println(symbol)
                goto loop
            }
        }
}
