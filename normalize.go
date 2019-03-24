package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// Normalize string options
const (
	NormalizeStringOptionLatinASCII       uint64 = C.LIBPOSTAL_NORMALIZE_STRING_LATIN_ASCII
	NormalizeStringOptionTransliterate    uint64 = C.LIBPOSTAL_NORMALIZE_STRING_TRANSLITERATE
	NormalizeStringOptionStripAccents     uint64 = C.LIBPOSTAL_NORMALIZE_STRING_STRIP_ACCENTS
	NormalizeStringOptionDecompose        uint64 = C.LIBPOSTAL_NORMALIZE_STRING_DECOMPOSE
	NormalizeStringOptionLowercase        uint64 = C.LIBPOSTAL_NORMALIZE_STRING_LOWERCASE
	NormalizeStringOptionTrim             uint64 = C.LIBPOSTAL_NORMALIZE_STRING_TRIM
	NormalizeStringOptionReplaceHyphens   uint64 = C.LIBPOSTAL_NORMALIZE_STRING_REPLACE_HYPHENS
	NormalizeStringOptionCompose          uint64 = C.LIBPOSTAL_NORMALIZE_STRING_COMPOSE
	NormalizeStringOptionSimpleLatinASCII uint64 = C.LIBPOSTAL_NORMALIZE_STRING_SIMPLE_LATIN_ASCII
	NormalizeStringOptionReplaceNumex     uint64 = C.LIBPOSTAL_NORMALIZE_STRING_REPLACE_NUMEX
	NormalizeStringOptionDefault          uint64 = (C.LIBPOSTAL_NORMALIZE_STRING_LATIN_ASCII | C.LIBPOSTAL_NORMALIZE_STRING_COMPOSE | C.LIBPOSTAL_NORMALIZE_STRING_TRIM | C.LIBPOSTAL_NORMALIZE_STRING_REPLACE_HYPHENS | C.LIBPOSTAL_NORMALIZE_STRING_STRIP_ACCENTS | C.LIBPOSTAL_NORMALIZE_STRING_LOWERCASE)
)

// Normalize token options
const (
	NormalizeTokenOptionReplaceHyphens             uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_REPLACE_HYPHENS
	NormalizeTokenOptionDeleteHyphens              uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_HYPHENS
	NormalizeTokenOptionDeleteFinalPeriod          uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_FINAL_PERIOD
	NormalizeTokenOptionDeleteAcronymPeriods       uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_ACRONYM_PERIODS
	NormalizeTokenOptionDropEnglishPossessives     uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_DROP_ENGLISH_POSSESSIVES
	NormalizeTokenOptionDeleteOtherApostrophe      uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_OTHER_APOSTROPHE
	NormalizeTokenOptionSplitAlphaFromNumeric      uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_SPLIT_ALPHA_FROM_NUMERIC
	NormalizeTokenOptionReplaceDigits              uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_REPLACE_DIGITS
	NormalizeTokenOptionReplaceNumericTokenLetters uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_REPLACE_NUMERIC_TOKEN_LETTERS
	NormalizeTokenOptionReplaceNumericHyphens      uint64 = C.LIBPOSTAL_NORMALIZE_TOKEN_REPLACE_NUMERIC_HYPHENS
	NormalizeTokenOptionDefault                    uint64 = (C.LIBPOSTAL_NORMALIZE_TOKEN_REPLACE_HYPHENS | C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_FINAL_PERIOD | C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_ACRONYM_PERIODS | C.LIBPOSTAL_NORMALIZE_TOKEN_DROP_ENGLISH_POSSESSIVES | C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_OTHER_APOSTROPHE)
	NormalizeTokenOptionDropPeriods                uint64 = (C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_FINAL_PERIOD | C.LIBPOSTAL_NORMALIZE_TOKEN_DELETE_ACRONYM_PERIODS)
	NormalizeTokenOptionsDefaultNumeric            uint64 = (C.LIBPOSTAL_NORMALIZE_DEFAULT_TOKEN_OPTIONS | C.LIBPOSTAL_NORMALIZE_TOKEN_SPLIT_ALPHA_FROM_NUMERIC)
)

type NormalizedToken struct {
	String string
	Token  Token
}

