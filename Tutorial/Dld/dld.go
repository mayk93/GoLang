package main

import ("fmt"; "encoding/json"; "io/ioutil"; "github.com/hashicorp/go-getter"; "strings"; "math"; "math/rand")

const pipe_name = "python_go_pipe"

type UrlList []struct {
	URL string `json:"url"`
}

func receiveFromPython() []byte {
    bytestring, err := ioutil.ReadFile(pipe_name)
    if err != nil {
        fmt.Println("Error in receive_from_python")
        fmt.Println(err)
        return make([]byte, 0)
    }

    return bytestring
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


func getBase(url string) string {
    var base string
    if strings.HasPrefix(url, "https://www.") {
        base = url[12:]
    } else if strings.HasPrefix(url, "http://www.") {
        base = url[11:]
    } else if strings.HasPrefix(url, "https://") {
        base = url[8:]
    } else if strings.HasPrefix(url, "http://") {
        base = url[7:]
    } else {
        fmt.Println("Process part 1 failed for url", url)
    }

    dotIndex := strings.Index(base, ".")

    base = base[:dotIndex]
    base += "_" + getRandomString(rand.Intn(1000))

    return base
}

func download(url string) {

    fmt.Println("Downloading", url, "to", "downloads/" + getBase(url))

    dld := new(getter.Client)
    mode := getter.ClientModeAny

    dld.Src = url
    dld.Dst = "downloads/" + getBase(url)
    dld.Mode = mode

    err := dld.Get()
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
    jsonByteString := receiveFromPython()
    var urlList UrlList
    err := json.Unmarshal(jsonByteString, &urlList)

    if err != nil {
        fmt.Println(err)
        return
    }

    for _, url := range urlList {
        fmt.Println("Sending to download:",url.URL)
        go download(url.URL)
    }

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
