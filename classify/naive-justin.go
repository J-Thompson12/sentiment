package classify

import (
	"math"
	"sync"
)

const threshold = 1.1

var lock sync.RWMutex

// Classifier creates a struct with the following elements
// words is a map of maps that represent the words that have been trained by the classifier
// example
// {
//     "positive": {
//         "good": 10,
//         "wonderful": 5,
//         "amazing": 7,
//     },
//     "negative": {
//         "awful": 6,
//         "loud": 4,
//     }
// }
// documentCategoryTotal stores the number of documents in each category
// {
//     "positive": 13,
//     "negative": 16,
// }
// wordCategoryTotal stores the number of documents in each category
// {
//     "positive": 64,
//     "negative": 53,
// }
// numDocSeen stores the number of documents contains a specific word
type Classifier struct {
	words                 map[string]map[string]float64
	totalDocuments        float64
	documentCategoryTotal map[string]float64
	wordCategoryTotal     map[string]float64
	totalWords            float64
	numDocSeen            map[string]float64
}

// CreateClassifier create and initialize the classifier
func CreateClassifier(categories []string) (c Classifier) {
	c = Classifier{
		words:                 make(map[string]map[string]float64),
		documentCategoryTotal: make(map[string]float64),
		wordCategoryTotal:     make(map[string]float64),
		totalDocuments:        0,
		totalWords:            0,
		numDocSeen:            make(map[string]float64),
	}

	for _, category := range categories {
		c.words[category] = make(map[string]float64)
		c.documentCategoryTotal[category] = 0
		c.wordCategoryTotal[category] = 0
	}

	return
}

// Train the classifier
func (c *Classifier) Train(category string, data string) {
	words, _ := countWords(data)
	for word, count := range words {
		lock.Lock()
		c.words[category][word] += count
		c.wordCategoryTotal[category] += count
		c.numDocSeen[word]++
		c.totalWords += count
		lock.Unlock()
	}
	c.documentCategoryTotal[category]++
	c.totalDocuments++
}

// clean up and split words in DataSet, then stem each word and count the occurrence. This will split the words individually and float64o ngrams of 2.
func countWords(document string) (wordCount map[string]float64, totalWords float64) {
	words, count := tokenize(document)
	words = append(words, tokenizeMulti(document, 2)...)
	wordCount = make(map[string]float64)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount, count
}

// Classify a DataSet
func (c *Classifier) Classify(document string) string {
	var result string

	// Mulitnomial naive bayes
	prob := c.probMulti(document)
	highNum := 0.0
	for category, probability := range prob {
		if probability > highNum {
			highNum = probability
			result = category
		}
	}

	//compliment naive bayes uses lowest number to determine category
	// prob := c.probabilitiesComp(document)
	// lowNum := 99999999999990.0
	// for category, probability := range prob {
	// 	if probability < lowNum {
	// 		lowNum = probability
	// 		result = category
	// 	}
	// }

	return result
}

// Gets the probablity of the probability of each of the categories
func (c *Classifier) categoryProb(category string) float64 {
	return c.documentCategoryTotal[category] / c.totalDocuments
}

func (c *Classifier) tfidf(word string) float64 {
	return math.Log(c.totalDocuments / c.numDocSeen[word])
}
