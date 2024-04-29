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