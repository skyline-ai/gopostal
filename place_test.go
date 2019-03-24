package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestPlaceLanguages(t *testing.T) {
	comps := map[string]string{
		postal.AddressLabelRoad:        "east beaver creek rd",
		postal.AddressLabelHouseNumber: "426",
		postal.AddressLabelCity:        "knoxville",
		postal.AddressLabelState:       "tn",
	}

	languages := postal.PlaceLanguages(comps)
	assert.Equal(t, []string{"en"}, languages)
}
