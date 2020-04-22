package main

import (
    "fmt"
    "github.com/jasonlvhit/gocron"
)

func task() {
    fmt.Println("I am running task.")
}

func main() {
    // Do jobs without params
    gocron.Every(1).Seconds().Do(task)

    // Start all the pending jobs
    <- gocron.Start()

}