package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"

const (
	AddressNone        uint16 = C.LIBPOSTAL_ADDRESS_NONE
	AddressAny         uint16 = C.LIBPOSTAL_ADDRESS_ANY
	AddressName        uint16 = C.LIBPOSTAL_ADDRESS_NAME
	AddressHouseNumber uint16 = C.LIBPOSTAL_ADDRESS_HOUSE_NUMBER
	AddressStreet      uint16 = C.LIBPOSTAL_ADDRESS_STREET
	AddressUnit        uint16 = C.LIBPOSTAL_ADDRESS_UNIT
	AddressLevel       uint16 = C.LIBPOSTAL_ADDRESS_LEVEL
	AddressStaircase   uint16 = C.LIBPOSTAL_ADDRESS_STAIRCASE
	AddressEntrance    uint16 = C.LIBPOSTAL_ADDRESS_ENTRANCE
	AddressCategory    uint16 = C.LIBPOSTAL_ADDRESS_CATEGORY
	AddressNear        uint16 = C.LIBPOSTAL_ADDRESS_NEAR
	AddressToponym     uint16 = C.LIBPOSTAL_ADDRESS_TOPONYM
	AddressPostalCode  uint16 = C.LIBPOSTAL_ADDRESS_POSTAL_CODE
	AddressPoBox       uint16 = C.LIBPOSTAL_ADDRESS_PO_BOX
	AddressAll         uint16 = C.LIBPOSTAL_ADDRESS_ALL
)

func AddressComponentString(a uint16) string {
	switch a {
	case AddressNone:
		return "None"
	case AddressAny:
		return "Any"
	case AddressName:
		return "Name"
	case AddressHouseNumber:
		return "HouseNumber"
	case AddressStreet:
		return "Street"
	case AddressUnit:
		return "Unit"
	case AddressLevel:
		return "Level"
	case AddressStaircase:
		return "Staircase"
	case AddressEntrance:
		return "Entrance"
	case AddressCategory:
		return "Category"
	case AddressNear:
		return "Near"
	case AddressToponym:
		return "Toponym"
	case AddressPostalCode:
		return "PostalCode"
	case AddressPoBox:
		return "PoBox"
	case AddressAll:
		return "All"
	}

	return ""
}
