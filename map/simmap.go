package simmap

import "math/rand/v2"

type Map struct {
	Size int
}

func NewMap(w int) *Map { return &Map{Size: w} }

func RandomPosition(m *Map) [2]int {
	w := rand.IntN(m.Size)
	h := rand.IntN(m.Size)
	return [2]int{w,h}
}