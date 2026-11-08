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
		case 4:
			e.TrierCinq()
		case 5:
			e.TrierCinq()
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

func (e *Etat) AmenerEnHaut(i int) {
	n := e.a.Taille()

	nbRa := (n - 1) - i                      
	nbRra := i + 1               

	if nbRa <= nbRra {
		for k := 0; k < nbRa; k++ {    
			e.Executer("ra")           
		}
	} else {
		for k := 0; k < i + 1; k++ {    
			e.Executer("rra")           
		}
	}
}


func (e *Etat) TrierCinq() {
	for e.a.Taille() > 3 {         
		e.AmenerEnHaut(e.a.IndexMin())          
		e.Executer("pb")               
	}
	e.TrierTrois()             

	for e.b.Taille() > 0 {        
		e.Executer("pa")               
	}
}