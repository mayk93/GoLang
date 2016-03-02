package main

import "fmt"

func main() {
    //  Short assignment.

    var_1 := 1
    var_2 := 2 - 1

    var_3 := "String"
    var_4 := "Different string."

    fmt.Println("Is var_1 (",var_1,") equal to var_2 (",var_2,") ? ",
                var_1 == var_2)

    fmt.Println("What about var_3 (",var_3,") and var_4 (",var_4,") ? ",
                var_3 == var_4)
}
