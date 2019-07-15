package classify

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var positiveString1 = "that was a really good movie"
var positiveString2 = "I really liked the food here"
var negativeString1 = "that was the worst place i have ever been to"
var negativeString2 = "I cant believe how horrible this was"
var neutralString = "The food wasnt great but it was okay"

var dataFile = "../all.txt"
var categories = []string{"positive", "negative"}
var train []document

var testData = []string{positiveString1, positiveString2, negativeString1, negativeString2, neutralString}

// datasets
type document struct {
	sentiment string
	text      string
}

func TestCountWords(t *testing.T) {
	tokens, _ := countWords("she was not my favorite. my favorite was the main guy")
	assert.Equal(t, 6, len(tokens), "these should be equal")
}

func TestClassify(t *testing.T) {
	setupData(dataFile)
	c := CreateClassifier(categories)

	for _, doc := range train {
		c.Train(doc.sentiment, doc.text)
	}

	sentiment := c.Classify(positiveString1)
	assert.Equal(t, "positive", sentiment, "this should be positive")

	sentiment = c.Classify(positiveString2)
	assert.Equal(t, "positive", sentiment, "this should be positive")

	sentiment = c.Classify(negativeString1)
	assert.Equal(t, "negative", sentiment, "this should be negative")

	sentiment = c.Classify(negativeString2)
	assert.Equal(t, "negative", sentiment, "this should be negative")

	// sentiment = c.Classify(neutralString)
	// assert.Equal(t, "neutral", sentiment, "this should be neutral")
}

func setupData(file string) {
	data, err := readLines(file)
	if err != nil {
		fmt.Println("Cannot read file", err)
		os.Exit(1)
	}
	for _, line := range data {
		s := strings.Split(line, "|")
		doc, sentiment := s[0], s[1]

		train = append(train, document{sentiment, doc})
	}
}

// read the file line by line
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
