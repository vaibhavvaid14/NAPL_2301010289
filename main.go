package main

import (
    "fmt"
    "time"
)

type Result struct {
    TaskName string
    Value    int
}

func main() {
    results := make(chan Result, 15)

    go func() {
        fmt.Println("[Start] Squares goroutine launched.")
        for i := 1; i <= 5; i++ {
            time.Sleep(50 * time.Millisecond)
            results <- Result{TaskName: "Square", Value: i * i}
        }
        fmt.Println("[End] Squares goroutine completed.")
    }()

    go func() {
        fmt.Println("[Start] Cubes goroutine launched.")
        for i := 1; i <= 5; i++ {
            time.Sleep(40 * time.Millisecond)
            results <- Result{TaskName: "Cube", Value: i * i * i}
        }
        fmt.Println("[End] Cubes goroutine completed.")
    }()

    go func() {
        fmt.Println("[Start] Fibonacci goroutine launched.")
        a, b := 0, 1
        for i := 1; i <= 5; i++ {
            time.Sleep(60 * time.Millisecond)
            results <- Result{TaskName: "Fibonacci", Value: b}
            a, b = b, a+b
        }
        fmt.Println("[End] Fibonacci goroutine completed.")
    }()

    for i := 0; i < 15; i++ {
        res := <-results
        fmt.Printf("  -> Collected: [%s] = %d\n", res.TaskName, res.Value)
    }

    fmt.Println("Main function finished execution.")
}
