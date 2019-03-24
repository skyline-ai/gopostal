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

func PlaceFromComponents(comps map[string]string) (*Place, error) {
	if len(comps) == 0 {
		return nil, nil
	}

	place := &Place{}

	for label, value := range comps {
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

func PlaceLanguages(comps map[string]string) []string {
	var cLabels, cValues []*C.char

	for label, value := range comps {
		cLabel := C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
		cLabels = append(cLabels, cLabel)

		cValue := C.CString(value)
		defer C.free(unsafe.Pointer(cValue))
		cValues = append(cValues, cValue)
	}

	cNumComponents := C.ulong(len(comps))
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
