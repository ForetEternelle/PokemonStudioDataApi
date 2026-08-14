package pkmn

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

// NewPokemonQueryFilter returns a filter function that checks if a PokemonWithForm matches the given query string in its ID or name in any of the specified languages.
func NewPokemonQueryFilter(query string, langs ...string) iter2.FilterFunc[PokemonWithForm] {
	return func(pwf PokemonWithForm) bool {
		queryUpper := strings.ToUpper(query)

		p := pwf.Pokemon

		idStr := strconv.Itoa(int(p.ID()))
		if strings.Contains(idStr, query) {
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

// NewPokemonFormTypesFilter returns a filter function that checks if a PokemonForm matches any of the given types.
func NewPokemonFormTypesFilter(types []string) iter2.FilterFunc[*PokemonForm] {
	if len(types) == 0 {
		return iter2.True
	}

	return func(pf *PokemonForm) bool {
		for _, t := range types {
			if pf.type1.dbSymbol == t ||
				(pf.type2 != nil && pf.type2.dbSymbol == t) {
				return true
			}
		}
		return false
	}
}

// PokemonWithFormTags returns the tags of a PokemonWithForm, from both the pokemon and its form.
func PokemonWithFormTags(pwf PokemonWithForm) []string {
	tags := pwf.Pokemon.Tags
	if f, ok := pwf.Pokemon.Form(pwf.FormId); ok {
		tags = append(tags, f.Tags...)
	}
	return tags
}

// NewPokemonWithTagsFilter returns a filter function that checks if a PokemonWithForm
// matches any of the given tags (on the pokemon or its form).
func NewPokemonWithTagsFilter(tags []string) iter2.FilterFunc[PokemonWithForm] {
	if len(tags) == 0 {
		return iter2.True
	}

	return func(pwf PokemonWithForm) bool {
		for _, tag := range PokemonWithFormTags(pwf) {
			if slices.Contains(tags, tag) {
				return true
			}
		}
		return false
	}
}

// NewPokemonWithoutTagsFilter returns a filter function that checks if a PokemonWithForm
// has none of the given tags (on the pokemon or its form).
func NewPokemonWithoutTagsFilter(tags []string) iter2.FilterFunc[PokemonWithForm] {
	if len(tags) == 0 {
		return iter2.True
	}

	return func(pwf PokemonWithForm) bool {
		for _, tag := range PokemonWithFormTags(pwf) {
			if slices.Contains(tags, tag) {
				return false
			}
		}
		return true
	}
}
