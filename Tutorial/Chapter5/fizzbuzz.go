/*
Write a program that prints the numbers from 1 to 100.
But for multiples of three print "Fizz" instead of the number and
for the multiples of five print "Buzz".
For numbers which are multiples of both three and five print "FizzBuzz".
*/

package main

import "fmt"  // Printing
import "os"  // Command Line Arguments
import "strconv"  // Convert string command line arguments to ints
import "errors"  // Return errors

const MaxInt = int(^uint(0) >> 1)  // Maximum Possible value

func clamp(input ... int) (int, error) {
    var min int
    var max int
    var val int
    // No input, nothing to clamp, retuning 0
    if len(input) == 0 {
        return 0, nil
    // One input value, clamp between 0 and infinity (max int)
    } else if len(input) == 1 {
        min = 0
        max = MaxInt
    // Two input values, the value to clamp and a minimum.
    } else if len(input) == 2 {
        min = input[1]
        max = MaxInt
    // Three input values, the value to clamp, min and max
    } else if len(input) == 3 {
        min = input[1]
        max = input[2]
    // Bad input.
    } else if len(input) > 3 {
        return 0, errors.New("Input too long.")
    }

    val = input[0]

    if min <= val && val <= max {
        return val, nil
    } else if val < min {
        return min, nil
    } else if val > max {
        return max, nil
    }

    return 0, errors.New("[1] Reached a bad point.")
}

func get_range() (int, int, error)  {
    arguments := os.Args[1:]  // The first argument is the program
    if len(arguments) == 0 {
        return 0, 100, nil  // This is the default range
    } else if len(arguments) == 1 {
        var parsed_value int64
        var clamped_value int
        var err error

        parsed_value, err = strconv.ParseInt(arguments[0], 10, 64)
        clamped_value, err = clamp(int(parsed_value))

        return 0, clamped_value, err  // The agument is considered
                                     // to be the end of the range
    } else if len(arguments) == 2 {
        var parsed_value_1 int64
        var parsed_value_2 int64
        var start int
        var end int
        var err error

        parsed_value_1, err = strconv.ParseInt(arguments[0], 10, 64)
        parsed_value_2, err = strconv.ParseInt(arguments[1], 10, 64)

        start, err = clamp(int(parsed_value_1))
        end, err = clamp(int(parsed_value_2))

        if start < end {
            return start, end, err
        } else {
            return end, start, err
        }

    } else if len(arguments) > 2 {
        return 0, 0, errors.New("Too many arguments.")
    }

    return 0, 0, errors.New("[0] Reached a bad point.")
}

func fizzBuzz(start, end int) {
    for index := start; index < end; index += 1 {
        if index % 3 == 0 && index % 5 == 0 {
            fmt.Println("FizzBuzz")
        } else if index % 3 == 0 {
            fmt.Println("Fizz")
        } else if index % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(index)
        }
    }
}

func main()  {
    var (start, end int)
    var err error
    start, end, err = get_range()

    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    fmt.Println("Fizzbuzz between:")
    fmt.Println("Start: ", start)
    fmt.Println("End: ", end)

    fizzBuzz(start, end)
}
