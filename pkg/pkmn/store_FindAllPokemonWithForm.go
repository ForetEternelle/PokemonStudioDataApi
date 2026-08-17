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
	LastForm      int32
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
		opts.LastForm = formId
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

		if !ok {
			startAt = 0
			slog.Debug("Last ID not found in pokemon list, starting from index 0", "lastId", *options.LastId)
		} else if options.MainFormOnly {
			startAt++
		}
	}

	var pwfIt iter.Seq[PokemonWithForm] = nil
	pkmnIt := slices.Values(s.pokemonList[startAt:])
	pkmnIt = iter2.Filter(pkmnIt, options.PokemonFilter)

	if options.MainFormOnly {
		pwfIt = pokemonWithMainFormIt(pkmnIt)
	} else {
		pwfIt = pokemonWithAllFormIt(pkmnIt)
		if options.LastId != nil {
			pwfIt = iter2.SkipUntil(pwfIt, func(pwf PokemonWithForm) bool {
				return pwf.Pokemon.ID != *options.LastId || pwf.FormId > options.LastForm
			})
		}
	}

	return iter2.Filter(pwfIt, func(pwf PokemonWithForm) bool {
		form, ok := pwf.Pokemon.Form(pwf.FormId)
		return ok && options.FormFilter(form)
	})
}

func pokemonWithMainFormIt(pokemonIt iter.Seq[*Pokemon]) iter.Seq[PokemonWithForm] {
	return iter2.Map(pokemonIt, func(pokemon *Pokemon) PokemonWithForm {
		return PokemonWithForm{
			Pokemon: pokemon,
			FormId:  0,
		}
	})
}

func pokemonWithAllFormIt(pokemonIt iter.Seq[*Pokemon]) iter.Seq[PokemonWithForm] {
	return func(yield func(PokemonWithForm) bool) {
		for pokemon := range pokemonIt {
			formIt := slices.Values(pokemon.Forms)
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
