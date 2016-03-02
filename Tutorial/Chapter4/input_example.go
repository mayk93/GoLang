package main

import "fmt"

const constant_factor float64 = 5.0

func main() {
    var input float64
    var output float64

    fmt.Println("Enter a number: ")
    fmt.Scanf("%f", &input)

    output = constant_factor * input

    fmt.Println("Result: ", output)
}
