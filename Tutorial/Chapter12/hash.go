package main

import ("fmt"; "crypto/md5"; "io/ioutil"; "encoding/hex")

const pipe_name = "python_go_pipe"

func receive_from_python(channel chan string)  {
    bytestring, err := ioutil.ReadFile(pipe_name)
    if err != nil {
        fmt.Println("Error in receive_from_python")
        fmt.Println(err)
        return
    }

    from_python := string(bytestring)
    channel <- from_python
}

func receive_from_go(channel chan string)  {
    hashAlgorithm := md5.New()
    to_hash := "This value must be hashed."
    hashAlgorithm.Write([]byte(to_hash))
    hashed := hex.EncodeToString(hashAlgorithm.Sum([]byte{}))
    from_python := <- channel
    fmt.Println("From Python:",from_python)
    fmt.Println("From Go:", hashed)
    fmt.Println("Same: ",from_python == hashed)
}

func main() {
    channel := make(chan string)
    go receive_from_python(channel)
    go receive_from_go(channel)
    var input string
    fmt.Println("Press x to stop.")
    for {
        fmt.Scanln(&input)
        if input == "x" {
            return
        } else {
            fmt.Println("Press x to stop.")
        }
    }
}
