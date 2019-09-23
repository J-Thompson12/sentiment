package compliment

import (
	"math"

	"github.com/nuvi/justin-sentiment/train"
	"github.com/nuvi/justin-sentiment/utilities"
)

// ProbabilitiesComp get the compliment naive bayes score
func ProbabilitiesComp(c *train.Classifier, document string) (p map[string]float64) {
	p = make(map[string]float64)

	for category := range c.Words {
		p[category] = compliment(c, category, document) * categoryProb(c, category)
	}
	return
}

func compliment(c *train.Classifier, category string, document string) (p float64) {
	p = 0.0
	words := utilities.CountWords(document)
	for word := range words {
		p = p + (words[word] * wncAll(c, word)) + (words[word] * wnc(c, category, word))
	}
	return p
}

func wnc(c *train.Classifier, category string, word string) float64 {
	sum := 0.0
	wordSum := 0.0
	for cat := range c.Words {
		if cat != category {
			sum += c.WordCategoryTotal[cat]
			wordSum += c.Words[cat][word]
		}
	}
	return math.Log((1 + wordSum) / (c.TotalWords + sum))
}

func wncAll(c *train.Classifier, word string) float64 {
	wordSum := 0.0
	for cat := range c.Words {
		wordSum += c.Words[cat][word]
	}
	return math.Log((1.0 + wordSum) / (c.TotalWords * 2))
}

func categoryProb(c *train.Classifier, category string) float64 {
	return c.DocumentCategoryTotal[category] / c.TotalDocuments
}
