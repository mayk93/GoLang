package main

import "fmt"

func main() {
    fmt.Println("[Length of 'Hello World!'] ", len("Hello World!"))
    fmt.Println("[Third character of 'Hello World!'] ", "Hello World!"[2])
    fmt.Println("[Concatenation of 'Hello', ' World' and '!'] ",
                "Hello " + "World" + "!")
}
