package classify

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {

	mentions := "@something@else#lookatme"
	cleanup(mentions)

	a := "@something #swagger Ohh look at me i'm testing emojis i'm freaking AWESOME!!! 😀 😁 😂 🤣 😃 😄"
	// U+1F600
	b := "Illuminati meeting, more hits on the way, be patient, excuse me. Lmao.🎶😂🎶 "
	c := "a #💩 #and #🍦 #😳"
	result := tokenize(a)
	fmt.Println(result)
	result = tokenize(b)
	fmt.Println(result)
	result = tokenize(c)
	fmt.Println(result)
}
