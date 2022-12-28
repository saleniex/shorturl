package shorturl

import (
	"math/rand"
	"strings"
	"time"
)

// SimpleIdGenerator generates simple identifier without collision prevention
type SimpleIdGenerator struct {
	Len int
}

var materia = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

func NewSimpleIdGenerator() SimpleIdGenerator {
	return SimpleIdGenerator{
		Len: 8,
	}
}

func (g *SimpleIdGenerator) Generate() string {
	tokens := strings.Split(materia, "")
	rand.Seed(time.Now().UnixMicro())
	rand.Shuffle(len(tokens), func(i, j int) {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	})
	shufledMateria := strings.Join(tokens, "")

	return shufledMateria[0:g.Len]
}
