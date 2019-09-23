package multinomial

import (
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

// Checks each individual word to see what category they belong to and multiplies them
func wordFrequency(c *train.Classifier, category string, document string) (p float64) {
	p = 1.0
	wnc := 0.0
	words := utilities.CountWords(document)
	for word := range words {
		wnc = (c.Words[category][word] + 1.0) / (c.WordCategoryTotal[category] + c.TotalWords)
		p = p * wnc
	}
	return p
}

func categoryProb(c *train.Classifier, category string) float64 {
	return c.DocumentCategoryTotal[category] / c.TotalDocuments
}
