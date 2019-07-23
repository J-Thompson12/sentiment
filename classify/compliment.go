package classify

import (
	"fmt"
	"math"
)

func (c *Classifier) probabilitiesComp(document string) (p map[string]float64) {
	p = make(map[string]float64)

	for category := range c.words {
		p[category] = c.compliment(category, document) * c.categoryProb(category)
	}
	return
}

func (c *Classifier) compliment(category string, document string) (p float64) {
	p = 0.0
	words := countWords(document)
	for word := range words {
		x := c.wncAll(word)
		y := c.wnc(category, word)
		if x > 0 || y > 0 {
			fmt.Println("a number is greater")
		}

		p = p + (words[word] * c.wncAll(word)) + (words[word] * c.wnc(category, word))
	}
	return p
}

func (c *Classifier) wnc(category string, word string) float64 {
	sum := 0.0
	wordSum := 0.0
	for cat := range c.words {
		if cat != category {
			sum += c.wordCategoryTotal[cat]
			wordSum += c.words[cat][word]
		}
	}
	return math.Log((1 + wordSum) / (c.totalWords + sum))
}

func (c *Classifier) wncAll(word string) float64 {
	wordSum := 0.0
	for cat := range c.words {
		wordSum += c.words[cat][word]
	}
	return math.Log((1.0 + wordSum) / (c.totalWords * 2))
}
