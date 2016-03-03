package mathlib

import "testing"

type testpair struct {
    values []Number
    average Number
}

func TestAverage(t *testing.T) {
    n1 := new(Number)
    n2 := new(Number)
    n3 := new(Number)
    n4 := new(Number)
    n5 := new(Number)

    n1.SetValue(1.2)
    n2.SetValue(3.8)
    n3.SetValue(4.0)

    n4.SetValue(17.3)
    n5.SetValue(2.7)

    res123 := new(Number)
    res45 := new(Number)
    resempty := new(Number)

    res123.SetValue(3.0)
    res45.SetValue(10.0)

    var (tp1, tp2, tp3 testpair)

    tp1.values = []Number{*n1, *n2, *n3}
    tp1.average = *res123

    tp2.values = []Number{*n4, *n5}
    tp2.average = *res45

    tp3.values = []Number{*resempty}
    tp3.average = *resempty

    var tests = []testpair {tp1, tp2, tp3 }

    for _, pair := range tests {
        result := Average(pair.values)
        if result.GetValue() != pair.average.GetValue() {
            t.Error("Missmatch between", result.GetValue(), "and", pair.average.GetValue())
        }
    }
}
