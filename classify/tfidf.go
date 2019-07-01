package classify

import (
	"math"
)

// TFIDF is used to get the weight of a word in a document. Currently a work in progress as it drops the accuracy by 30%
func TFIDF(word string, sentence string, c *Classifier) float64 {
	document := countWords(sentence)
	x := termFrequency(word, document)
	y := inverseDocumentFrequency(word, c)
	return x * y
}

func termFrequency(word string, document map[string]int) float64 {
	rawFreq := float64(document[word]) / float64(len(document))
	var maxFreq int
	for i := range document {
		if document[i] > maxFreq {
			maxFreq = document[i]
		}
	}

	return 0.5 + 0.5*(rawFreq/float64(maxFreq))
}

func inverseDocumentFrequency(word string, c *Classifier) float64 {
	return math.Log((float64(c.TotalDocuments) + 1) / float64(c.numDocSeen[word]))
}
