package subsumamax

// Dado un arreglo devuelve la posicion inicial, la posición final y el valor
// de la subsecuencia cuya suma es máxima dentro del arreglo original. Misma función pero
//pretende de ser O(n)
func SubSumaMaxOdeN(arreglo []int) (int, int, int) {
	sumaMaxima := -1
	posInicial, posFinal := -1, -1
	primerPositivo := -1
	posibleSuma := 0
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] >= 0 {
			if primerPositivo == -1 {

				sumaMaxima = arreglo[i]
				posInicial = i
				posFinal = i

				if arreglo[i]+arreglo[i+1] >= 0 {
					primerPositivo = arreglo[i]
				}
			} else if posibleSuma == 0 {
				sumaMaxima += arreglo[i]
				posFinal = i
			} else {
				posibleSuma += arreglo[i]
				if posibleSuma >= 0 {
					sumaMaxima += posibleSuma
					posibleSuma = 0
					posFinal = i
				}
			}
		} else if primerPositivo != -1 {
			posibleSuma += arreglo[i]
		}

	}
	return sumaMaxima, posInicial, posFinal
}

// Lo que hago es buscar el primer positivo que me encuentre (el cero lo tomo como positivo).
// cuando lo encuentro guardo lo guardo en la sumaMaxima.
// Si se piensa que la suma máxima podría llegar a tener
// más de un único valor, este primer positivo también debe cumplir que su suma con el subsiguiente sea positiva.
// Sino sigo buscando el llamado primerPositivo (aunque igualmente ya tengo guardado el primer positivo
// verdadero en la sumaMaxima por si llega a ser el único positivo de todo el arreglo).

// Una vez encontrado el primerPositvo, desde allí pregunto, el subsiguiente, ¿Es positivo? Si es así, lo meto en la subsecuencia.
// El tercero, ¿Lo es? si es lo meto también. Si no es positivo lo meto en otra variable llamada posibleSuma.
// Voy metiendo los siguientes valores del arreglo en posibleSuma hasta encontrarme con un positivo.
// Si al meter el positivo en posibleSuma, este la vuelve positiva, sumo a posibleSuma en la subsecuencia
// y a posibleSuma le guardo el cero (para reiniciarla).
// Si posibleSuma nunca logra ser positiva, la subsecuencia termina sin modificaciones, siendo únicamente esos primeros n números positivos.
