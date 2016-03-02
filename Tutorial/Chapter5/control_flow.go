package main

import "fmt"

func verbose_for()  {
    fmt.Println("Verbose for")
    index := 0
    for index < 10 {
        fmt.Println("[Verbose] Now at iteration: ", index)
        index += 1
    }
}

func concise_for()  {
    fmt.Println("Concise for")
    for index := 0; index < 10; index += 1 {
        fmt.Println("[Concise] Now at iteration: ", index)
    }
}

func if_for()  {
    fmt.Println("If else for")
    for index := 0; index < 10; index += 1 {
        if index % 2 == 0  {
            fmt.Println(index, " is an even number.")
        } else {
            fmt.Println(index, " is an odd number.")
        }
    }
}

func get_number_literal(number int) string  {
    switch number {
        case 0: return "Zero"
        case 1: return "One"
        case 2: return "Two"
        case 3: return "Three"
        case 4: return "Four"
        case 5: return "Five"
        default: return "Case not built yet."
    }
}

func switch_example()  {
    fmt.Println("Switch example")
    for index := 0; index < 10; index += 1 {
        fmt.Println("The number ", index, " is written as ", get_number_literal(index))
    }
}

func main() {
    verbose_for()
    fmt.Println("----------")
    concise_for()
    fmt.Println("----------")
    if_for()
    fmt.Println("----------")
    switch_example()
    fmt.Println("----------")
}
