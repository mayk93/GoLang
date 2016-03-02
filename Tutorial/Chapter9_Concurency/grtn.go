package main

import ("fmt"; "time"; "strconv"; "os"; "errors")

func goroutine(value int, start int, result []int) {
    fmt.Println("Start Goroutine with value ", value)
    total := 0
    for index := value; index < 2 * value; index += 1 {
        total += (index + value) * (total + 1)
        fmt.Println("Goroutine ",index," - Total so far: ", total)
        time.Sleep(700 * time.Millisecond)
        if index % 10 == 0 {
            fmt.Println("Still processing.")
        }
    }

    result[value-start] = total
}

func parse_args(arguments []string) (int, int, error) {
    start, err := strconv.ParseInt(arguments[0], 10, 64)
    if err != nil {
        return 0, 0, err
    }
    end, err := strconv.ParseInt(arguments[1], 10, 64)
    if err != nil {
        return 0, 0, err
    }

    return int(start), int(end), nil
}

func get_args() (int, int, error)  {
    var (start, end int)
    var err error

    defer func () {
        stack_trace := recover()
        if stack_trace != nil {
			fmt.Println("Run time panic: ", stack_trace)
            start = 0
            end = 0
            err = errors.New("[1] Invalid arguments")
		}
    }()

    arguments := os.Args[1:]
    if len(arguments) != 2 {
        panic("[0] Invalid number of arguments")
    } else {
        start, end, err = parse_args(arguments)
        if err != nil {
            panic(err)
        }
    }

    return start, end, err
}

func main()  {
    start, end, err := get_args()
    if err != nil {
        fmt.Println("Invalid arguments.")
        return
    } else {
        fmt.Println("[Start] Starting ", end-start, " goroutines between ", start, " and ", end, ".")
        result := make([]int, 0)
        for index := 0; index < end-start; index += 1 {
            result = append(result, 0)
        }
        goroutine_number := start
        for index := start; index < end; index += 1 {
            defer func () {
                fmt.Println("Goroutine ", goroutine_number, " returned ", result)
                goroutine_number += 1
            }()
            go goroutine(index, start, result)
        }
    }
    var new_start int
    //var new_end int

    fmt.Println("Enter new start value.")
    fmt.Scanln(&new_start)
    fmt.Println(new_start)
}
