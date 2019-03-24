package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

const (
	TokenTypeEnd               uint16 = C.LIBPOSTAL_TOKEN_TYPE_END
	TokenTypeWord              uint16 = C.LIBPOSTAL_TOKEN_TYPE_WORD
	TokenTypeAbbreviation      uint16 = C.LIBPOSTAL_TOKEN_TYPE_ABBREVIATION
	TokenTypeIdeographicChar   uint16 = C.LIBPOSTAL_TOKEN_TYPE_IDEOGRAPHIC_CHAR
	TokenTypeHangulSyllable    uint16 = C.LIBPOSTAL_TOKEN_TYPE_HANGUL_SYLLABLE
	TokenTypeAcronym           uint16 = C.LIBPOSTAL_TOKEN_TYPE_ACRONYM
	TokenTypePhrase            uint16 = C.LIBPOSTAL_TOKEN_TYPE_PHRASE
	TokenTypeEmail             uint16 = C.LIBPOSTAL_TOKEN_TYPE_EMAIL
	TokenTypeURL               uint16 = C.LIBPOSTAL_TOKEN_TYPE_URL
	TokenTypeUsPhone           uint16 = C.LIBPOSTAL_TOKEN_TYPE_US_PHONE
	TokenTypeIntlPhone         uint16 = C.LIBPOSTAL_TOKEN_TYPE_INTL_PHONE
	TokenTypeNumeric           uint16 = C.LIBPOSTAL_TOKEN_TYPE_NUMERIC
	TokenTypeOrdinal           uint16 = C.LIBPOSTAL_TOKEN_TYPE_ORDINAL
	TokenTypeRomanNumeral      uint16 = C.LIBPOSTAL_TOKEN_TYPE_ROMAN_NUMERAL
	TokenTypeIdeographicNumber uint16 = C.LIBPOSTAL_TOKEN_TYPE_IDEOGRAPHIC_NUMBER
	TokenTypePeriod            uint16 = C.LIBPOSTAL_TOKEN_TYPE_PERIOD
	TokenTypeExclamation       uint16 = C.LIBPOSTAL_TOKEN_TYPE_EXCLAMATION
	TokenTypeQuestionMark      uint16 = C.LIBPOSTAL_TOKEN_TYPE_QUESTION_MARK
	TokenTypeComma             uint16 = C.LIBPOSTAL_TOKEN_TYPE_COMMA
	TokenTypeColon             uint16 = C.LIBPOSTAL_TOKEN_TYPE_COLON
	TokenTypeSemicolon         uint16 = C.LIBPOSTAL_TOKEN_TYPE_SEMICOLON
	TokenTypePlus              uint16 = C.LIBPOSTAL_TOKEN_TYPE_PLUS
	TokenTypeAmpersand         uint16 = C.LIBPOSTAL_TOKEN_TYPE_AMPERSAND
	TokenTypeAtSign            uint16 = C.LIBPOSTAL_TOKEN_TYPE_AT_SIGN
	TokenTypePound             uint16 = C.LIBPOSTAL_TOKEN_TYPE_POUND
	TokenTypeEllipsis          uint16 = C.LIBPOSTAL_TOKEN_TYPE_ELLIPSIS
	TokenTypeDash              uint16 = C.LIBPOSTAL_TOKEN_TYPE_DASH
	TokenTypeBreakingDash      uint16 = C.LIBPOSTAL_TOKEN_TYPE_BREAKING_DASH
	TokenTypeHyphen            uint16 = C.LIBPOSTAL_TOKEN_TYPE_HYPHEN
	TokenTypePunctOpen         uint16 = C.LIBPOSTAL_TOKEN_TYPE_PUNCT_OPEN
	TokenTypePunctClose        uint16 = C.LIBPOSTAL_TOKEN_TYPE_PUNCT_CLOSE
	TokenTypeDoubleQuote       uint16 = C.LIBPOSTAL_TOKEN_TYPE_DOUBLE_QUOTE
	TokenTypeSingleQuote       uint16 = C.LIBPOSTAL_TOKEN_TYPE_SINGLE_QUOTE
	TokenTypeOpenQuote         uint16 = C.LIBPOSTAL_TOKEN_TYPE_OPEN_QUOTE
	TokenTypeCloseQuote        uint16 = C.LIBPOSTAL_TOKEN_TYPE_CLOSE_QUOTE
	TokenTypeSlash             uint16 = C.LIBPOSTAL_TOKEN_TYPE_SLASH
	TokenTypeBackslash         uint16 = C.LIBPOSTAL_TOKEN_TYPE_BACKSLASH
	TokenTypeGreaterThan       uint16 = C.LIBPOSTAL_TOKEN_TYPE_GREATER_THAN
	TokenTypeLessThan          uint16 = C.LIBPOSTAL_TOKEN_TYPE_LESS_THAN
	TokenTypeOther             uint16 = C.LIBPOSTAL_TOKEN_TYPE_OTHER
	TokenTypeWhitespace        uint16 = C.LIBPOSTAL_TOKEN_TYPE_WHITESPACE
	TokenTypeNewline           uint16 = C.LIBPOSTAL_TOKEN_TYPE_NEWLINE
	TokenTypeInvalidChar       uint16 = C.LIBPOSTAL_TOKEN_TYPE_INVALID_CHAR
)

type Token struct {
	Offset int
	Len    int
	Type   uint16
}

func Tokenize(input string, whitespace bool) []Token {
	cInput := C.CString(input)
	defer C.free(unsafe.Pointer(cInput))

	cWhitespace := C.bool(whitespace)
	var cNumTokens = C.size_t(0)

	cTokens := C.libpostal_tokenize(cInput, cWhitespace, &cNumTokens)
	numTokens := uint64(cNumTokens)

	cTokensPtr := (*[1 << 28]C.libpostal_token_t)(unsafe.Pointer(cTokens))[:cNumTokens:cNumTokens]

	var tokens []Token
	var i uint64
	for i = 0; i < numTokens; i++ {
		token := Token{
			Offset: int(cTokensPtr[i].offset),
			Len:    int(cTokensPtr[i].len),
			Type:   uint16(cTokensPtr[i]._type),
		}
		tokens = append(tokens, token)
	}

	return tokens
}
