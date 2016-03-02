package main

import ("fmt"; "math"; "os"; "strconv")

type segment struct {
    x float64
    y float64
}

func no_struct_distance(x1, x2, y1, y2 float64) float64 {
    delta_x := x2 - x1
    delta_y := y2 - y1

    return math.Sqrt(math.Pow(delta_x, 2) + math.Pow(delta_y, 2))
}

func struct_distance(segment_1, segment_2 segment) float64 {
    delta_x := segment_2.x - segment_1.x
    delta_y := segment_2.y - segment_1.y

    return math.Sqrt(math.Pow(delta_x, 2) + math.Pow(delta_y, 2))
}

/* Will add a method here */
func (self *segment) distanceTo(other segment) float64 {
    delta_x := other.x - self.x
    delta_y := other.y - self.y

    return math.Sqrt(math.Pow(delta_x, 2) + math.Pow(delta_y, 2))
}

func no_args()  {
    /*
    5
    |        b
    |
    |
    |  a
    |__.__.__.__.__
    |

    a = (x1, y1) = (0.9, 1.1)
    b = (x2, y2) = (3.2, 4.3)

    */

    var (x1, x2, y1, y2 float64)
    x1 = 0.9
    x2 = 3.2
    y1 = 1.1
    y2 = 4.3

    distance := no_struct_distance(x1, x2, y1, y2)

    fmt.Println("[No Struct] Distance between (",x1,",",y1,") and (",x2,",",y2,"): ", distance)

    var segment_1 segment
    var segment_2 segment

    segment_1.x = 0.9
    segment_1.y = 1.1

    segment_2.x = 3.2
    segment_2.y = 4.3

    distance = struct_distance(segment_1, segment_2)

    fmt.Println("[Struct] Distance between (",segment_1,") and (",segment_2,"): ", distance)
}

func get_float(to_parse string) float64 {
    var return_value float64
    // This is basically an anonymous function that runs in the background
    defer func() {
        stack_trace := recover()
        if stack_trace != nil {
            fmt.Println("Run time panic: ", stack_trace)
            return_value = 0.0
        }
    }()

    return_value, err := strconv.ParseFloat(to_parse, 64)
    if err != nil {
        panic(err)
    }

    return return_value
}

func with_args(arguments []string)  {
    var (x1, y1, x2, y2 float64)
    for index, value := range arguments {
        switch index {
        case 0:
            x1 = get_float(value)
        case 1:
            y1 = get_float(value)
        case 2:
            x2 = get_float(value)
        case 3:
            y2 = get_float(value)
        default:
            fmt.Println("There must never be more than 4 values.")
            panic("Number of values exceded.")
        }
    }

    segment_1 := segment{x: x1, y: y1}
    segment_2 := segment{x: x2, y: y2}

    fmt.Println("[Method] Distance: ", segment_1.distanceTo(segment_2))
}

func main()  {
    arguments := os.Args[1:]
    if len(arguments) == 4 {
        with_args(arguments)
    } else if len(arguments) == 0 {
        no_args()
    } else {
        fmt.Println("Must have 0 or 4 arguments. With 0 arguments, will use default values, with 4 will use x1, y1, x2, y2")
    }
}
