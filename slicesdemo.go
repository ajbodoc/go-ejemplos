package main

import (
    "fmt"
)

func main() {
        var nums [] int 
        for i:=0; i<10; i++{
            nums = append(nums,i*10)
        }
    fmt.Println("nums =", nums)
    
    for _, e:= range nums{
         fmt.Println("e =", e)
    }
    
    frutas := [] string {"manzanas","naranjas","ciruelas","pomelos","limas"}
    fmt.Println("frutas =", frutas) 
    
    subs := frutas[0:3] 
    
    for i, fruta := range frutas {
         fmt.Printf("indice = [%d] fruta = %s \n", i, fruta)
    }
    
    for i, fruta := range subs {
         fmt.Printf("[%d] %s \n", i, fruta)
    }
}
