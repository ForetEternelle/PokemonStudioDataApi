package studio

// PokemonMapper maps Pokemon descriptors to Pokemon entities.
type PokemonMapper struct {
	store *Store
}

// NewPokemonMapper creates a new PokemonMapper.
func NewPokemonMapper(store *Store) *PokemonMapper {
	return &PokemonMapper{store: store}
}

// MapPokemonDescriptorToPokemon maps a PokemonDescriptor to a Pokemon.
func (m *PokemonMapper) MapPokemonDescriptorToPokemon(desc PokemonDescriptor) *Pokemon {
	forms := make(map[int32]PokemonForm, len(desc.Forms))
	for _, formDesc := range desc.Forms {
		forms[formDesc.Form] = *m.MapFormDescriptorToPokemonForm(formDesc)
	}

	pokemon := NewPokemonBuilder().
		ID(desc.ID).
		DbSymbol(desc.DbSymbol).
		Forms(forms).
		Build()

	return pokemon
}

// MapFormDescriptorToPokemonForm maps a FormDescriptor to a PokemonForm.
func (m *PokemonMapper) MapFormDescriptorToPokemonForm(desc FormDescriptor) *PokemonForm {
	form := NewPokemonFormBuilder().
		Form(desc.Form).
		Type1(m.store.FindTypeBySymbol(desc.Type1)).
		Height(desc.Height).
		Weight(desc.Weight).
		BaseHp(desc.BaseHp).
		BaseAtk(desc.BaseAtk).
		BaseDfe(desc.BaseDfe).
		BaseSpd(desc.BaseSpd).
		BaseAts(desc.BaseAts).
		BaseDfs(desc.BaseDfs).
		EvHp(desc.EvHp).
		EvAtk(desc.EvAtk).
		EvDfe(desc.EvDfe).
		EvSpd(desc.EvSpd).
		EvAts(desc.EvAts).
		EvDfs(desc.EvDfs).
		ExperienceType(ExperienceTypeMap[desc.ExperienceType]).
		BaseExperience(desc.BaseExperience).
		BaseLoyalty(desc.BaseLoyalty).
		CatchRate(desc.CatchRate).
		FemaleRate(desc.FemaleRate).
		BreedGroups(m.MapBreedGroups(desc.BreedGroups)).
		HatchSteps(desc.HatchSteps).
		BabyForm(desc.BabyForm).
		BabyDbSymbol(desc.BabyDbSymbol).
		FrontOffsetY(desc.FrontOffsetY).
		Name(desc.Name).
		Description(desc.Description).
		Evolutions(m.MapEvolutions(desc.Evolutions)).
		ItemHeld(m.MapItemHelds(desc.ItemHeld)).
		CustomProperties(make(map[string]any)).
		Resources(m.MapPokemonResources(desc.Resources)).
		Build()

	if desc.Type2 != nil && *desc.Type2 != "" && *desc.Type2 != UndefType {
		form.type2 = m.store.FindTypeBySymbol(*desc.Type2)
	}

	abilities := make([]*Ability, 0, len(desc.Abilities))
	for _, abilitySymbol := range desc.Abilities {
		if ability := m.store.FindAbilityBySymbol(abilitySymbol); ability != nil {
			abilities = append(abilities, ability)
		}
	}
	form.abilities = abilities

	return form
}

// MapBreedGroups maps breed group descriptors to breed group strings.
func (m *PokemonMapper) MapBreedGroups(breedGroupInts []int32) []string {
	breedGroups := make([]string, len(breedGroupInts))
	for i, bgInt := range breedGroupInts {
		breedGroups[i] = BreedMap[BreedGroupDescriptor(bgInt)]
	}
	return breedGroups
}

// MapEvolutions maps evolution descriptors to evolutions.
func (m *PokemonMapper) MapEvolutions(evolutions []EvolutionDescriptor) []Evolution {
	if len(evolutions) == 0 {
		return nil
	}

	mapped := make([]Evolution, len(evolutions))
	for i, evoDesc := range evolutions {
		mapped[i] = Evolution{
			DbSymbol:   evoDesc.DbSymbol,
			Form:       evoDesc.Form,
			Conditions: m.MapConditions(evoDesc.Conditions),
		}
	}
	return mapped
}

// MapConditions maps condition descriptors to conditions.
func (m *PokemonMapper) MapConditions(conditions []ConditionDescriptor) []Condition {
	if len(conditions) == 0 {
		return nil
	}

	mapped := make([]Condition, len(conditions))
	for i, condDesc := range conditions {
		mapped[i] = Condition{Type: condDesc.Type}
	}
	return mapped
}

// MapItemHelds maps item held descriptors to item held entities.
func (m *PokemonMapper) MapItemHelds(itemHelds []ItemHeldDescriptor) []*ItemHeld {
	if len(itemHelds) == 0 {
		return nil
	}

	mapped := make([]*ItemHeld, len(itemHelds))
	for i, ihDesc := range itemHelds {
		mapped[i] = &ItemHeld{
			DbSymbol: ihDesc.DbSymbol,
			Chance:   ihDesc.Chance,
		}
	}
	return mapped
}

// MapPokemonResources maps Pokemon resources descriptors to Pokemon resources.
func (m *PokemonMapper) MapPokemonResources(resources PokemonResourcesDescriptor) PokemonResources {
	return PokemonResources{
		Icon:            resources.Icon,
		IconF:           resources.IconF,
		IconShiny:       resources.IconShiny,
		IconShinyF:      resources.IconShinyF,
		Front:           resources.Front,
		FrontF:          resources.FrontF,
		FrontShiny:      resources.FrontShiny,
		FrontShinyF:     resources.FrontShinyF,
		Back:            resources.Back,
		BackF:           resources.BackF,
		BackShiny:       resources.BackShiny,
		BackShinyF:      resources.BackShinyF,
		Footprint:       resources.Footprint,
		Character:       resources.Character,
		CharacterF:      resources.CharacterF,
		CharacterShiny:  resources.CharacterShiny,
		CharacterShinyF: resources.CharacterShinyF,
		Cry:             resources.Cry,
		HasFemale:       resources.HasFemale,
		Egg:             resources.Egg,
		IconEgg:         resources.IconEgg,
	}
}
