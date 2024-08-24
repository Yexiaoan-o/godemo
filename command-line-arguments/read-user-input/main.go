package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Please provide two numbers")
        return
    }

    num1, err1 := strconv.Atoi(os.Args[1])
    num2, err2 := strconv.Atoi(os.Args[2])

    if err1 != nil || err2 != nil {
        fmt.Println("Invalid numbers")
        return
    }

    sum := num1 + num2
    fmt.Println("Sum:", sum)
}
