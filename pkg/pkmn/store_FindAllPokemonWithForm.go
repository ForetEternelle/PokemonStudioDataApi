package pkmn

import (
	"iter"
	"log/slog"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

type PokemonWithForm struct {
	Pokemon *Pokemon
	FormId  int32
}

type findAllWithFormOptions struct {
	PokemonFilter iter2.FilterFunc[*Pokemon]
	FormFilter    iter2.FilterFunc[*PokemonForm]
	LastId        *int32
	LastForm      *int32
	MainFormOnly  bool
}
type FindAllPokemonWithFormOption func(*findAllWithFormOptions)

func newFindAllWithFormOptions(opts ...FindAllPokemonWithFormOption) *findAllWithFormOptions {
	res := &findAllWithFormOptions{
		PokemonFilter: iter2.True[*Pokemon],
		FormFilter:    iter2.True[*PokemonForm],
		MainFormOnly:  false,
	}

	for _, opt := range opts {
		opt(res)
	}
	return res
}

var WithPokemonFilter = func(filter iter2.FilterFunc[*Pokemon]) func(*findAllWithFormOptions) {
	return func(opts *findAllWithFormOptions) {
		opts.PokemonFilter = filter
	}
}

var WithFormFilter = func(filter iter2.FilterFunc[*PokemonForm]) func(*findAllWithFormOptions) {
	return func(opts *findAllWithFormOptions) {
		opts.FormFilter = filter
	}
}

var WithLastId = func(id int32) func(*findAllWithFormOptions) {
	return func(opts *findAllWithFormOptions) {
		opts.LastId = &id
	}
}

var WithLastForm = func(formId int32) func(*findAllWithFormOptions) {
	return func(opts *findAllWithFormOptions) {
		opts.LastForm = &formId
	}
}

var WithMainFormsOnly = func() func(*findAllWithFormOptions) {
	return func(opts *findAllWithFormOptions) {
		opts.MainFormOnly = true
	}
}

func (s *Store) FindAllPokemonWithForm(opts ...FindAllPokemonWithFormOption) iter.Seq[PokemonWithForm] {
	options := *newFindAllWithFormOptions(opts...)

	startAt := 0
	if options.LastId != nil {
		var ok bool
		startAt, ok = slices.BinarySearchFunc(s.pokemonList, *options.LastId, func(p *Pokemon, id int32) int {
			return int(p.ID) - int(id)
		})
		if ok {
			startAt++
			slog.Debug("Last ID found in pokemon list, starting from next index", "lastId", *options.LastId, "startAt", startAt)
		}
	}

	return func(yield func(PokemonWithForm) bool) {
		for _, pokemon := range s.pokemonList[startAt:] {
			if !options.PokemonFilter(pokemon) {
				continue
			}

			if options.MainFormOnly {
				mainForm, _ := pokemon.Form(0)
				if mainForm == nil || !options.FormFilter(mainForm) {
					continue
				}

				if !yield(PokemonWithForm{
					Pokemon: pokemon,
					FormId:  0,
				}) {
					return
				}
				continue
			}

			formIt := slices.Values(pokemon.Forms)
			formIt = iter2.Filter(formIt, options.FormFilter)
			for form := range formIt {
				if !yield(PokemonWithForm{
					Pokemon: pokemon,
					FormId:  form.Form,
				}) {
					return
				}
			}
		}
	}

}
