package pkmn

import (
	"testing"
)

func TestPokemonTypeResisted(t *testing.T) {
	tests := []struct {
		name     string
		attacker *PokemonType
		defType1 string
		defType2 string
		want     float32
	}{
		{
			name: "no damage to first type (immune)",
			attacker: newTestType(t, map[string]float32{
				"ghost":  0,
				"normal": 1,
			}),
			defType1: "ghost",
			defType2: "normal",
			want:     0,
		},
		{
			name: "no damage to second type (immune)",
			attacker: newTestType(t, map[string]float32{
				"ghost":  2,
				"normal": 0,
			}),
			defType1: "normal",
			defType2: "ghost",
			want:     0,
		},
		{
			name: "super effective against both types (4x damage)",
			attacker: newTestType(t, map[string]float32{
				"fire":  2,
				"grass": 2,
			}),
			defType1: "fire",
			defType2: "grass",
			want:     4,
		},
		{
			name: "not very effective against both types (0.25x damage)",
			attacker: newTestType(t, map[string]float32{
				"fire":  0.5,
				"grass": 0.5,
			}),
			defType1: "fire",
			defType2: "grass",
			want:    0.25,
		},
		{
			name: "neutral against both types",
			attacker: newTestType(t, map[string]float32{
				"water":  1,
				"ground": 1,
			}),
			defType1: "water",
			defType2: "ground",
			want:     1,
		},
		{
			name: "super effective against first type only",
			attacker: newTestType(t, map[string]float32{
				"fire":  2,
				"water": 1,
			}),
			defType1: "fire",
			defType2: "water",
			want:     2,
		},
		{
			name: "not very effective against first type only",
			attacker: newTestType(t, map[string]float32{
				"fire":  0.5,
				"water": 1,
			}),
			defType1: "fire",
			defType2: "water",
			want:     0.5,
		},
		{
			name: "unknown first type",
			attacker: newTestType(t, map[string]float32{
				"fire": 2,
			}),
			defType1: "unknown",
			defType2: "water",
			want:     1,
		},
		{
			name: "unknown second type",
			attacker: newTestType(t, map[string]float32{
				"fire":  2,
				"water": 1,
			}),
			defType1: "water",
			defType2: "unknown",
			want:     1,
		},
		{
			name: "empty defending types",
			attacker: newTestType(t, map[string]float32{
				"fire": 2,
			}),
			defType1: "",
			defType2: "",
			want:     1,
		},
		{
			name: "first type immune stops second type calculation",
			attacker: newTestType(t, map[string]float32{
				"ghost":  0,
				"normal": 2,
			}),
			defType1: "ghost",
			defType2: "normal",
			want:     0,
		},
		{
			name: "half damage to both types (0.5 each = 4x resist)",
			attacker: newTestType(t, map[string]float32{
				"fire":  0.5,
				"grass": 0.5,
			}),
			defType1: "fire",
			defType2: "grass",
			want:     0.25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.attacker.DamageToTypes(tt.defType1, tt.defType2)
			if got != tt.want {
				t.Errorf("Resisted(%q, %q) = %v, want %v", tt.defType1, tt.defType2, got, tt.want)
			}
		})
	}
}

