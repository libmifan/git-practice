package main

import (
    "fmt"
    "os"
)

type error interface {
    Error() string
}

type RPCError struct {
    Code    int64
    Message string
}

func (e *RPCError) Error() string {
    return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func main() {
	fmt.Println("vim-go")
    fmt.Println("ars: ", os.Args[1:])
    for idx,v := range os.Args[1:] {
        fmt.Printf("Args[%d] = %s\n", idx, v)
    }
}
