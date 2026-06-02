package studio

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"iter"
	"log/slog"
	"os"
	"path"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/file"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

const (
	pokemonNameTranslationFile            = "100000.csv"
	pokemonDescriptionTranslationFile     = "100002.csv"
	pokemonFormNameTranslationFile        = "100067.csv"
	pokemonFormDescriptionTranslationFile = "100068.csv"
	typeNameTranslationFile               = "100003.csv"
	abilityNameTranslationFile            = "100004.csv"
	abilityDescriptionTranslationFile     = "100005.csv"
	moveNameTranslationFile               = "100006.csv"
	moveDescriptionTranslationFile        = "100007.csv"
)

func Load(folder string) (*pkmn.Store, error) {
	store := pkmn.NewStore()
	translationFolder := path.Join(folder, pkmn.LanguageFolder)
	studioFolder := path.Join(folder, pkmn.StudioFolder)

	typeIterator, err := ImportTypes(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load pokemon types")
		return nil, err
	}
	for descriptor := range typeIterator {
		store.AddType(pkmn.AddTypeDto{
			DbSymbol: descriptor.DbSymbol,
			Color:    descriptor.Color,
			TextId:   descriptor.TextId,
			Name:     descriptor.Name,
			DamageTo: mapTypeDamages(descriptor.DamageTo),
		})
	}

	abilityIterator, err := ImportAbility(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load abilities")
		return nil, err
	}
	for descriptor := range abilityIterator {
		store.AddAbility(pkmn.AddAbilityDto{
			DbSymbol:    descriptor.DbSymbol,
			Id:          descriptor.Id,
			TextId:      descriptor.TextID,
			Name:        descriptor.Name,
			Description: descriptor.Description,
		})
	}

	moveIterator, err := ImportMoves(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load moves")
		return nil, err
	}
	for descriptor := range moveIterator {
		store.AddMove(pkmn.AddMoveDto{
			Id:               descriptor.Id,
			DbSymbol:         descriptor.DbSymbol,
			Type:             descriptor.Type,
			Category:         pkmn.MoveCategory(descriptor.Category),
			Power:            descriptor.Power,
			Accuracy:         descriptor.Accuracy,
			PP:               descriptor.PP,
			CriticalRate:     descriptor.MoveCriticalRate,
			Priority:         descriptor.Priority,
			MapUse:           descriptor.MapUse,
			Targeting:        mapTargeting(*descriptor),
			Execution:        mapExecution(*descriptor),
			MechanicalTags:   mapMechanicalTags(*descriptor),
			Interactions:     mapInteractions(*descriptor),
			SecondaryEffects: mapSecondaryEffects(*descriptor),
			Name:             descriptor.Name,
			Description:      descriptor.Description,
		})
	}

	pokemonIterator, err := ImportPokemon(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load pokemon")
		return nil, err
	}
	for descriptor := range pokemonIterator {
		forms := make([]pkmn.AddPokemonFormDto, len(descriptor.Forms))
		for i, f := range descriptor.Forms {
			abilities := make([]string, len(f.Abilities))
			copy(abilities, f.Abilities)

			forms[i] = pkmn.AddPokemonFormDto{
				Form:             f.Form,
				Type1:            f.Type1,
				Type2:            f.Type2,
				Height:           f.Height,
				Weight:           f.Weight,
				BaseHp:           f.BaseHp,
				BaseAtk:          f.BaseAtk,
				BaseDfe:          f.BaseDfe,
				BaseSpd:          f.BaseSpd,
				BaseAts:          f.BaseAts,
				BaseDfs:          f.BaseDfs,
				EvHp:             f.EvHp,
				EvAtk:            f.EvAtk,
				EvDfe:            f.EvDfe,
				EvSpd:            f.EvSpd,
				EvAts:            f.EvAts,
				EvDfs:            f.EvDfs,
				Evolutions:       mapEvolutions(f.Evolutions),
				ExperienceType:   ExperienceTypeMap[f.ExperienceType],
				BaseExperience:   f.BaseExperience,
				BaseLoyalty:      f.BaseLoyalty,
				CatchRate:        f.CatchRate,
				FemaleRate:       f.FemaleRate,
				BreedGroups:      mapBreedGroups(f.BreedGroups),
				HatchSteps:       f.HatchSteps,
				BabyDbSymbol:     f.BabyDbSymbol,
				BabyForm:         f.BabyForm,
				ItemHeld:         mapItemHelds(f.ItemHeld),
				Abilities:        abilities,
				FrontOffsetY:     f.FrontOffsetY,
				Name:             f.Name,
				Description:      f.Description,
				CustomProperties: make(map[string]any),
				Resources:        mapPokemonResources(f.Resources),
			}
		}

		store.AddPokemon(pkmn.AddPokemonDto{
			ID:       descriptor.ID,
			DbSymbol: descriptor.DbSymbol,
			Forms:    forms,
		})
	}

	return store, nil
}

