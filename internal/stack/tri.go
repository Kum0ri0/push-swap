package stack


func (p *Pile) IndexMax() int {

	indexMax := 0
	for index, v := range p.contenu {
		if v > p.contenu[indexMax] {
			indexMax = index
		}
	}
	return indexMax
}

func (e *Etat) TrierDeux() {
	haut := e.a.contenu[1]
	bas := e.a.contenu[0]

	if haut > bas {
		e.Executer("sa")
	}
}

func (e *Etat) TrierTrois() {
	indexMax := e.a.IndexMax()
	switch indexMax {
		case 2:
			e.Executer("ra")
		case 1:
			e.Executer("rra")
	}
	haut := e.a.contenu[2]
	milieu := e.a.contenu[1]

	if haut > milieu {
		e.Executer("sa")
	}
}

func (e *Etat) Trier() {
	if e.EstResolu() {
		return 
	}
	switch e.a.Taille() {
		case 2:
			e.TrierDeux()
		case 3:
			e.TrierTrois()
	}
}

func (p *Pile) IndexMin() int {

	IndexMin := (len((p.contenu))-1)
	for index, v := range p.contenu {
		if v < p.contenu[IndexMin] {
			IndexMin = index
		}
	}
	return IndexMin
}

func (e *Etat) AmenerEnHaut() {
	

}