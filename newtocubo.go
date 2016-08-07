package main

import (
    "fmt"
    "math/cmplx"
)


func Cbrt(x complex128, n int) complex128 {
    var z complex128 = 1
    
    if n==0 {n=1000}
    
    for i:=0; i<n; i++ {
       z =  z - ( (z * z * z) - x ) / ((z * z) * 3)    
    }
    
    return z
}

func main() {
    fmt.Println(Cbrt(2,10))
    var p complex128 = 1.0 / 3.0
    fmt.Println(cmplx.Pow(2.0,p))
}
