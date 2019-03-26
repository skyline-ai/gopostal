package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type DuplicateStatus int

func (d DuplicateStatus) String() string {
	switch d {
	case DuplicateStatusNull:
		return "Null Duplicate"
	case DuplicateStatusNon:
		return "Non Duplicate"
	case DuplicateStatusPossible:
		return "Possible Duplicate Needs Review"
	case DuplicateStatusLikely:
		return "Likely Duplicate"
	case DuplicateStatusExact:
		return "Exact Duplicate"
	}
	return "Unknown Duplicate Status"
}

type DuplicateOptions struct {
	Languages []string
}

type FuzzyDuplicateOptions struct {
	Languages            []string
	NeedsReviewThreshold float64
	LikelyDupeThreshold  float64
}

func DefaultFuzzyDuplicateOptions() FuzzyDuplicateOptions {
	cDefaultFuzzyDuplicateOptions := C.libpostal_get_default_fuzzy_duplicate_options()
	return FuzzyDuplicateOptions{
		NeedsReviewThreshold: float64(cDefaultFuzzyDuplicateOptions.needs_review_threshold),
		LikelyDupeThreshold:  float64(cDefaultFuzzyDuplicateOptions.likely_dupe_threshold),
	}
}

const (
	DuplicateStatusNull     DuplicateStatus = C.LIBPOSTAL_NULL_DUPLICATE_STATUS
	DuplicateStatusNon      DuplicateStatus = C.LIBPOSTAL_NON_DUPLICATE
	DuplicateStatusPossible DuplicateStatus = C.LIBPOSTAL_POSSIBLE_DUPLICATE_NEEDS_REVIEW
	DuplicateStatusLikely   DuplicateStatus = C.LIBPOSTAL_LIKELY_DUPLICATE
	DuplicateStatusExact    DuplicateStatus = C.LIBPOSTAL_EXACT_DUPLICATE
)

func DefaultDuplicateOptions() DuplicateOptions {
	return DuplicateOptions{}
}

func IsDuplicate(addressComponent uint16, value1, value2 string, options DuplicateOptions) (DuplicateStatus, error) {
	cValue1 := C.CString(value1)
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := C.CString(value2)
	defer C.free(unsafe.Pointer(cValue2))

	cOptions := C.libpostal_get_default_duplicate_options()

	if options.Languages != nil {
		cLanguages := make([]*C.char, len(options.Languages))

		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}

		cOptions.languages = &cLanguages[0]
		cOptions.num_languages = C.size_t(len(options.Languages))
	} else {
		cOptions.num_languages = 0
	}

	var status C.libpostal_duplicate_status_t
	switch addressComponent {
	case AddressStreet:
		status = C.libpostal_is_street_duplicate(cValue1, cValue2, cOptions)
	case AddressName:
		status = C.libpostal_is_name_duplicate(cValue1, cValue2, cOptions)
	case AddressHouseNumber:
		status = C.libpostal_is_house_number_duplicate(cValue1, cValue2, cOptions)
	case AddressPoBox:
		status = C.libpostal_is_po_box_duplicate(cValue1, cValue2, cOptions)
	case AddressUnit:
		status = C.libpostal_is_unit_duplicate(cValue1, cValue2, cOptions)
	case AddressPostalCode:
		status = C.libpostal_is_postal_code_duplicate(cValue1, cValue2, cOptions)
	default:
		return 0, fmt.Errorf("unsupported address component: %s", AddressComponentString(addressComponent))
	}

	return DuplicateStatus(status), nil
}

