package main

import "fmt"

func main() {
    var hello_world = "Hello World!"  // Assign directly, type is infered

    /*
    Declared a string variable and assigned after declaration.
    */
    var new_hello_world string
    new_hello_world = "New Hello World!"

    fmt.Println(hello_world)
    fmt.Println(new_hello_world)
}
