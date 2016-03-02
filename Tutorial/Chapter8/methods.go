package main

import ("fmt"; "math"; "strconv")

type Circle struct {
    x float64
    y float64
    r float64
}

/*
func (object *ObjectType) functionName(parameters ... parametersType) returnType

This is how a method for an object of type ObjectType is created.
All objects of type ObjectType will be able to call functionName like this:

object.functionName()

The method functionName will take parameters of type parametersType and return
a value of type returnType.
*/

// All circle objects will have a method called area
// self is a pointer to the object calling this methd
func (self *Circle) area() float64 {
    return math.Pi * math.Pow(self.r, 2)
}

type Person struct {
    name string
    _ageString string
    _age int
}

func (self *Person) setAge(newAge int) {
    self._age = newAge
    ageString := strconv.Itoa(self._age)
    self._ageString = ageString
}

func (self *Person) talk() {
    fmt.Println("Hello. I am ", self.name, " and I am ", self._ageString, " years old.")
}

func main() {
    circle_ptr := new(Circle)
    fmt.Println(circle_ptr.area())  // Should be 0
    circle_ptr_2 := Circle{r: 10}
    fmt.Println(circle_ptr_2.area())
    circle_ptr.r = 5
    fmt.Println(circle_ptr.area())

    person := Person{name: "John Smith"}
    person.setAge(22)
    person.talk()
    person.setAge(34)
    person.talk()
}
