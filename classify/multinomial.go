package classify

// Using basic naive bayes equation to get the probability that the dataset belongs to a specific category
func (c *Classifier) probMulti(document string) (p map[string]float64) {
	p = make(map[string]float64)

	for category := range c.words {
		p[category] = (c.wordFrequency(category, document) * c.categoryProb(category))
	}
	return
}

// Checks each individual word to see what category they belong to and multiplies them
func (c *Classifier) wordFrequency(category string, document string) (p float64) {
	p = 1.0
	words, totalWords := countWords(document)
	for word := range words {
		p = p * (((c.words[category][word] + 1.0) * idf(words[word], totalWords)) / (c.wordCategoryTotal[category] + 1.0 + c.totalWords))
	}

	return p
}

func idf(wordCount float64, totalWords float64) float64 {
	return (wordCount + 1.0) / totalWords
}

func (c *Classifier) normalFreq(category string, word string) float64 {
	return (c.words[category][word] + 1.0) / c.wordCategoryTotal[category]
}
