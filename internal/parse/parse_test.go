package parse

import (
	"slices"
	"testing"
)

func TestParseValide(t *testing.T) {
	cas := []struct {
		args    []string
		attendu []int
	}{
		{[]string{"3 1 2"}, []int{3, 1, 2}},
		{[]string{"3", "1", "2"}, []int{3, 1, 2}},
		{[]string{"  -5   42 +7 "}, []int{-5, 42, 7}},
		{[]string{"2147483647 -2147483648"}, []int{2147483647, -2147483648}},
	}

	for _, c := range cas {
		obtenu, err := Parse(c.args)
		if err != nil {
			t.Errorf("%q : erreur inattendue : %v", c.args, err)
			continue
		}
		if !slices.Equal(obtenu, c.attendu) {
			t.Errorf("%q : attendu %v, obtenu %v", c.args, c.attendu, obtenu)
		}
	}
}

func TestParseErreurs(t *testing.T) {
	invalides := []string{"1 a 3", "1 2.5 3", "1 2 1", "0 -0", "-", "2147483648", "-2147483649"}

	for _, entree := range invalides {
		if _, err := Parse([]string{entree}); err == nil {
			t.Errorf("%q : une erreur était attendue", entree)
		}
	}
}
