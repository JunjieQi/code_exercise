package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

}

func TestParseMentions(t *testing.T) {
	m := parseMentions("@bob @john (success) such a cool feature")

	expected := []string{"bob", "john"}

	assert.Equal(t, expected, m)

	// empty case
	m = parseMentions("@& such a cool feature;")
	assert.Empty(t, m)

}

func TestParseEmoticons(t *testing.T) {
	e := parseEmoticons("@bob @john (success) (smile) (thisisoverfifteensentence)")

	expected := []string{"success", "smile"}

	assert.Equal(t, expected, e)

	// empty case
	e = parseEmoticons("")
	assert.Empty(t, e)
}

func TestGetUrls(t *testing.T) {
	l := getUrls("Olympics are starting soon; http://www.nbcolympics.com; https://twitter.com/jdorfman/status/430511497475670016")

	expected := []string{"http://www.nbcolympics.com", "https://twitter.com/jdorfman/status/430511497475670016"}

	assert.Equal(t, expected, l)
	// empty case
	l = getUrls("Olympics are starting soon;")
	assert.Empty(t, l)
}

func TestGetTitle(t *testing.T) {
	title, err := getTitle("http://www.nbcolympics.com")

	assert.Nil(t, err)
	assert.NotEmpty(t, title)
}

func TestParseLinks(t *testing.T) {
	links := parseLinks("Olympics are starting soon; http://www.nbcolympics.com; https://twitter.com/jdorfman/status/430511497475670016")

	assert.Len(t, links, 2)
}
