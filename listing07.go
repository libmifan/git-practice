package main

import (
    "fmt"
    "runtime"
    "sync"
)

var wg sync.WaitGroup

func main() {

    runtime.GOMAXPROCS(2)
    wg.Add(2)

    fmt.Println("Starting program...")

    go func() {
        defer wg.Done()

        for count:=0; count<3; count++ {
            for char:='a'; char<'a'+26; char++ {
                fmt.Printf("%c ", char)
            }
        }
    }()

    go func() {
        wg.Done()

        for count:=0; count<3; count++ {
            for char:='A'; char<'A'+26; char++ {
                fmt.Printf("%c ", char)
            }
        }
    }()

    fmt.Println("Waiting for finish")
    wg.Wait()

    fmt.Println("Done!")
}
