package main

import "fmt"

func increment(value_ptr *int, ammount int) {
    *value_ptr += ammount
}

func main()  {
    value := 15
    ammount := 10
    fmt.Println("Value before: ", value)
    increment(&value, ammount)
    fmt.Println("Value after: ", value)
}
