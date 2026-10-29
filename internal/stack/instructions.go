package stack

func (p *Pile) Echanger() {
	if len(p.contenu) < 2 {
		return
	}
	p.contenu[(len(p.contenu)-2)], p.contenu[(len(p.contenu)-1)] = p.contenu[(len(p.contenu)-1)], p.contenu[(len(p.contenu)-2)]
}

func (p *Pile) TournerInverse() {
	if len(p.contenu) < 2 {
		return
	}
	fond := p.contenu[0]

	p.contenu = p.contenu[1:]
	p.contenu = append(p.contenu, fond)

}

func (p *Pile) Tourner() {
	if len(p.contenu) < 2 {
		return
	}
	sommet, _ := p.Retirer()

	NouvSlice := []int{sommet}

	NouvSlice = append(NouvSlice, p.contenu...)

	p.contenu = NouvSlice
}

func (e *Etat) Sa() {
	e.a.Echanger()
}

func (e *Etat) Sb() {
	e.b.Echanger()
}

func (e *Etat) Ss() {
	e.Sa()
	e.Sb()
}

func (e *Etat) Pa() {
	result, booleen := e.b.Retirer()
	if booleen == false {
		return
	}
	e.a.Ajouter(result)
}

func (e *Etat) Pb() {
	result, booleen := e.a.Retirer()
	if booleen == false {
		return
	}
	e.b.Ajouter(result)
}

func (e *Etat) Ra() {
	e.a.Tourner()
}

func (e *Etat) Rb() {
	e.b.Tourner()
}

func (e *Etat) Rr() {
	e.Ra()
	e.Rb()
}

func (e *Etat) Rra() {
	e.a.TournerInverse()
}

func (e *Etat) Rrb() {
	e.b.TournerInverse()
}

func (e *Etat) Rrr() {
	e.Rra()
	e.Rrb()
}