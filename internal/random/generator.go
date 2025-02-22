package random

import "math/rand"

type Generator interface {
	Generate(upperBound int) uint8
}

var _ Generator = (*RealGenerator)(nil)

type RealGenerator struct {
}

func (r *RealGenerator) Generate(upperbound int) uint8 {
	return uint8(rand.Intn(upperbound))
}
