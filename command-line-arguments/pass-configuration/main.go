package main

import (
    "fmt"
    "os"
)

func main() {
    debug := false

    for _, arg := range os.Args {
        if arg == "--debug" {
            debug = true
        }
    }

    if debug {
        fmt.Println("Debug mode enabled")
    } else {
        fmt.Println("Debug mode disabled")
    }
}
