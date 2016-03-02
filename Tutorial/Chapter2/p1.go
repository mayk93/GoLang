/*
Reading the program:

Line 13: Create a new executable program,
Line 15: which references the fmt library
Line 27: and contains one function called main.
Line 27: That function takes no arguments,
Line 27: doesn't return anything and does the following:
Line 28: Access the Println function contained inside of the fmt package and
Line 28: invoke it using one argument – the string Hello World.
*/

package main  // This is the name of the package - The file we are currently in.

import "fmt"  // This is a package (a file, a library) we are importing.
              // fmt stands for format - We use it to print stuff.

/*
All functions:
1. start with the keyword func,
2. followed by the name of the function (main in this case),
3. a list of zero or more “parameters” surrounded by parentheses,
4. an optional return type and
5. a “body” which is surrounded by curly braces.
*/

func main() {
    fmt.Println("Hello World!")
}
