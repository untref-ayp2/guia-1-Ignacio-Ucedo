package ordenamientos_test

import (
	"ordenamientos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble1(t *testing.T) {

	e := []int{1, 5, 2, 90, 2, 9}

	assert.Equal(t, []int{1, 2, 2, 5, 9, 90}, ordenamientos.Burbujeo(e))
}

func TestBubble2(t *testing.T) {

	e := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, ordenamientos.Burbujeo(e))
}

func TestBubble3(t *testing.T) {

	e := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 88, 2, -5}

	assert.Equal(t, []int{-5, 0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 88}, ordenamientos.Burbujeo(e))
}
