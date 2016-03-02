package main

import ("fmt"; "time"; "math"; "math/rand")

type Object struct {
    filed_1 int
    filed_2 string
}

func getRandomLetter(seed int) string {
    var letter string
    if rand.Intn(10) % 2 == 0 {
        switch seed {
        case 0:
            letter = "a"
        case 1:
            letter = "b"
        case 2:
            letter = "c"
        case 3:
            letter = "d"
        case 4:
            letter = "e"
        }
    } else {
        switch seed {
        case 0:
            letter = "f"
        case 1:
            letter = "g"
        case 2:
            letter = "h"
        case 3:
            letter = "i"
        case 4:
            letter = "j"
        }
    }

    return letter
}

func getRandomString(seed int) string {
    randomString := ""
    for index := 0; index < rand.Intn(100); index += 1 {
        var letter_seed int
        if seed % 3 == 0 {
            letter_seed = int((4.0 * math.Pow(float64(seed - index), 2.0) - float64(index+12-(seed/3)) + float64(rand.Intn(10)))) % 5
        } else if seed % 3 == 1 {
            letter_seed = int((2.0 * math.Pow(float64(seed + index), 4.0) - float64(index*seed+31-(seed*seed)) - float64(rand.Intn(10)))) % 5
        } else if seed % 3 == 2 {
            letter_seed = int((float64(rand.Intn(10)) * math.Pow(float64(seed*index - index*index), 2.0) - float64(index+33-seed) * float64(rand.Intn(10)))) % 5
        }
        randomString += getRandomLetter(letter_seed)
    }

    return randomString
}

func sender(channel chan Object) {
    for {
        channel <- Object{filed_1: rand.Intn(500), filed_2: getRandomString(rand.Intn(100))}
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
}

func sender_2(channel chan Object) {
    for {
        channel <- Object{filed_1: rand.Intn(500), filed_2: "sender_2"}
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    }
}

func receiver(channel chan Object) {
    for {
        obj := <- channel
        fmt.Println("Received: ",obj)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    var channel chan Object = make(chan Object)

    go sender(channel)
    go sender_2(channel)
    go receiver(channel)

    time.Sleep(60 * 1000 * time.Millisecond)
}
