package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command")
        return
    }

    command := os.Args[1]

    switch command {
    case "greet":
        fmt.Println("Hello, World!")
    case "help":
        fmt.Println("Available commands: greet, help")
    default:
        fmt.Println("Unknown command:", command)
    }
}
