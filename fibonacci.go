package main

import "fmt"

// fibonacci es una función que devuelve
// una función que devuelve un int.
func fibonacci() func() int {
    a:=1
    b:=0
    return func() int {
        c:=a+b
        a=b
        b=c
        return c
    } 
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
