package main

import (
    "fmt"
    "time"
)

func main() {
    stop := make(chan bool)

    go func() {
        for {
            select {
            case <-stop:
                fmt.Println("goroutine 1 stopping")
                return
            default:
                fmt.Println("goroutine 1 says hi")
                time.Sleep(3 * time.Second)
            }
        }
    }()

    go func() {
        time.Sleep(10 * time.Second) // big work
        stop <- true
    }()

    time.Sleep(12 * time.Second)
}
