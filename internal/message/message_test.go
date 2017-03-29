package message

import (
	"testing"
	"net/http"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
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
	defer gock.Off()

	gock.New("http://www.nbcolympics.com").
		Reply(http.StatusOK).
		BodyString("<title>hello</title>")

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	title, err := getTitle("http://www.nbcolympics.com", client)

	assert.Nil(t, err)
	assert.Equal(t, "hello", title)
}

func TestParseLinks(t *testing.T) {
	links := parseLinks("Olympics are starting soon; http://www.nbcolympics.com; https://twitter.com/jdorfman/status/430511497475670016")

	assert.Len(t, links, 2)
}
