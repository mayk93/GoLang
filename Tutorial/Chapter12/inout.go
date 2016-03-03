package main

import ("fmt"; "os"; "errors"; "io/ioutil"; "strings"; "strconv"; "sort")

type Object struct {
    number float64
    name string
    _id int
}

type ObjectList []Object
type ObjectListName []Object

func (self ObjectList) Len() int {
    return len(self)
}

func (self ObjectList) Less(index_1, index_2 int) bool {
    return self[index_1].number <  self[index_2].number
}

func (self ObjectList) Swap(index_1, index_2 int) {
    self[index_1], self[index_2] = self[index_2], self[index_1]
}

func (self ObjectListName) Len() int {
    return len(self)
}

func (self ObjectListName) Less(index_1, index_2 int) bool {
    return self[index_1].name <  self[index_2].name
}

func (self ObjectListName) Swap(index_1, index_2 int) {
    self[index_1], self[index_2] = self[index_2], self[index_1]
}

func generateId() func() int  {
    start := 0

    return func() int {
        start += 1
        return start
    }
}

func getInputFile() (string, error) {
    arguments := os.Args[1:]
    if len(arguments) != 1 {
        return "", errors.New("Invalid number of arguments. Call with only 1 existing file name.")
    }

    return arguments[0], nil
}

func getNumbers(input string) []float64 {
    numbers := make([]float64, 0)
    for _, value := range strings.Split(input, " ") {
        number, err := strconv.ParseFloat(value, 64)
        if err == nil {
            numbers = append(numbers, number)
        }
    }
    return numbers
}

func getObjects(numbers []float64, getId func() int) []Object {
    objects := make([]Object, 0)
    for _, value := range numbers {
        newObject := Object{number: value, name: strconv.FormatFloat(value, 'f', 3, 64), _id: getId()}
        objects = append(objects, newObject)
    }

    return objects
}

func main() {
    getId := generateId()
    inputFile, err := getInputFile()

    if err != nil {
        fmt.Println(err)
        return
    }
    bytestring, err := ioutil.ReadFile(inputFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    content := string(bytestring)
    numbers := getNumbers(content)
    objects := getObjects(numbers, getId)

    fmt.Println("Objects:",objects)

    sort.Sort(ObjectList(objects))
    fmt.Println("Objects ordered by number:",objects)

    sort.Sort(ObjectListName(objects))
    fmt.Println("Objects ordered by name:",objects)

    outputFile, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer outputFile.Close()

    for _, object := range objects {
        outputFile.WriteString(object.name + "\n")
    }
}