func mapTypeDamages(damages []TypeDamageDescriptor) map[string]float32 {
	mapped := make(map[string]float32, len(damages))
	for _, tdDesc := range damages {
		mapped[tdDesc.DefensiveType] = tdDesc.Factor
	}
	return mapped
}

func mapBreedGroups(breedGroupInts []int32) []string {
	breedGroups := make([]string, len(breedGroupInts))
	for i, bgInt := range breedGroupInts {
		breedGroups[i] = BreedMap[BreedGroupDescriptor(bgInt)]
	}
	return breedGroups
}

func mapEvolutions(evolutions []EvolutionDescriptor) []pkmn.Evolution {
	if len(evolutions) == 0 {
		return nil
	}

	mapped := make([]pkmn.Evolution, len(evolutions))
	for i, evoDesc := range evolutions {
		mapped[i] = pkmn.Evolution{
			DbSymbol:   evoDesc.DbSymbol,
			Form:       evoDesc.Form,
			Conditions: mapConditions(evoDesc.Conditions),
		}
	}
	return mapped
}

func mapConditions(conditions []ConditionDescriptor) []pkmn.Condition {
	if len(conditions) == 0 {
		return nil
	}

	mapped := make([]pkmn.Condition, len(conditions))
	for i, condDesc := range conditions {
		mapped[i] = pkmn.Condition{Type: condDesc.Type}
	}
	return mapped
}

func mapItemHelds(itemHelds []ItemHeldDescriptor) []*pkmn.ItemHeld {
	if len(itemHelds) == 0 {
		return nil
	}

	mapped := make([]*pkmn.ItemHeld, len(itemHelds))
	for i, ihDesc := range itemHelds {
		mapped[i] = &pkmn.ItemHeld{
			DbSymbol: ihDesc.DbSymbol,
			Chance:   ihDesc.Chance,
		}
	}
	return mapped
}

