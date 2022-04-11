package main

import (
    "github.com/arcs/pastee/handlers"
)

func main() {
    s := handlers.New()
    s.Router().Run(":8888")
}
