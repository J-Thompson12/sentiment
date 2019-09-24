package compliment

import (
	"math"

	"github.com/nuvi/justin-sentiment/train"
	"github.com/nuvi/justin-sentiment/utilities"
)

func ClassifyComp(c *train.Classifier, document string) (result string, score float64) {

	//compliment naive bayes uses lowest number to determine category
	prob := probComp(c, document)
	score = 99999999999990.0
	for category, probability := range prob {
		if probability < score {
			score = probability
			result = category
		}
	}
	return result, score
}

// ProbabilitiesComp get the compliment naive bayes score
func probComp(c *train.Classifier, document string) (p map[string]float64) {
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
		p = p + (words[word] * frequencyAllCatagories(c, word)) + (words[word] * frequencyAllButGiven(c, category, word))
	}
	return p
}

func frequencyAllButGiven(c *train.Classifier, category string, word string) float64 {
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

func frequencyAllCatagories(c *train.Classifier, word string) float64 {
	wordSum := 0.0
	for cat := range c.Words {
		wordSum += c.Words[cat][word]
	}
	return math.Log((1.0 + wordSum) / (c.TotalWords * 2))
}

func categoryProb(c *train.Classifier, category string) float64 {
	return c.DocumentCategoryTotal[category] / c.TotalDocuments
}