type NormalizeOptions struct {
	Languages              []string
	AddressComponents      uint16
	LatinASCII             bool
	Transliterate          bool
	StripAccents           bool
	Decompose              bool
	Lowercase              bool
	TrimString             bool
	ReplaceWordHyphens     bool
	DeleteWordHyphens      bool
	ReplaceNumericHyphens  bool
	DeleteNumericHyphens   bool
	SplitAlphaFromNumeric  bool
	DeleteFinalPeriods     bool
	DeleteAcronymPeriods   bool
	DropEnglishPossessives bool
	DeleteApostrophes      bool
	ExpandNumex            bool
	RomanNumerals          bool
}

func DefaultNormalizeOptions() NormalizeOptions {
	cDefaultNormalizeOptions := C.libpostal_get_default_options()
	return NormalizeOptions{
		Languages:              nil,
		AddressComponents:      uint16(cDefaultNormalizeOptions.address_components),
		LatinASCII:             bool(cDefaultNormalizeOptions.latin_ascii),
		Transliterate:          bool(cDefaultNormalizeOptions.transliterate),
		StripAccents:           bool(cDefaultNormalizeOptions.strip_accents),
		Decompose:              bool(cDefaultNormalizeOptions.decompose),
		Lowercase:              bool(cDefaultNormalizeOptions.lowercase),
		TrimString:             bool(cDefaultNormalizeOptions.trim_string),
		ReplaceWordHyphens:     bool(cDefaultNormalizeOptions.replace_word_hyphens),
		DeleteWordHyphens:      bool(cDefaultNormalizeOptions.delete_word_hyphens),
		ReplaceNumericHyphens:  bool(cDefaultNormalizeOptions.replace_numeric_hyphens),
		DeleteNumericHyphens:   bool(cDefaultNormalizeOptions.delete_numeric_hyphens),
		SplitAlphaFromNumeric:  bool(cDefaultNormalizeOptions.split_alpha_from_numeric),
		DeleteFinalPeriods:     bool(cDefaultNormalizeOptions.delete_final_periods),
		DeleteAcronymPeriods:   bool(cDefaultNormalizeOptions.delete_acronym_periods),
		DropEnglishPossessives: bool(cDefaultNormalizeOptions.drop_english_possessives),
		DeleteApostrophes:      bool(cDefaultNormalizeOptions.delete_apostrophes),
		ExpandNumex:            bool(cDefaultNormalizeOptions.expand_numex),
		RomanNumerals:          bool(cDefaultNormalizeOptions.roman_numerals),
	}
}

func NormalizedTokens(input string, stringOptions uint64, tokenOptions uint64, whitespace bool, languages []string) NormalizedToken {
	cInput := C.CString(input)
	cStringOptions := C.uint64_t(stringOptions)
	cTokenOptions := C.uint64_t(tokenOptions)
	cWhitespace := C.bool(whitespace)
	cn := C.size_t(len(input))

	cNumLanguages := C.size_t(0)
	var cLanguages []*C.char

	if languages != nil {
		cLanguages = make([]*C.char, len(languages))
		cNumLanguages = C.size_t(len(languages))

		for i := 0; i < len(languages); i++ {
			cLang := C.CString(languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}
	}

	var cToken *C.libpostal_normalized_token_t
	if cNumLanguages > 0 {
		cToken = C.libpostal_normalized_tokens_languages(cInput, cStringOptions, cTokenOptions, cWhitespace, cNumLanguages, &cLanguages[0], &cn)
	} else {
		cToken = C.libpostal_normalized_tokens(cInput, cStringOptions, cTokenOptions, cWhitespace, &cn)
	}

	return NormalizedToken{
		String: C.GoString(cToken.str),
		Token: Token{
			Offset: int(cToken.token.offset),
			Len:    int(cToken.token.len),
			Type:   uint16(cToken.token._type),
		},
	}
}

func NormalizeString(input string, stringOptions uint64, languages []string) string {
	cInput := C.CString(input)
	cStringOptions := C.uint64_t(stringOptions)

	cNumLanguages := C.size_t(0)
	var cLanguages []*C.char

	if languages != nil {
		cLanguages = make([]*C.char, len(languages))
		cNumLanguages = C.size_t(len(languages))

		for i := 0; i < len(languages); i++ {
			cLang := C.CString(languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}
	}

	if cNumLanguages > 0 {
		return C.GoString(C.libpostal_normalize_string_languages(cInput, cStringOptions, cNumLanguages, &cLanguages[0]))
	}

	return C.GoString(C.libpostal_normalize_string(cInput, cStringOptions))
}
