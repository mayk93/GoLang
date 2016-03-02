package main

import ("fmt"; "time")

func f(end int) {
    for index := 0; index < end; index += 1 {
        fmt.Println(index, "out of", end)
    }
    fmt.Println("-----=====-----")
}

func main() {
    go f(10)
    var input int
    fmt.Scanln(&input)
    go f(input)
    time.Sleep(1000 * time.Millisecond)
    for index := 0; index < input; index += 1 {
        go f(index)
    }
    time.Sleep(1000 * time.Millisecond)
}
