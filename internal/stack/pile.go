package stack

type Pile struct {
	contenu []int
}

func (p *Pile) Taille() int {
	result := len(p.contenu)
	return result
}

func (p *Pile) Ajouter(nombre int) []int {
	p.contenu = append(p.contenu, nombre)
	return p.contenu
}

func (p *Pile) Retirer() (int, bool) {
	if p.Taille() == 0 {
		return 0, false
	}
	derniern := (len(p.contenu) - 1)
	nombre := p.contenu[(len(p.contenu) - 1)]
	p.contenu = p.contenu[0:derniern]

	return nombre, true
}