func IsDuplicateFuzzy(addressComponent uint16, tokens1 []string, scores1 []float64, tokens2 []string, scores2 []float64, options FuzzyDuplicateOptions) (DuplicateStatus, float64, error) {
	if len(tokens1) != len(scores1) {
		return 0, 0, fmt.Errorf("tokens1 and scores1 arrays must be of equal length")
	}
	if len(tokens2) != len(scores2) {
		return 0, 0, fmt.Errorf("tokens1 and scores1 arrays must be of equal length")
	}

	// avoid segfault
	if len(tokens1) > len(tokens2) {
		for i := len(tokens2); i < len(tokens1); i++ {
			tokens2 = append(tokens2, "")
			scores2 = append(scores2, 0.0)
		}
	}
	if len(tokens2) > len(tokens1) {
		for i := len(tokens1); i < len(tokens2); i++ {
			tokens1 = append(tokens1, "")
			scores1 = append(scores1, 0.0)
		}
	}

	cOptions := C.libpostal_get_default_fuzzy_duplicate_options()
	cOptions.needs_review_threshold = C.double(options.NeedsReviewThreshold)
	cOptions.likely_dupe_threshold = C.double(options.LikelyDupeThreshold)

	if options.Languages != nil {
		cLanguages := make([]*C.char, len(options.Languages))

		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}

		cOptions.languages = &cLanguages[0]
		cOptions.num_languages = C.size_t(len(options.Languages))
	}

	var cTokens1 []*C.char
	for _, token := range tokens1 {
		cToken := C.CString(token)
		defer C.free(unsafe.Pointer(cToken))
		cTokens1 = append(cTokens1, cToken)
	}

	var cTokens2 []*C.char
	for _, token := range tokens2 {
		cToken := C.CString(token)
		defer C.free(unsafe.Pointer(cToken))
		cTokens2 = append(cTokens2, cToken)
	}

	var cScores1 []C.double
	for _, score := range scores1 {
		cScores1 = append(cScores1, C.double(score))
	}

	var cScores2 []C.double
	for _, score := range scores2 {
		cScores2 = append(cScores2, C.double(score))
	}

	cNumTokens1 := C.ulong(len(tokens1))
	cNumTokens2 := C.ulong(len(tokens2))

	var cStatus C.libpostal_fuzzy_duplicate_status_t
	switch addressComponent {
	case AddressStreet:
		cStatus = C.libpostal_is_street_duplicate_fuzzy(cNumTokens1, &cTokens1[0], &cScores1[0], cNumTokens2, &cTokens2[0], &cScores2[0], cOptions)
	case AddressName:
		cStatus = C.libpostal_is_name_duplicate_fuzzy(cNumTokens1, &cTokens1[0], &cScores1[0], cNumTokens2, &cTokens2[0], &cScores2[0], cOptions)
	default:
		return 0, 0, fmt.Errorf("unsupported address component: %s", AddressComponentString(addressComponent))
	}

	return DuplicateStatus(cStatus.status), float64(cStatus.similarity), nil
}

func IsToponymDuplicate(comps1 map[string]string, comps2 map[string]string, options DuplicateOptions) DuplicateStatus {
	cOptions := C.libpostal_get_default_duplicate_options()
	if options.Languages != nil {
		cLanguages := make([]*C.char, len(options.Languages))
		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}
		cOptions.languages = &cLanguages[0]
		cOptions.num_languages = C.size_t(len(options.Languages))
	} else {
		cOptions.num_languages = 0
	}

	var cLabels1, cValues1, cLabels2, cValues2 []*C.char

	for label, value := range comps1 {
		cLabel := C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
		cLabels1 = append(cLabels1, cLabel)

		cValue := C.CString(value)
		defer C.free(unsafe.Pointer(cValue))
		cValues1 = append(cValues1, cValue)
	}

	for label, value := range comps2 {
		cLabel := C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
		cLabels2 = append(cLabels2, cLabel)

		cValue := C.CString(value)
		defer C.free(unsafe.Pointer(cValue))
		cValues2 = append(cValues2, cValue)
	}

	cNumComponents1 := C.ulong(len(comps1))
	cNumComponents2 := C.ulong(len(comps2))

	cStatus := C.libpostal_is_toponym_duplicate(cNumComponents1, &cLabels1[0], &cValues1[0], cNumComponents2, &cLabels2[0], &cValues2[0], cOptions)

	return DuplicateStatus(cStatus)
}
