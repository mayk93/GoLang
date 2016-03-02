package main

import ("fmt"; "math")

type Color struct {
    r float64
    g float64
    b float64
}

func (self *Color) add(other Color)  {
    self.r += other.r
    self.g += other.g
    self.b += other.b
}

type Circle struct {
    x float64
    y float64
    r float64
    _color Color
}

type Rectange struct {
    // x and y refer to the LOWER LEFT corner
    x float64
    y float64
    length float64
    width float64
    _color Color
}

func (self *Circle) area() float64 {
    return math.Pi * math.Pow(self.r, 2)
}

func (self *Rectange) area() float64 {
    return self.length * self.width
}

func (self *Circle) color() Color {
    return self._color
}

func (self *Rectange) color() Color {
    return self._color
}

type Shape interface {
    area() float64
    color() Color
}

func totalArea(shapes ...Shape) float64 {
    var total float64
    for _, shape := range shapes {
        total += shape.area()
    }
    return total
}

func get_average(total Color, count float64) Color {
    return Color{r: total.r/count, g: total.g/count, b: total.b/count}
}

func colorAverage(shapes ...Shape) Color {
    var total Color
    for _, shape := range shapes {
        total.add(shape.color())
    }
    return get_average(total, float64(len(shapes)))
}

func main() {
    rectange_1 := Rectange{x: 10, y: 10, length: 15, width: 5, _color: Color{r: 10, g: 15, b: 20}}
    rectange_2 := Rectange{x: 30, y: 35, length: 7, width: 10, _color: Color{r: 25, g: 30, b: 45}}
    circle_1 := Circle{x: 100, y: 95, r: 10, _color: Color{r:10, g: 100, b: 255}}
    circle_2 := Circle{x: 155, y: 201, r: 13, _color: Color{r:255, g: 200, b: 35}}

    fmt.Println("Total area: ", totalArea(&rectange_1, &rectange_2, &circle_1, &circle_2))
    fmt.Println("Avg Color: ", colorAverage(&rectange_1, &rectange_2, &circle_1, &circle_2))
}
