package main

import (
	"fmt"
)

func main() {
        deps:= make(map[string]int)
 
        deps["desarrolladores"]=10
        deps["administración"]=5
        deps["ventas"]=15
        deps["mantenimiento"]=8
        deps["jefes"]=3

	fmt.Println("departamentos =", deps)
	
	for key, value:= range deps{
	   fmt.Printf("departamento: %s personas: %d \n", key, value)
	}
	
	nums := []int {2,4,6,8}
	sum:=0
	for _, num:= range nums{
	   sum+=num
	   fmt.Printf("num= %d \n", num)
	}
	fmt.Printf("suma =%d \n", sum)
	
}