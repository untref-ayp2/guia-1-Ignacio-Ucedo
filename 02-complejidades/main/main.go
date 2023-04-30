package main

import (
	"complejidades/complejidades"
	"fmt"
	"time"
)

func main() {

	n := 30

	inicio := time.Now()

	//ologn
	// complejidades.OlogN(n)
	// fmt.Println("OlogN: ", time.Since(inicio))

	//on
	// complejidades.On(n)
	// fmt.Println("O(n): ", time.Since(inicio))

	// //onlogn
	// complejidades.Onlogn(n)
	// fmt.Println("Onlog(n): ", time.Since(inicio))

	//on2
	// complejidades.On2(n)
	// fmt.Println("O(n)^2: ", time.Since(inicio))

	//on3
	// complejidades.On3(n)
	// fmt.Println("O(n)^3: ", time.Since(inicio))

	//o(2^n)
	// complejidades.O2n(n)
	// fmt.Println("O(2)^n: ", time.Since(inicio))

	//o(n!)
	complejidades.OnFactorial(n)
	fmt.Println("O(n!): ", time.Since(inicio))

}
