package subsumamax_test

import (
	"subsumamax"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsecuenciaSumaMaxima(t *testing.T) {

	e := []int{-1, 6, -2, 5, -1, 4, 3, -4, 3}

	sumaMaxima, posInicial, posFinal := subsumamax.SubSumaMaxOdeN(e)

	assert.Equal(t, 15, sumaMaxima)
	assert.Equal(t, 1, posInicial)
	assert.Equal(t, 6, posFinal)
}

func TestSubsecuenciaSumaMaxima2(t *testing.T) {

	e := []int{3, -4, 3, 4, -1, 5, -2, 6, -1}

	sumaMaxima, posInicial, posFinal := subsumamax.SubSumaMaxOdeN(e)

	assert.Equal(t, 15, sumaMaxima)
	assert.Equal(t, 2, posInicial)
	assert.Equal(t, 7, posFinal)
}

func TestSubsecuenciaSumaMaxima3(t *testing.T) {

	e := []int{-3, -4, -3, -4, -1, -5, -2, -6, -1}

	sumaMaxima, posInicial, posFinal := subsumamax.SubSumaMaxOdeN(e)

	assert.Equal(t, -1, sumaMaxima)
	assert.Equal(t, -1, posInicial)
	assert.Equal(t, -1, posFinal)
}

func TestSubsecuenciaSumaMaxima4(t *testing.T) {

	e := []int{-3, -4, -3, 4, -8, -5, -2, -6, -1}

	sumaMaxima, posInicial, posFinal := subsumamax.SubSumaMaxOdeN(e)

	assert.Equal(t, 4, sumaMaxima)
	assert.Equal(t, 3, posInicial)
	assert.Equal(t, 3, posFinal)
}

func TestSubsecuenciaSumaMaxima5(t *testing.T) {

	e := []int{-3, -4, -3, 4, -1, 5, -2, -6, -1}

	sumaMaxima, posInicial, posFinal := subsumamax.SubSumaMaxOdeN(e)

	assert.Equal(t, 8, sumaMaxima)
	assert.Equal(t, 3, posInicial)
	assert.Equal(t, 5, posFinal)

}
