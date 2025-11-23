package studio

import (
	"testing"
)

const (
	InvalidPath        = "invalid/path"
	TranslationInvalid = "../../test/test_resources/100003-invalid.csv"
	TranslationValid   = "../../test/test_resources/valid-data/Text/Dialogs/100003.csv"
)

func TestImportTranslations_FileNotFound(t *testing.T) {
	_, err := ImportTranslations(InvalidPath)
	if err == nil {
		t.Error("Import translation at invalid path should return error")
	}
}

func TestImportTranslations_InvalidFormat(t *testing.T) {
	_, err := ImportTranslations(TranslationInvalid)
	if err == nil {
		t.Error("Import translation with invalid format should return error")
	}

}

func TestImportTranslations_Ok(t *testing.T) {
	_, err := ImportTranslations(TranslationValid)
	if err != nil {
		t.Error("Import translation with valid file should not return error")
	}
}

func TestParseAcceptLanguageParam(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string returns default en",
			input:    "",
			expected: "en",
		},
		{
			name:     "Simple language code",
			input:    "fr",
			expected: "fr",
		},
		{
			name:     "Language with region",
			input:    "en-US",
			expected: "en",
		},
		{
			name:     "Language with quality factor",
			input:    "fr;q=0.9",
			expected: "fr",
		},
		{
			name:     "Multiple languages with quality factors",
			input:    "en-US,en;q=0.5",
			expected: "en",
		},
		{
			name:     "Language with region and quality factor",
			input:    "fr-CA;q=0.8",
			expected: "fr",
		},
		{
			name:     "Complex accept language header",
			input:    "fr-FR, fr;q=0.9, en;q=0.8, *;q=0.5",
			expected: "fr",
		},
		{
			name:     "Language with whitespace",
			input:    " es ",
			expected: "es",
		},
		{
			name:     "Language with region and whitespace",
			input:    " pt-BR ",
			expected: "pt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseAcceptLanguageParam(tt.input)
			if result != tt.expected {
				t.Errorf("ParseAcceptLanguageParam(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
