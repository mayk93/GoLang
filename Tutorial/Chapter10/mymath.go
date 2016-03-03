package mymath

func Average(array []float64) float64 {
    var total float64
    for _, value := range array {
        total += value
    }

    if len(array) == 0 {
        return 0.0
    }
    return total / float64(len(array))
}
