package main

import ("fmt"; "mymath")

func main() {
    array := []float64{3.5, 1.23, 7.42, 9.8135, 2.22, 6.21, 7.0}
    average := mymath.Average(array)
    fmt.Println("Average:",average)
}
