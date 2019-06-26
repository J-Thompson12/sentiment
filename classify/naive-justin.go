package classify

const (
	positive = "positive"
	negative = "negative"
)

// Classifier creates a struct with the following elements
// Words is a map of maps that represent the words that have been trained by the classifier
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
// DocumentCategoryTotal stores the number of documents in each category
// {
//     "positive": 13,
//     "negative": 16,
// }
// WordCategoryTotal stores the number of documents in each category
// {
//     "positive": 64,
//     "negative": 53,
// }
type Classifier struct {
	Words                 map[string]map[string]int
	TotalDocuments        int
	DocumentCategoryTotal map[string]int
	WordCategoryTotal     map[string]int
	TotalWords            int
}

// CreateClassifier create and initialize the classifier
func CreateClassifier() (c Classifier) {
	c = Classifier{
		Words:                 make(map[string]map[string]int),
		DocumentCategoryTotal: make(map[string]int),
		WordCategoryTotal:     make(map[string]int),
		TotalDocuments:        0,
		TotalWords:            0,
	}

	c.Words[positive] = make(map[string]int)
	c.Words[negative] = make(map[string]int)
	c.DocumentCategoryTotal[positive] = 0
	c.DocumentCategoryTotal[negative] = 0
	c.WordCategoryTotal[positive] = 0
	c.WordCategoryTotal[negative] = 0

	return
}

// Train the classifier
func (c *Classifier) Train(category string, data string) {
	for word, count := range countWords(data) {
		c.Words[category][word] += count
		c.WordCategoryTotal[category] += count
		c.TotalWords += count
	}
	c.DocumentCategoryTotal[category]++
	c.TotalDocuments++
}

// clean up and split Words in DataSet, then stem each word and count the occurrence
func countWords(document string) (wordCount map[string]int) {
	words := tokenize(document)
	wordCount = make(map[string]int)
	for _, word := range words {
		// key := stem(strings.ToLower(word))
		wordCount[word]++
	}
	return
}

// Classify a DataSet
func (c *Classifier) Classify(document string) (category string) {
	// get all the probabilities of each category
	prob := c.probabilities(document)
	if prob[positive] > prob[negative] {
		category = positive
	} else if prob[positive] < prob[negative] {
		category = negative
	} else {
		category = "neutral"
	}

	return
}

func (c *Classifier) probabilities(document string) (p map[string]float64) {
	p = make(map[string]float64)
	for category := range c.Words {
		p[category] = c.setCategory(category, document)
	}
	return
}

func (c *Classifier) setCategory(category string, document string) (p float64) {
	p = 1.0
	for word := range countWords(document) {
		p = p * c.pWordCategory(category, word)
	}
	return p
}

func (c *Classifier) pCategory(category string) float64 {
	return float64(c.DocumentCategoryTotal[category]) / float64(c.TotalDocuments)
}

func (c *Classifier) pWordCategory(category string, word string) float64 {
	return float64(c.Words[category][word]+1) / float64(c.WordCategoryTotal[category]+14)
}

func (c *Classifier) pCategoryDocument(category string, document string) float64 {
	return c.setCategory(category, document) * c.pCategory(category)
}
