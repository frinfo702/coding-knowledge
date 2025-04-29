package main

import (
    "log"
)

func main() {
    srv := NewServer(WithOnError(func(err error) { log.Printf("ERR: %v", err) }))
    if err := srv.Listen(":8080"); err != nil {
        log.Fatal(err)
    }
}
