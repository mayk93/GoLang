package main

import "fmt"

func main() {
    slice_1 := make([]float64, 0)
    fmt.Println("Slice 1: ", slice_1)
    fmt.Println("Len: ", len(slice_1))
    slice_1 = append(slice_1, 10)
    fmt.Println("Slice 1: ", slice_1)
    fmt.Println("Len: ", len(slice_1))
    var set [10]int
    for index := 0; index < 10; index += 1 {
        if index > 0 {
            set[index] = set[index-1] + index
        } else {
            set[index] = index
        }
    }
    for _, value := range set {
        slice_1 = append(slice_1, 2.35 * float64(value))
        }
    fmt.Println("Slice 1: ", slice_1)
    fmt.Println("Len: ", len(slice_1))
}
