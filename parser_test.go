package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestParseUSAddress(t *testing.T) {
	labels, values := postal.ParseAddress("781 Franklin Ave Crown Heights Brooklyn NYC NY 11216 USA", postal.DefaultParserOptions())
	assert.Equal(t, []string{"house_number", "road", "suburb", "city_district", "city", "state", "postcode", "country"}, labels)
	assert.Equal(t, []string{"781", "franklin ave", "crown heights", "brooklyn", "nyc", "ny", "11216", "usa"}, values)
}

func TestParserPrintFeatures(t *testing.T) {
	b := postal.ParserPrintFeatures(true)
	assert.True(t, b)
}
