package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestIsDuplicate(t *testing.T) {
	status, err := postal.IsDuplicate(postal.AddressStreet, "6020 Churchland St #1", "6020 Churchland Blvd #1", postal.DefaultDuplicateOptions())
	assert.Nil(t, err)
	assert.Equal(t, postal.DuplicateStatusPossible, status)
	status, err = postal.IsDuplicate(postal.AddressName, "Home I", "Home 2", postal.DefaultDuplicateOptions())
	assert.Nil(t, err)
	assert.Equal(t, postal.DuplicateStatusNon, status)
}

func TestIsToponymDuplicate(t *testing.T) {
	comps1 := map[string]string{
		postal.AddressLabelRoad:    "426 East Beaver Creek Rd",
		postal.AddressLabelCity:    "Knoxville",
		postal.AddressLabelState:   "TN",
		postal.AddressLabelCountry: "USA",
		postal.AddressLabelHouse:   "home 1",
	}
	comps2 := map[string]string{
		postal.AddressLabelRoad:    "426 East Beaver Creek Rd",
		postal.AddressLabelCity:    "Knoxville",
		postal.AddressLabelState:   "TN",
		postal.AddressLabelCountry: "USA",
		postal.AddressLabelHouse:   "home1",
	}

	status := postal.IsToponymDuplicate(comps1, comps2, postal.DefaultDuplicateOptions())
	assert.Equal(t, postal.DuplicateStatusExact, status)
}

func TestIsNameDuplicateFuzzy(t *testing.T) {
	tokens1 := []string{"The", "Name 1"}
	scores1 := 1.0
	tokens2 := []string{"The", "Name 2"}
	scores2 := 1.0
	opts := postal.DefaultFuzzyDuplicateOptions()
	status, sim, err := postal.IsDuplicateFuzzy(postal.AddressName, tokens1, scores1, tokens2, scores2, opts)
	assert.Nil(t, err)
	assert.Equal(t, postal.DuplicateStatusLikely.String(), status.String())
	assert.Equal(t, 1.0, sim)
}

func TestIsStreetDuplicateFuzzy(t *testing.T) {
	tokens1 := []string{"East", "Beaver", "Creek", "Rd"}
	scores1 := 1.0
	tokens2 := []string{"East", "Beaver", "Creek", "Road"}
	scores2 := 1.0
	opts := postal.DefaultFuzzyDuplicateOptions()
	status, sim, err := postal.IsDuplicateFuzzy(postal.AddressStreet, tokens1, scores1, tokens2, scores2, opts)
	assert.Nil(t, err)
	assert.Equal(t, postal.DuplicateStatusPossible.String(), status.String())
	assert.Equal(t, 0.7071067811865475, sim)
}
