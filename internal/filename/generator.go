package filename

import (
	"math/rand"
	"strconv"
)

type Generator interface {
	Generate() string
}

var generator Generator = (*RealGenerator)(nil)

type RealGenerator struct {
}

func (r *RealGenerator) Generate() string {
	return "/" + strconv.Itoa(rand.Int()) + ".png"
}
