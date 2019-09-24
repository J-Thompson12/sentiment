package multinomial

import (
	"github.com/nuvi/justin-sentiment/train"
	"github.com/nuvi/justin-sentiment/utilities"
)

func ClassifyMulti(c *train.Classifier, document string) (result string, score float64) {
	prob := probMulti(c, document)
	score = 0.0
	for category, probability := range prob {
		if probability >= score {
			score = probability
			result = category
		}
	}
	return result, score
}

// probMulti Using basic naive bayes equation to get the probability that the dataset belongs to a specific category
func probMulti(c *train.Classifier, document string) (p map[string]float64) {
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
