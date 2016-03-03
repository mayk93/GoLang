package main

import ("fmt"; "time"; "math/rand")

// Arrays cannot be declared as constant.
var NAMES = [...]string {"John", "Dan", "Jim", "Bob", "Hellen", "Andrea", "Rob"}

type Person struct {
    age int
    name string
    channel int
}

func (self *Person) talk()  {
    fmt.Println("Hello, I am",self.name,"and I am",self.age,"years old. I arrived on channel ", self.channel)
}

func sender(channel chan Person, channelNumber int) {
    for {
        channel <- Person{age: 20 + rand.Intn(20), name: NAMES[rand.Intn(len(NAMES))], channel: channelNumber}
    }
}

func receiver(channel_1 chan Person, channel_2 chan Person, channel_3 chan Person) {
    for {
      select {
        case person_1 := <- channel_1:
            person_1.talk()
        case person_2 := <- channel_2:
            person_2.talk()
        case person_3 := <- channel_3:
            person_3.talk()
        time.Sleep(time.Duration(1) * time.Second)
      }
    }
}

func main() {
    var senders int
    var receivers int
    var input string

    channel_1 := make(chan Person, 10)
    channel_2 := make(chan Person, 10)
    channel_3 := make(chan Person, 10)
    channels := [...]chan Person{channel_1, channel_2, channel_3}

    fmt.Println("Input number of senders:")
    fmt.Scanln(&senders)

    for index := 0; index < senders; index += 1 {
        go sender(channels[index % len(channels)], index % len(channels))
    }

    fmt.Println("Input number of receivers:")
    fmt.Scanln(&receivers)

    for index := 0; index < receivers; index += 1 {
        go receiver(channel_1, channel_2, channel_3)
    }

    fmt.Println("Press ENTER to stop.")
    for {
            fmt.Scanln(&input)
            if input == "\n" {
                return
            } else {
                fmt.Println("Press ENTER to stop.")
            }
    }
}
