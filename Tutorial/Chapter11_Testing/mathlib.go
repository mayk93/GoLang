package mathlib

import ("math")

type number struct {
    baseValue float64
    floor int
    ceil int
}

type Number struct {
    _number number
}

func (self *Number) Add (value Number) {
    self.SetValue(self._number.baseValue + value._number.baseValue)
}

func (self *Number) Sub (value float64) {
    self.SetValue(self._number.baseValue - value)
}

func (self *Number) SetValue (value float64) {
    self._number.baseValue = value
    self._number.floor = int(math.Floor(self._number.baseValue))
    self._number.ceil = int(math.Ceil(self._number.baseValue))
}

func (self *Number) GetValue () float64 {
    return self._number.baseValue
}

func (self *Number) GetFloor () int {
    return self._number.floor
}

func (self *Number) GetCeil () int {
    return self._number.ceil
}

func Average(numbers []Number) Number {
    total := new(Number)
    for _, value := range numbers {
        total.Add(value)
    }

    total.SetValue(total._number.baseValue/float64(len(numbers)))

    return *total
}
