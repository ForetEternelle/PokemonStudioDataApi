package studio

var ExperienceTypeMap = map[ExperienceTypeDescriptor]string{
	ExperienceErraticNum:     ExperienceErratic,
	ExperienceFastNum:        ExperienceFast,
	ExperienceMediumFastNum:  ExperienceMediumFast,
	ExperienceMediumSlowNum:  ExperienceMediumSlow,
	ExperienceSlowNum:        ExperienceSlow,
	ExperienceFluctuatingNum: ExperienceFluctuating,
}

var BreedMap = map[BreedGroupDescriptor]string{
	BreedMonsterNum:      BreedMonster,
	BreedWater1Num:       BreedWater1,
	BreedBugNum:          BreedBug,
	BreedFlyingNum:       BreedFlying,
	BreedFieldNum:        BreedField,
	BreedFairyNum:        BreedFairy,
	BreedGrassNum:        BreedGrass,
	BreedHumanNum:        BreedHuman,
	BreedWater3Num:       BreedWater3,
	BreedMineralNum:      BreedMineral,
	BreedAmorphousNum:    BreedAmorphous,
	BreedWater2Num:       BreedWater2,
	BreedDittoNum:        BreedDitto,
	BreedDragonNum:       BreedDragon,
	BreedUndiscoveredNum: BreedUndiscovered,
}

// PokemonMapper handles mapping of Pokemon descriptors with reference resolution using a store
type PokemonMapper struct {
	store *Store
}

// NewPokemonMapper creates a new PokemonMapper with the given store
func NewPokemonMapper(store *Store) *PokemonMapper {
	return &PokemonMapper{store: store}
}

// MapPokemonDescriptorToPokemon converts a PokemonDescriptor to a Pokemon domain struct with reference resolution
func (m *PokemonMapper) MapPokemonDescriptorToPokemon(desc PokemonDescriptor) *Pokemon {
	forms := make([]PokemonForm, len(desc.Forms))
	for i, formDesc := range desc.Forms {
		forms[i] = *m.MapFormDescriptorToPokemonForm(formDesc)
	}

	pokemon := &Pokemon{
		Id:               desc.ID,
		DbSymbol:         desc.DbSymbol,
		Forms:            forms,
		CustomProperties: make(map[string]any),
	}

	return pokemon
}

// MapFormDescriptorToPokemonForm converts a FormDescriptor to a PokemonForm domain struct with reference resolution
func (m *PokemonMapper) MapFormDescriptorToPokemonForm(desc FormDescriptor) *PokemonForm {
	form := &PokemonForm{
		Name:           desc.Name,
		Description:    desc.Description,
		Form:           desc.Form,
		Height:         desc.Height,
		Weight:         desc.Weight,
		BaseHp:         desc.BaseHp,
		BaseAtk:        desc.BaseAtk,
		BaseDfe:        desc.BaseDfe,
		BaseSpd:        desc.BaseSpd,
		BaseAts:        desc.BaseAts,
		BaseDfs:        desc.BaseDfs,
		EvHp:           desc.EvHp,
		EvAtk:          desc.EvAtk,
		EvDfe:          desc.EvDfe,
		EvSpd:          desc.EvSpd,
		EvAts:          desc.EvAts,
		EvDfs:          desc.EvDfs,
		ExperienceType: ExperienceTypeMap[desc.ExperienceType],
		BaseExperience: desc.BaseExperience,
		BaseLoyalty:    desc.BaseLoyalty,
		CatchRate:      desc.CatchRate,
		FemaleRate:     desc.FemaleRate,
		BreedGroups:    MapBreedGroups(desc.BreedGroups),
		HatchSteps:     desc.HatchSteps,
		BabyForm:       desc.BabyForm,
		FrontOffsetY:   desc.FrontOffsetY,
	}

	if desc.BabyDbSymbol != nil {
		form.BabyDbSymbol = desc.BabyDbSymbol
	}

	form.Evolutions = MapEvolutions(desc.Evolutions)
	form.ItemHeld = MapItemHelds(desc.ItemHeld)

	form.Type1 = m.store.FindTypeBySymbol(desc.Type1)
	if desc.Type2 != nil && *desc.Type2 != "" && *desc.Type2 != UndefType {
		form.Type2 = m.store.FindTypeBySymbol(*desc.Type2)
	}

	form.Abilities = make([]*Ability, 0, len(desc.Abilities))
	for _, abilitySymbol := range desc.Abilities {
		if ability := m.store.FindAbilityBySymbol(abilitySymbol); ability != nil {
			form.Abilities = append(form.Abilities, ability)
		}
	}

	return form
}

// TypeMapper handles mapping of PokemonType descriptors using a store
type TypeMapper struct {
	store *Store
}

// NewTypeMapper creates a new TypeMapper with the given store
func NewTypeMapper(store *Store) *TypeMapper {
	return &TypeMapper{store: store}
}

// MapPokemonTypeDescriptorToPokemonType converts a PokemonTypeDescriptor to a PokemonType domain struct
func (m *TypeMapper) MapPokemonTypeDescriptorToPokemonType(desc PokemonTypeDescriptor) *PokemonType {
	pokemonType := &PokemonType{
		DbSymbol: desc.DbSymbol,
		Color:    desc.Color,
		TextId:   desc.TextId,
		Name:     desc.Name,
	}

	pokemonType.DamageTo = MapTypeDamages(desc.DamageTo)
	return pokemonType
}

// AbilityMapper handles mapping of Ability descriptors using a store
type AbilityMapper struct {
	store *Store
}

// NewAbilityMapper creates a new AbilityMapper with the given store
func NewAbilityMapper(store *Store) *AbilityMapper {
	return &AbilityMapper{store: store}
}

// MapAbilityDescriptorToAbility converts an AbilityDescriptor to an Ability domain struct
func (m *AbilityMapper) MapAbilityDescriptorToAbility(desc AbilityDescriptor) *Ability {
	ability := &Ability{
		DbSymbol:    desc.DbSymbol,
		Id:          desc.Id,
		TextID:      desc.TextID,
		Description: desc.Description,
		Name:        desc.Name,
	}

	return ability
}

func MapBreedGroups(breedGroupInts []int32) []string {
	breedGroups := make([]string, len(breedGroupInts))
	for i, bgInt := range breedGroupInts {
		breedGroups[i] = BreedMap[BreedGroupDescriptor(bgInt)]
	}
	return breedGroups
}

func MapEvolutions(evolutions []EvolutionDescriptor) []Evolution {
	if len(evolutions) == 0 {
		return nil
	}

	mapped := make([]Evolution, len(evolutions))
	for i, evoDesc := range evolutions {
		mapped[i] = Evolution{
			DbSymbol:   evoDesc.DbSymbol,
			Form:       evoDesc.Form,
			Conditions: MapConditions(evoDesc.Conditions),
		}
	}
	return mapped
}

func MapConditions(conditions []ConditionDescriptor) []Condition {
	if len(conditions) == 0 {
		return nil
	}

	mapped := make([]Condition, len(conditions))
	for i, condDesc := range conditions {
		mapped[i] = Condition{Type: condDesc.Type}
	}
	return mapped
}

func MapItemHelds(itemHelds []ItemHeldDescriptor) []*ItemHeld {
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

func MapTypeDamages(damages []TypeDamageDescriptor) []TypeDamage {
	if len(damages) == 0 {
		return nil
	}

	mapped := make([]TypeDamage, len(damages))
	for i, tdDesc := range damages {
		mapped[i] = TypeDamage{
			DefensiveType: tdDesc.DefensiveType,
			Factor:        tdDesc.Factor,
		}
	}
	return mapped
}
