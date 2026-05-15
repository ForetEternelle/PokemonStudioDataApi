package pkmn

import (
	"strconv"
	"strings"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

func NewPokemonQueryFilter(query string, langs ...string) iter2.FilterFunc[PokemonWithForm] {
	return func(pwf PokemonWithForm) bool {
		queryUpper := strings.ToUpper(query)

		p := pwf.Pokemon
		symbolUpper := strings.ToUpper(p.dbSymbol)


		if strings.Contains(symbolUpper, queryUpper) {
			return true
		}

		idStr := strconv.Itoa(int(p.ID()))
		if strings.Contains(idStr, queryUpper) {
			return true
		}

		f, _ := p.Form(pwf.FormId)
		for _, lang := range langs {
			nameUpper := strings.ToUpper(f.Name(lang))
			if strings.Contains(nameUpper, queryUpper) {
				return true
			}
		}
		return false
	}
}
