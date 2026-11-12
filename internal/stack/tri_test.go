package stack

import (
	"math/rand"
	"slices"
	"testing"
)

func TestExempleSujet(t *testing.T) {
	etat := NewEtat([]int{2, 1, 3, 6, 8, 5})
	instructions := []string{"sa", "pb", "pb", "pb", "rb", "rrr", "pa", "pa", "pa"}

	for _, ins := range instructions {
		if !etat.Executer(ins) {
			t.Fatalf("l'instruction %q a été refusée", ins)
		}
	}
	if !etat.EstResolu() {
		t.Errorf("l'exemple du sujet devrait donner une pile triée")
	}
}

func TestInstructionInconnue(t *testing.T) {
	etat := NewEtat([]int{2, 1, 3})

	for _, ins := range []string{"hello", "SA", "sa ", "s a", ""} {
		if etat.Executer(ins) {
			t.Errorf("l'instruction %q aurait dû être refusée", ins)
		}
	}
}

func verifierTri(t *testing.T, nombres []int) {
	etat := NewEtat(nombres)
	etat.Trier()
	if !etat.EstResolu() {
		t.Errorf("%v : la pile n'est pas triée après Trier", nombres)
	}

	verif := NewEtat(nombres)
	for _, coup := range Nettoyer(etat.Coups()) {
		verif.Executer(coup)
	}
	if !verif.EstResolu() {
		t.Errorf("%v : les coups nettoyés ne trient pas la pile", nombres)
	}
}
func TestTrierCasLimites(t *testing.T) {
	cas := [][]int{
		{42},
		{2, 1},
		{1, 2},
		{3, 2, 1},
		{2, 3, 1},
		{4, 67, 3, 87, 23},
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7, 1, 2},
		{10, -5, 0, 2147483647, -2147483648, 7, 3, 99, -1, 42},
	}
	for _, nombres := range cas {
		verifierTri(t, nombres)
	}
}

func TestTrierAleatoire(t *testing.T) {
	for i := 0; i < 20; i++ {
		verifierTri(t, rand.Perm(100))
	}
}

func TestNettoyer(t *testing.T) {
	entree := []string{"ra", "pa", "pa", "pb", "pb", "ra", "rra", "pb"}
	attendu := []string{"ra", "pb"}

	obtenu := Nettoyer(entree)
	if !slices.Equal(obtenu, attendu) {
		t.Errorf("attendu %v, obtenu %v", attendu, obtenu)
	}
}
