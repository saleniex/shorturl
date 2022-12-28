package shorturl

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	g := NewSimpleIdGenerator()
	s01 := g.Generate()
	s02 := g.Generate()
	s03 := g.Generate()
	if s01 == s02 || s01 == s03 || s02 == s03 {
		t.Errorf("Collision while generate IDs using SimpleIdGenerator")
	}
}
