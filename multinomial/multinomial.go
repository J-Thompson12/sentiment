package multinomial

import (
	"math"

	"github.com/nuvi/justin-sentiment/train"
	"github.com/nuvi/justin-sentiment/utilities"
)

// ProbMulti Using basic naive bayes equation to get the probability that the dataset belongs to a specific category
func ProbMulti(c *train.Classifier, document string) (p map[string]float64) {
	p = make(map[string]float64)

	for category := range c.Words {
		p[category] = wordFrequency(c, category, document) * categoryProb(c, category)
	}
	return
}

//TfFreq gives the word frequency using TFIDF
func TfFreq(c *train.Classifier, category string, data string) {
	words := utilities.CountWords(data)
	for word, count := range words {
		c.NormalFreq[category][word] += normFreq(c, count, len(words), word)
	}
}

// Checks each individual word to see what category they belong to and multiplies them
func wordFrequency(c *train.Classifier, category string, document string) (p float64) {
	p = 1.0
	wnc := 0.0
	words := utilities.CountWords(document)
	for word := range words {
		if c.HasIDF {
			wnc = (c.NormalFreq[category][word] + 1.0) / (c.WordCategoryTotal[category] + c.TotalWords)
		} else {
			wnc = (c.Words[category][word] + 1.0) / (c.WordCategoryTotal[category] + c.TotalWords)
		}
		p = p * wnc
	}

	return p
}

func normFreq(c *train.Classifier, wordCount float64, totalWords int, word string) float64 {
	return (wordCount / float64(totalWords)) * idf(c, word)
}

func idf(c *train.Classifier, word string) float64 {
	return math.Log(c.TotalDocuments / c.NumDocSeen[word])
}

func categoryProb(c *train.Classifier, category string) float64 {
	return c.DocumentCategoryTotal[category] / c.TotalDocuments
}