func mapPokemonResources(resources PokemonResourcesDescriptor) pkmn.PokemonResources {
	return pkmn.PokemonResources{
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

func mapTargeting(desc MoveDescriptor) pkmn.MoveTargeting {
	targeting := pkmn.MoveTargeting{
		AimedTarget: pkmn.AimedTarget(desc.BattleEngineAimedTarget),
	}

	if desc.IsDirect {
		targeting.ContactType = pkmn.ContactTypeDirect
	} else if desc.IsDistance {
		targeting.ContactType = pkmn.ContactTypeDistant
	} else {
		targeting.ContactType = pkmn.ContactTypeNone
	}

	return targeting
}

func mapExecution(desc MoveDescriptor) pkmn.MoveExecution {
	return pkmn.MoveExecution{
		Method:   pkmn.ExecutionMethod(desc.BattleEngineMethod),
		Charge:   desc.IsCharge,
		Recharge: desc.IsRecharge,
	}
}

func mapMechanicalTags(desc MoveDescriptor) []pkmn.MoveMechanicalTag {
	tags := make([]pkmn.MoveMechanicalTag, 0)

	if desc.IsAuthentic {
		tags = append(tags, pkmn.MechanicalTagAuthentic)
	}
	if desc.IsBallistics {
		tags = append(tags, pkmn.MechanicalTagBallistic)
	}
	if desc.IsBite {
		tags = append(tags, pkmn.MechanicalTagBite)
	}
	if desc.IsDance {
		tags = append(tags, pkmn.MechanicalTagDance)
	}
	if desc.IsPunch {
		tags = append(tags, pkmn.MechanicalTagPunch)
	}
	if desc.IsSlicingAttack {
		tags = append(tags, pkmn.MechanicalTagSlice)
	}
	if desc.IsSoundAttack {
		tags = append(tags, pkmn.MechanicalTagSound)
	}
	if desc.IsWind {
		tags = append(tags, pkmn.MechanicalTagWind)
	}
	if desc.IsPulse {
		tags = append(tags, pkmn.MechanicalTagPulse)
	}
	if desc.IsPowder {
		tags = append(tags, pkmn.MechanicalTagPowder)
	}
	if desc.IsMental {
		tags = append(tags, pkmn.MechanicalTagMental)
	}

	return tags
}

func mapInteractions(desc MoveDescriptor) []pkmn.MoveInteraction {
	interactions := make([]pkmn.MoveInteraction, 0)

	if desc.IsBlocable {
		interactions = append(interactions, pkmn.InteractionBlocable)
	}
	if desc.IsMirrorMove {
		interactions = append(interactions, pkmn.InteractionMirrorMove)
	}
	if desc.IsSnatchable {
		interactions = append(interactions, pkmn.InteractionSnatchable)
	}
	if desc.IsMagicCoatAffected {
		interactions = append(interactions, pkmn.InteractionMagicCoatAffected)
	}
	if desc.IsKingRockUtility {
		interactions = append(interactions, pkmn.InteractionKingRockUtility)
	}
	if desc.IsGravity {
		interactions = append(interactions, pkmn.InteractionAffectedByGravity)
	}
	if desc.IsNonSkyBattle {
		interactions = append(interactions, pkmn.InteractionNonSkyBattle)
	}

	return interactions
}

func mapSecondaryEffects(desc MoveDescriptor) pkmn.MoveSecondaryEffects {
	effects := pkmn.MoveSecondaryEffects{
		Chance: desc.EffectChance,
	}

	if len(desc.MoveStatus) > 0 {
		effects.StatusEffects = make([]pkmn.MoveStatusEffect, len(desc.MoveStatus))
		for i, status := range desc.MoveStatus {
			effects.StatusEffects[i] = pkmn.MoveStatusEffect{
				Status:   status.Status,
				LuckRate: status.LuckRate,
			}
		}
	}

	if len(desc.BattleStageMod) > 0 {
		effects.StatStageChanges = make([]pkmn.MoveStatStageChange, len(desc.BattleStageMod))
		for i, stageMod := range desc.BattleStageMod {
			effects.StatStageChanges[i] = pkmn.MoveStatStageChange{
				BattleStage: pkmn.BattleStage(stageMod.BattleStage),
				Modificator: stageMod.Modificator,
			}
		}
	}

	return effects
}

func ImportAbility(studioFolder, translationFolder string) (iter.Seq[*AbilityDescriptor], error) {
	abilityFolderPath := path.Join(studioFolder, "abilities/")
	slog.Info("Importing ability folder", "path", abilityFolderPath)
	abilityFileIterator, err := file.ImportFolder(abilityFolderPath)
	if err != nil {
		return nil, err
	}

	// Import translations
	abilityNametranslationsPath := path.Join(translationFolder, abilityNameTranslationFile)
	abilityTranslations, err := ImportTranslations(abilityNametranslationsPath)
	if err != nil {
		slog.Warn("Failed to import ability translations", "path", abilityNametranslationsPath, "error", err)
		abilityTranslations = []pkmn.Translation{}
	}

	abilityDescriptiontranslationsPath := path.Join(translationFolder, abilityDescriptionTranslationFile)
	abilityDescTranslations, err := ImportTranslations(abilityDescriptiontranslationsPath)
	if err != nil {
		slog.Warn("Failed to import ability description translations", "path", abilityDescriptiontranslationsPath, "error", err)
		abilityDescTranslations = []pkmn.Translation{}
	}

	return func(yield func(*AbilityDescriptor) bool) {
		for abilityFile := range abilityFileIterator {
			abilityDesc, err := UnmarshalAbilityDescriptor(abilityFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal ability descriptor content", "file", abilityFile.Path, "error", err)
				continue
			}

			abilityDesc.Name = MapTranslation(abilityDesc.TextID, abilityTranslations)
			abilityDesc.Description = MapTranslation(abilityDesc.TextID, abilityDescTranslations)

			if abilityDesc != nil {
				if !yield(abilityDesc) {
					break
				}
			}
		}
	}, nil
}

// ImportPokemon import a pokemon folder to a store
// studioFolder pokemon studio folder
// translationFolder the translation folder
// store the store that import is sending data to
func ImportPokemon(studioFolder, translationFolder string) (iter.Seq[*PokemonDescriptor], error) {
	pokemonFolderPath := path.Join(studioFolder, "pokemon/")
	slog.Info("Importing pokemon folder", "path", pokemonFolderPath)
	pokemonFileIterator, err := file.ImportFolder(pokemonFolderPath)
	if err != nil {
		return nil, err
	}

	pokemonNameTranslationPath := path.Join(translationFolder, pokemonNameTranslationFile)
	pokemonNameTranslations := ImportTranslationsOrEmpty(pokemonNameTranslationPath)

	pokemonDescriptionTranslationPath := path.Join(translationFolder, pokemonDescriptionTranslationFile)
	pokemonDescriptionTranslations := ImportTranslationsOrEmpty(pokemonDescriptionTranslationPath)

	pokemonFormNameTranslationPath := path.Join(translationFolder, pokemonFormNameTranslationFile)
	pokemonFormNameTranslations := ImportTranslationsOrEmpty(pokemonFormNameTranslationPath)

	pokemonFormDescriptionTranslationPath := path.Join(translationFolder, pokemonFormDescriptionTranslationFile)
	pokemonFormDescriptionTranslations := ImportTranslationsOrEmpty(pokemonFormDescriptionTranslationPath)

	return func(yield func(*PokemonDescriptor) bool) {
		for pokemonFile := range pokemonFileIterator {
			pokemonDesc, err := UnmarshalPokemonDescriptor(pokemonFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal pokemon descriptor content", "file", pokemonFile.Path, "error", err)
				continue
			}

			for i := range pokemonDesc.Forms {
				form := &pokemonDesc.Forms[i]
				formTextId := form.FormTextId

				if formTextId.Name == 0{
					form.Name = MapTranslation(int(pokemonDesc.ID), pokemonNameTranslations)
				}else{
					form.Name = MapTranslation(formTextId.Name, pokemonFormNameTranslations)
				}

				if formTextId.Description == 0 {
					form.Description = MapTranslation(int(pokemonDesc.ID), pokemonDescriptionTranslations)
				} else {
					form.Description = MapTranslation(formTextId.Description, pokemonFormDescriptionTranslations)
				}
			}

			if pokemonDesc != nil {
				if !yield(pokemonDesc) {
					break
				}
			}
		}
	}, nil
}

// ImportTypes import a type folder to a store
// studioFolder pokemon studio folder
// translationFolder the translation folder
// store the store that import is sending data to
func ImportTypes(studioFolder, translationFolder string) (iter.Seq[*PokemonTypeDescriptor], error) {
	typeFolderPath := path.Join(studioFolder, "types/")
	slog.Info("Importing type folder", "path", typeFolderPath)
	typeFileIterator, err := file.ImportFolder(typeFolderPath)
	if err != nil {
		return nil, err
	}

	// Import translations
	translationsPath := path.Join(translationFolder, typeNameTranslationFile)
	translations, err := ImportTranslations(translationsPath)
	if err != nil {
		slog.Warn("Failed to import type translations", "path", translationsPath, "error", err)
		translations = []pkmn.Translation{}
	}

	return func(yield func(*PokemonTypeDescriptor) bool) {
		for typeFile := range typeFileIterator {
			typeDesc, err := UnmarshalTypeDescriptor(typeFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal type descriptor content", "file", typeFile.Path, "error", err)
				continue
			}

			typeDesc.Name = MapTranslation(typeDesc.TextId, translations)

			if typeDesc != nil {
				if !yield(typeDesc) {
					break
				}
			}
		}
	}, nil
}

// ImportMoves import a moves folder to a store
// studioFolder pokemon studio folder
// translationFolder the translation folder
func ImportMoves(studioFolder, translationFolder string) (iter.Seq[*MoveDescriptor], error) {
	moveFolderPath := path.Join(studioFolder, "moves/")
	slog.Info("Importing move folder", "path", moveFolderPath)
	moveFileIterator, err := file.ImportFolder(moveFolderPath)
	if err != nil {
		return nil, err
	}

	// Import translations
	moveNameTranslationsPath := path.Join(translationFolder, moveNameTranslationFile)
	moveNameTranslations, err := ImportTranslations(moveNameTranslationsPath)
	if err != nil {
		slog.Warn("Failed to import move name translations", "path", moveNameTranslationsPath, "error", err)
		moveNameTranslations = []pkmn.Translation{}
	}

	moveDescTranslationsPath := path.Join(translationFolder, moveDescriptionTranslationFile)
	moveDescTranslations, err := ImportTranslations(moveDescTranslationsPath)
	if err != nil {
		slog.Warn("Failed to import move description translations", "path", moveDescTranslationsPath, "error", err)
		moveDescTranslations = []pkmn.Translation{}
	}

	return func(yield func(*MoveDescriptor) bool) {
		for moveFile := range moveFileIterator {
			moveDesc, err := UnmarshalMoveDescriptor(moveFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal move descriptor content", "file", moveFile.Path, "error", err)
				continue
			}

			moveDesc.Name = MapTranslation(moveDesc.Id, moveNameTranslations)
			moveDesc.Description = MapTranslation(moveDesc.Id, moveDescTranslations)

			if moveDesc != nil {
				if !yield(moveDesc) {
					break
				}
			}
		}
	}, nil
}

// UnmarshalAbilityDescriptor unmarshal a json encoded ability to a descriptor
// abilityContent the encoded ability
func UnmarshalAbilityDescriptor(abilityContent []byte) (*AbilityDescriptor, error) {
	abilityDescriptor := &AbilityDescriptor{}
	if err := json.Unmarshal(abilityContent, abilityDescriptor); err != nil {
		return nil, err
	}

	return abilityDescriptor, nil
}

// UnmarshalPokemonDescriptor unmarshal a json encoded pokemon to a descriptor
// pokemonContent the encoded pokemon
func UnmarshalPokemonDescriptor(pokemonContent []byte) (*PokemonDescriptor, error) {
	pokemonDescriptor := &PokemonDescriptor{}
	if err := json.Unmarshal(pokemonContent, pokemonDescriptor); err != nil {
		return nil, err
	}

	// Handle __undef__ values by converting them to nil
	for i := range pokemonDescriptor.Forms {
		if pokemonDescriptor.Forms[i].Type2 != nil && *pokemonDescriptor.Forms[i].Type2 == pkmn.UndefType {
			pokemonDescriptor.Forms[i].Type2 = nil
		}
	}

	return pokemonDescriptor, nil
}

// UnmarshalTypeDescriptor unmarshal a json encoded type to a descriptor
// typeContent the encoded type
func UnmarshalTypeDescriptor(typeContent []byte) (*PokemonTypeDescriptor, error) {
	typeDescriptor := &PokemonTypeDescriptor{}
	if err := json.Unmarshal(typeContent, typeDescriptor); err != nil {
		return nil, err
	}

	return typeDescriptor, nil
}

// UnmarshalMoveDescriptor unmarshal a json encoded move to a descriptor
// moveContent the encoded move
func UnmarshalMoveDescriptor(moveContent []byte) (*MoveDescriptor, error) {
	moveDescriptor := &MoveDescriptor{}
	if err := json.Unmarshal(moveContent, moveDescriptor); err != nil {
		return nil, err
	}

	return moveDescriptor, nil
}

// ImportTranslationsOrEmpty import translations from file, if it fails log the error and return an empty list
func ImportTranslationsOrEmpty(path string) []pkmn.Translation {
	translations, err := ImportTranslations(path)
	if err != nil {
		slog.Error("Failed to import translations", "path", path, "error", err)
		return []pkmn.Translation{}
	}
	return translations
}

// ImportTranslations import translations from file
// path the path of the file to import
func ImportTranslations(path string) ([]pkmn.Translation, error) {
	curPath, _ := os.Getwd()
	slog.Info("Import translation file", "path", path, "currentPath", curPath)
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Warn("Failed to close translation file", "path", path, "error", err)
		}
	}(file)

	reader := csv.NewReader(file)

	langs, err := reader.Read()
	if err != nil {
		return nil, err
	}

	results := make([]pkmn.Translation, 0)
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		translationMap := make(pkmn.Translation)
		for index := range len(records) {
			translationMap[langs[index]] = records[index]
		}

		results = append(results, translationMap)
	}

	return results, nil
}

func MapTranslation(textId int, translations []pkmn.Translation) pkmn.Translation {
	if textId >= 0 && textId < len(translations) {
		translation := translations[textId]
		return translation
	}
	slog.Warn("Could not find translation", "TextID", textId, "available", len(translations))
	return nil
}
