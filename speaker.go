// Range

package main 

import "fmt"

func main() {
	 numeros := []int{2,4,6}
	 suma := 0
	 for _, numero := range numeros {
	 suma += numero
	 }

	 fmt.Printf("suma: %d", suma)

	 for numero := range numeros {
	     fmt.Printf("index: %d ", numero)
	 } 
}

