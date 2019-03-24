package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestNearDupeHashes(t *testing.T) {
	comps := map[string]string{
		postal.AddressLabelRoad:        "east beaver creek rd",
		postal.AddressLabelHouseNumber: "426",
		postal.AddressLabelCity:        "knoxville",
		postal.AddressLabelState:       "tn",
	}

	opts := postal.DefaultNearDupeHashOptions()
	opts.AddressOnlyKeys = true
	opts.WithLatLon = true
	opts.Latitude = 35.85821
	opts.Longitude = -84.08088
	opts.Languages = []string{"en"}
	hashes := postal.NearDupeHashes(comps, opts)
	assert.Len(t, hashes, 20)
}
