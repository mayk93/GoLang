package main

import "fmt"

func main() {
    var variable string
    variable = "This is a string."
    fmt.Println(variable)
    variable = variable + ` This is another string that was
                            concatenated to the initial string.`
    fmt.Println(variable)
}
