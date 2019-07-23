package classify

import (
	"math"
)

// Using basic naive bayes equation to get the probability that the dataset belongs to a specific category
func (c *Classifier) probMulti(document string) (p map[string]float64) {
	p = make(map[string]float64)

	for category := range c.words {
		p[category] = c.wordFrequency(category, document) * c.categoryProb(category)
	}
	return
}

// Checks each individual word to see what category they belong to and multiplies them
func (c *Classifier) wordFrequency(category string, document string) (p float64) {
	p = 1.0
	wnc := 0.0
	words := countWords(document)
	for word := range words {
		if c.hasIDF {
			wnc = (c.normalFreq[category][word]) / (c.wordCategoryTotal[category] + c.totalWords)
		} else {
			wnc = (c.words[category][word] + 1.0) / (c.wordCategoryTotal[category] + c.totalWords)
		}

		p = p * wnc
	}

	return p
}

func (c *Classifier) normFreq(wordCount float64, totalWords int, word string) float64 {
	return ((wordCount + 1.0) / float64(totalWords)) * c.idf(word)
}

func (c *Classifier) idf(word string) float64 {
	return math.Log(c.totalDocuments / c.numDocSeen[word])
}
