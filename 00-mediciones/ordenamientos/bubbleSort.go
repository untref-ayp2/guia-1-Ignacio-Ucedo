package ordenamientos

//aplica el algoritmo de ordenamiento por burbujeo para un arreglo de numeros
func Burbujeo(e []int) []int {

	ordenado := false

	for !ordenado {

		ordenado = true //puede ser mentira, por eso lo chequea abajo
		for i := 0; i < len(e)-1; i++ {
			if e[i] > e[i+1] {
				ordenado = false
				e[i], e[i+1] = e[i+1], e[i]
			}
		}

	}

	return e
}
