package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var positiveString1 = "that was a really good movie"
var positiveString2 = "I really liked the food here"
var negativeString1 = "that was the worst place i have ever been to"
var negativeString2 = "I cant believe how horrible this was"
var neutralString = "The food wasnt great but it was okay"

var testData = []string{positiveString1, positiveString2, negativeString1, negativeString2, neutralString}

func TestClassify(t *testing.T) {
	setupData(dataFile)

	sentiment := classify(positiveString1)
	assert.Equal(t, "positive", sentiment, "this should be positive")

	sentiment = classify(positiveString2)
	assert.Equal(t, "positive", sentiment, "this should be positive")

	sentiment = classify(negativeString1)
	assert.Equal(t, "negative", sentiment, "this should be negative")

	sentiment = classify(negativeString2)
	assert.Equal(t, "negative", sentiment, "this should be negative")

	// sentiment = c.Classify(neutralString)
	// assert.Equal(t, "neutral", sentiment, "this should be neutral")
}
