package random

import "math/rand"

type Generator interface {
	Generate(upperBound int32) uint16
}

var _ Generator = (*RealGenerator)(nil)

type RealGenerator struct {
}

func (r *RealGenerator) Generate(upperBound int32) uint16 {
	return uint16(rand.Int31n(upperBound))
}
