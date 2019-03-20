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

type Place struct {
	Name          string
	HouseNumber   string
	Street        string
	Building      string
	Entrance      string
	Staircase     string
	Level         string
	Unit          string
	PoBox         string
	MetroStation  string
	Suburb        string
	CityDistrict  string
	City          string
	StateDistrict string
	Island        string
	State         string
	CountryRegion string
	Country       string
	WorldRegion   string
	PostalCode    string
	Telephone     string
	Website       string
}

func PlaceFromComponents(labels []string, values []string) (*Place, error) {
	if len(labels) != len(values) {
		return nil, fmt.Errorf("lables and values length must be equal")
	}

	if len(labels) == 0 {
		return nil, fmt.Errorf("lables is empty")
	}

	place := &Place{}

	for i := range labels {
		label := labels[i]
		value := values[i]

		switch label {
		case AddressLabelHouse:
			place.Name = value
		case AddressLabelHouseNumber:
			place.HouseNumber = value
		case AddressLabelPoBox:
			place.PoBox = value
		case AddressLabelBuilding:
			place.Building = value
		case AddressLabelEntrance:
			place.Entrance = value
		case AddressLabelStaircase:
			place.Staircase = value
		case AddressLabelLevel:
			place.Level = value
		case AddressLabelUnit:
			place.Unit = value
		case AddressLabelRoad:
			place.Street = value
		case AddressLabelMetroStation:
			place.MetroStation = value
		case AddressLabelSuburb:
			place.Suburb = value
		case AddressLabelCityDistrict:
			place.CityDistrict = value
		case AddressLabelCity:
			place.City = value
		case AddressLabelStateDistrict:
			place.StateDistrict = value
		case AddressLabelIsland:
			place.Island = value
		case AddressLabelState:
			place.State = value
		case AddressLabelPostalCode:
			place.PostalCode = value
		case AddressLabelCountryRegion:
			place.CountryRegion = value
		case AddressLabelCountry:
			place.Country = value
		case AddressLabelWorldRegion:
			place.WorldRegion = value
		case AddressLabelWebsite:
			place.Website = value
		case AddressLabelTelephone:
			place.Telephone = value
		default:
			fmt.Println("unsupported label:", label)
		}
	}

	return place, nil
}

func PlaceLanguages(labels []string, values []string) []string {
	cLabels := make([]*C.char, len(labels))
	for i, label := range labels {
		cLabel := C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
		cLabels[i] = cLabel
	}

	cValues := make([]*C.char, len(values))
	for i, value := range values {
		cValue := C.CString(value)
		defer C.free(unsafe.Pointer(cValue))
		cValues[i] = cValue
	}
	cNumComponents := C.ulong(len(labels))
	cNumLanguages := C.size_t(0)
	cLanguages := C.libpostal_place_languages(cNumComponents, &cLabels[0], &cValues[0], &cNumLanguages)
	cLanguagesPtr := (*[1 << 30](*C.char))(unsafe.Pointer(cLanguages))

	var languages []string
	var i uint64
	for i = 0; i < uint64(cNumLanguages); i++ {
		languages = append(languages, C.GoString(cLanguagesPtr[i]))
	}

	return languages
}
