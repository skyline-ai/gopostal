package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"

import (
	"log"
	"unicode/utf8"
	"unsafe"
)

var (
	parserLoaded = false
)

type ParserOptions struct {
	Language string
	Country  string
}

func DefaultParserOptions() ParserOptions {
	return ParserOptions{}
}

func setupParser() {
	if !parserLoaded {
		parserLoaded = bool(C.libpostal_setup_parser())
	}
	if !parserLoaded {
		log.Fatalf("Could not load libpostal_setup_parser")
	}
}

func teardownParser() {
	C.libpostal_teardown_parser()
}

func ParseAddress(address string, options ParserOptions) ([]string, []string) {
	setupParser()

	if !utf8.ValidString(address) {
		return nil, nil
	}

	cAddress := C.CString(address)
	defer C.free(unsafe.Pointer(cAddress))

	cOptions := C.libpostal_get_address_parser_default_options()
	if options.Language != "" {
		cLanguage := C.CString(options.Language)
		defer C.free(unsafe.Pointer(cLanguage))

		cOptions.language = cLanguage
	}

	if options.Country != "" {
		cCountry := C.CString(options.Country)
		defer C.free(unsafe.Pointer(cCountry))

		cOptions.country = cCountry
	}

	cAddressParserResponsePtr := C.libpostal_parse_address(cAddress, cOptions)

	if cAddressParserResponsePtr == nil {
		return nil, nil
	}

	cAddressParserResponse := *cAddressParserResponsePtr

	cNumComponents := cAddressParserResponse.num_components
	cComponents := cAddressParserResponse.components
	cLabels := cAddressParserResponse.labels

	numComponents := uint64(cNumComponents)

	cComponentsPtr := (*[1 << 30](*C.char))(unsafe.Pointer(cComponents))
	cLabelsPtr := (*[1 << 30](*C.char))(unsafe.Pointer(cLabels))

	var labels []string
	var values []string
	var i uint64
	for i = 0; i < numComponents; i++ {
		labels = append(labels, C.GoString(cLabelsPtr[i]))
		values = append(values, C.GoString(cComponentsPtr[i]))
	}

	C.libpostal_address_parser_response_destroy(cAddressParserResponsePtr)

	return labels, values
}

func ParserPrintFeatures(b bool) bool {
	return bool(C.libpostal_parser_print_features(C.bool(b)))
}
