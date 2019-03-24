package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestParseUSAddress(t *testing.T) {
	actual := postal.ParseAddress("781 Franklin Ave Crown Heights Brooklyn NYC NY 11216 USA", postal.DefaultParserOptions())
	expected := map[string]string{
		"country":       "usa",
		"house_number":  "781",
		"road":          "franklin ave",
		"suburb":        "crown heights",
		"city_district": "brooklyn",
		"city":          "nyc",
		"state":         "ny",
		"postcode":      "11216",
	}
	assert.Equal(t, expected, actual)
}

func TestParserPrintFeatures(t *testing.T) {
	b := postal.ParserPrintFeatures(true)
	assert.True(t, b)
}
