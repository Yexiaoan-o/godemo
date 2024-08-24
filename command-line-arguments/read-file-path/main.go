package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a file path")
        return
    }

    filePath := os.Args[1]
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println("File content:")
    fmt.Println(string(content))
}
