package stack

type Etat struct {
	a Pile
	b Pile
}

func NewEtat(nombres []int) *Etat {
	etat := Etat{}

	for i := (len(nombres) - 1); i >= 0; i-- {
		etat.a.contenu = append(etat.a.contenu, nombres[i])
	}
	return &etat
}
