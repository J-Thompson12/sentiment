package classify

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
// numDocSeen stores the number of documents contains a specific word
type Classifier struct {
	Words                 map[string]map[string]int
	TotalDocuments        int
	DocumentCategoryTotal map[string]int
	WordCategoryTotal     map[string]int
	TotalWords            int
	numDocSeen            map[string]int
}

// CreateClassifier create and initialize the classifier
func CreateClassifier(categories []string) (c Classifier) {
	c = Classifier{
		Words:                 make(map[string]map[string]int),
		DocumentCategoryTotal: make(map[string]int),
		WordCategoryTotal:     make(map[string]int),
		TotalDocuments:        0,
		TotalWords:            0,
		numDocSeen:            make(map[string]int),
	}

	for _, category := range categories {
		c.Words[category] = make(map[string]int)
		c.DocumentCategoryTotal[category] = 0
		c.WordCategoryTotal[category] = 0
	}

	return
}

// Train the classifier
func (c *Classifier) Train(category string, data string) {
	for word, count := range countWords(data) {
		c.Words[category][word] += count
		c.WordCategoryTotal[category] += count
		c.TotalWords += count
		c.numDocSeen[word]++
	}
	c.DocumentCategoryTotal[category]++
	c.TotalDocuments++
}

// clean up and split Words in DataSet, then stem each word and count the occurrence. This will split the words individually and into ngrams of 2.
func countWords(document string) (wordCount map[string]int) {
	words := tokenize(document)
	words = append(words, tokenizeMulti(document, 2)...)
	wordCount = make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return
}

// Classify a DataSet
func (c *Classifier) Classify(document string) string {
	// get all the probabilities of each category
	// this is used to sort based of -1 to 1
	// prob := c.probabilities(document)
	// value := prob[positive] + (prob[negative] * -1)
	// if value > 0 {
	// 	category = positive
	// } else if value < 0 {
	// 	category = negative
	// } else {
	// 	category = "neutral"
	// }

	prob := c.probabilities(document)

	//this picks the category with the largest property
	var result string
	highNum := 0.0
	for category, probability := range prob {
		if probability > highNum {
			highNum = probability
			result = category
		}
	}

	return result
}

// Using basic naive bayes equation to get the probability that the dataset belongs to a specific category
func (c *Classifier) probabilities(document string) (p map[string]float64) {
	p = make(map[string]float64)

	for category := range c.Words {
		p[category] = c.setCategory(category, document) * c.categoryProb(category)
	}
	return
}

// Checks each individual word to see what category they belong to and multiplies them
func (c *Classifier) setCategory(category string, document string) (p float64) {
	p = 1.0
	for word := range countWords(document) {
		p = p * c.wordProb(category, word)
	}
	return p
}

// Gets the probablity of the probability of each of the categories
func (c *Classifier) categoryProb(category string) float64 {
	return float64(c.DocumentCategoryTotal[category]) / float64(c.TotalDocuments)
}

// Gets the probablity of each word with a laplace estimator of +1 for smoothing
func (c *Classifier) wordProb(category string, word string) float64 {
	//return (TFIDF(word, document, c)) / (float64(c.WordCategoryTotal[category]) + float64(c.TotalWords))
	return (float64(c.Words[category][word]) + 1) / (float64(c.WordCategoryTotal[category]) + float64(c.TotalWords))
}
