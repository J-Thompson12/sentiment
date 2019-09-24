package train

import "github.com/nuvi/sentiment-trainer-api/utilities"

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
	Words                 map[string]map[string]float64
	TotalDocuments        float64
	DocumentCategoryTotal map[string]float64
	WordCategoryTotal     map[string]float64
	TotalWords            float64
	NumDocSeen            map[string]float64
	NormalFreq            map[string]map[string]float64
}

// CreateClassifier create and initialize the classifier
func CreateClassifier(categories []string) (c Classifier) {
	c = Classifier{
		Words:                 make(map[string]map[string]float64),
		DocumentCategoryTotal: make(map[string]float64),
		WordCategoryTotal:     make(map[string]float64),
		TotalDocuments:        0,
		TotalWords:            0,
		NumDocSeen:            make(map[string]float64),
		NormalFreq:            make(map[string]map[string]float64),
	}

	for _, category := range categories {
		c.Words[category] = make(map[string]float64)
		c.NormalFreq[category] = make(map[string]float64)
		c.DocumentCategoryTotal[category] = 0
		c.WordCategoryTotal[category] = 0
	}

	return
}

// Train the classifier
func Train(c *Classifier, category string, data string) {
	words := utilities.CountWords(data)
	for word, count := range words {
		c.Words[category][word] += count
		c.WordCategoryTotal[category] += count
		c.NumDocSeen[word]++
		c.TotalWords += count
	}
	c.DocumentCategoryTotal[category]++
	c.TotalDocuments++
}
