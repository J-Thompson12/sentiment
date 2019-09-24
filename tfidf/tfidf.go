package tfidf

import (
	"math"

	"github.com/nuvi/justin-sentiment/train"
	"github.com/nuvi/justin-sentiment/utilities"
)

func ClassifyTF(c *train.Classifier, document string) (result string, score float64) {
	prob := probTf(c, document)
	score = 0.0
	for category, probability := range prob {
		if probability >= score {
			score = probability
			result = category
		}
	}
	return result, score
}

// ProbTf Using basic naive bayes equation to get the probability that the dataset belongs to a specific category
func probTf(c *train.Classifier, document string) (p map[string]float64) {
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
		wnc = (c.NormalFreq[category][word] + 1.0) / (c.WordCategoryTotal[category] + c.TotalWords)
		p = p * wnc
	}
	return p
}

//TfFreq sets the normalized frequency for each word
func TfFreq(c *train.Classifier, category string, data string) {
	words := utilities.CountWords(data)
	for word, count := range words {
		c.NormalFreq[category][word] += normFreq(c, count, len(words), word)
	}
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
