package main

import "fmt"

// This function must not be called on it's own
func error_example() int {
    // We can't do this, maps need to be initialized
    var simple_map map[string]int
    simple_map["kw_0"] = 10
    return simple_map["kw_0"]
}

func correct_example() int {
    simple_map := make(map[string]int)
    simple_map["kw_0"] = 10
    return simple_map["kw_0"]
}

// Feed this function a function that may cause errors
func protect(to_run func() int) int {
	defer func() {
		stack_trace := recover()
        if stack_trace != nil {
			fmt.Println("Run time panic: ", stack_trace)
		}
	}()

	result := to_run()

    return result
}

func main()  {
    result_err := protect(error_example)
    fmt.Println("Error Res: ", result_err)
    result_cor := protect(correct_example)
    fmt.Println("Correct Res: ", result_cor)

    dictionary := make(map[string]int)
    dictionary["kw_0"] = 15
    dictionary["kw_1"] = 25
    if name, ok := dictionary["not_in_dict"]; ok {
      fmt.Println(name, ok)
    } else {
        fmt.Println("Not found.")
    }

    if name, ok := dictionary["kw_0"]; ok {
      fmt.Println(name, ok)
    } else {
        fmt.Println("Not found.")
    }
}
