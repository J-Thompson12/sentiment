package utilities

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {

	mentions := "@something@else#lookatme"
	cleanup(mentions)

	a := "@something #swagger Ohh look at me i'm testing emojis i'm freaking AWESOME!!! 😀 😁 😂 🤣 😃 😄"
	// U+1F600
	b := "Illuminati meeting, more hits on the way, be patient, excuse me. Lmao.🎶😂🎶 "
	c := "a #💩 #and #🍦 #😳"
	result := Tokenize(a)
	fmt.Println(result)
	result = Tokenize(b)
	fmt.Println(result)
	result = Tokenize(c)
	fmt.Println(result)
}

func TestCountWords(t *testing.T) {
	tokens := CountWords("she was not my favorite. my favorite was the main guy")
	assert.Equal(t, 6, len(tokens), "these should be equal")
}
