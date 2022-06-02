package main

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"

    "github.com/arcs/pastee/handlers"
    "github.com/arcs/pastee/pkg/clean"
)

const CmdServer = "server"
const CmdCleanup = "cleanup"

func main() {
    if len(os.Args) != 2 || (os.Args[1] != CmdServer && os.Args[1] != CmdCleanup) {
        fmt.Printf("Server must be called with exactly 1 argument. One of [%s, %s]\n", CmdServer, CmdCleanup)
        os.Exit(1)
    }

    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Could not load env: %v\n", err)
        os.Exit(2)
    }

    runMode := os.Args[1]
    if runMode == CmdServer {
        s := handlers.New()
        s.Router().Run(":8888")
    }
    if runMode == CmdCleanup {
        c := clean.New()
        err := c.PerformCleanup()
        if err != nil {
            fmt.Printf("Error performing cleanup: %v\n", err)
            os.Exit(3)
        }
    }
}
