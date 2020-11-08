package main

import (
    "fmt"
    "time"
)

func server(id int){
    for {
        fmt.Printf("Hi from id: %d\n", id)
        time.Sleep(10 * time.Second)
    }
}

func main(){
    go server(1)
    go server(2)
    for{
    }
}
