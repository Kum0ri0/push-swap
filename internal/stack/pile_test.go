package stack

import (
	"slices"
	"testing"
)

func TestRetirer(t *testing.T) {
	p := Pile{contenu: []int{5, 8}}
	z := Pile{contenu: []int{}}
	result, deuxieme := p.Retirer()
	result, deuxieme = z.Retirer()

	if !slices.Equal(p.contenu, []int{5}) {
		t.Errorf("nombre retiré attendu 5, obtenu %v", result)
	}

	if result != 8 {
		t.Errorf("nombre retiré attendu 8, obtenu %v", result)
	}
	if deuxieme != true {
		t.Errorf("nombre retiré attendu true, obtenu %v", deuxieme)
	}

}
