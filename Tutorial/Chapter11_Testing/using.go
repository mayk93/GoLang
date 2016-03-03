package main

import ("fmt"; "mathlib")

func main() {
    n1 := new(mathlib.Number)
    n2 := new(mathlib.Number)

    n1.SetValue(3.5)
    n2.SetValue(5.21)

    fmt.Println("N1:",n1.GetValue())
    fmt.Println("N1 floor:",n1.GetFloor())
    fmt.Println("N2:",n2.GetValue())
    fmt.Println("N2 ceil:",n2.GetCeil())

    naux := new(mathlib.Number)
    naux.SetValue(6.5)
    n1.Add(*naux)

    fmt.Println("N1:",n1.GetValue())
    fmt.Println("N1 floor:",n1.GetFloor())

    avg := mathlib.Average([]mathlib.Number{*n1, *n2})
    fmt.Println("Avg:",avg.GetValue())
}
