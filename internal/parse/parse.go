package parse

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(s []string) ([]int, error) {
	carte := map[int]bool{}
	text := []string{}

	for _, v := range s {
		mot := strings.Fields(v)
		text = append(text, mot...)
	}

	SliceNbr := []int{}

	for _, r := range text {
		nbr, err := strconv.ParseInt(r, 10, 32)
		if err != nil {
			return nil, err
		}
		if carte[int(nbr)] {
			return nil, errors.New("il ne peut pas y avoir de doublon")
		}
		carte[int(nbr)] = true
		SliceNbr = append(SliceNbr, int(nbr))
	}
	return SliceNbr, nil
}
