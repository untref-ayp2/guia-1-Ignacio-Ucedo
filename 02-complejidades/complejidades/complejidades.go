package complejidades

import (
	"utiles"
)

// funcion de complejidad O(log(n)). Aclaración: la función no calcula el logaritmo de n.
func OlogN(n int) int {

	if n == 0 {
		return 0
	}

	n = n / 2

	return OlogN(n)

}

// función de complejidad O(n).
func On(n int) {
	a := utiles.GenerarArreglo(10, n)

	for i := 0; i < len(a); i++ {
		a[i] = a[i+0] //asignación que no hace nada
	}
}

// función de complejidad Onlog(n)
func Onlogn(n int) {
	x := n
	for n > 1 {
		n = n / 2 //representa las log n veces

		for i := 0; i <= x; i++ { //reperesenta las n veces
			x = x + 0 //asignación qu no hace nada
		}
	}

}

// funcion de complejidad on^2
func On2(n int) {
	f := utiles.GenerarArreglo(n, n)
	c := utiles.GenerarArreglo(n, n)

	for i := 0; i < len(f); i++ {
		f[i] = c[i] // asignación de ejemplo

		for j := 0; j < len(c); j++ {
			c[j] = f[j] // asignación de ejemplo
		}
	}
}

// funcion de complejidad on^3
func On3(n int) {
	i := utiles.GenerarArreglo(n, n)
	j := utiles.GenerarArreglo(n, n)
	k := utiles.GenerarArreglo(n, n)

	for a := 0; a < len(i); a++ {
		i[a] = j[a] // asignación de ejemplo

		for b := 0; b < len(j); b++ {
			j[b] = i[b] // asignación  de ejemplo

			for c := 0; c < len(k); c++ {
				k[c] = i[c] // asignación de ejemplo
			}
		}
	}
}

// funcion de complejidad O(2^n)
func O2n(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return O2n(n-1) + O2n(n-2)
}

// funcion de complejidad O(n!)
func OnFactorial(n int) {

	for i := 0; i < n; i++ {
		OnFactorial(n - 1)
	}

}
