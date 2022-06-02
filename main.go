package main

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"

    "github.com/arcs/pastee/handlers"
    "github.com/arcs/pastee/pkg/clean"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Server must be called with exactly 1 argument. One of [server, cleanup]")
        os.Exit(1)
    }

    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Could not load env: %v\n", err)
        os.Exit(2)
    }

    runMode := os.Args[1]
    if runMode == "server" {
        s := handlers.New()
        s.Router().Run(":8888")
    }
    if runMode == "cleanup" {
        c := clean.New()
        err := c.PerformCleanup()
        if err != nil {
            fmt.Printf("Error performing cleanup: %v\n", err)
            os.Exit(3)
        }
    }
}
