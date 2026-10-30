package stack

func (e *Etat) EstResolu() bool {
	if e.b.Taille() != 0 {
		return false
	}
	for i := 0; i < len(e.a.contenu)-1; i++ {
		if e.a.contenu[i] < e.a.contenu[i+1] {
			return false
		}
	}
	return true
}

func (e *Etat) Executer(s string) bool {

	switch s {
	case "sa":
		e.Sa()
	case "sb":
		e.Sb()
	case "pa":
		e.Pa()
	case "pb":
		e.Pb()
	case "ss":
		e.Ss()
	case "ra":
		e.Ra()
	case "rb":
		e.Rb()
	case "rr":
		e.Rr()
	case "rra":
		e.Rra()
	case "rrb":
		e.Rrb()
	case "rrr":
		e.Rrr()
	default:
		return false
	}
	return true
}
