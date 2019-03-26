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
	scores1 := []float64{1.0, 1.0}
	tokens2 := []string{"The", "Name 2"}
	scores2 := []float64{1.0, 1.0}
	opts := postal.DefaultFuzzyDuplicateOptions()
	status, _, err := postal.IsDuplicateFuzzy(postal.AddressName, tokens1, scores1, tokens2, scores2, opts)
	assert.Nil(t, err)
	assert.Equal(t, postal.DuplicateStatusLikely.String(), status.String())
}

func TestIsStreetDuplicateFuzzy(t *testing.T) {
	tokens1 := []string{"Creekridge", "Road"}
	scores1 := []float64{1.0, 1.0}
	tokens2 := []string{"Creek", "ridge", "Road"}
	scores2 := []float64{1.0, 1.0, 1.0}
	opts := postal.DefaultFuzzyDuplicateOptions()
	status, _, err := postal.IsDuplicateFuzzy(postal.AddressStreet, tokens1, scores1, tokens2, scores2, opts)
	assert.Nil(t, err)
	assert.Equal(t, postal.DuplicateStatusPossible.String(), status.String())
}
