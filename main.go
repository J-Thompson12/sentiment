package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	classify "github.com/nuvi/justin-sentiment/classify"
)

var categories = []string{"positive", "negative"}

// parameters
var dataFile = "databig.txt"
var train []document
var test []document
var testPercentage = 0.2
var wg sync.WaitGroup
var count, accurates, neutral = 0, 0, 0

// datasets
type document struct {
	sentiment string
	text      string
}

func main() {
	hasIDF := false
	start := time.Now()
	setupData(dataFile)
	fmt.Println("number of docs in TRAIN dataset:", len(train))
	fmt.Println("number of docs in TEST dataset:", len(test))
	c := classify.CreateClassifier(categories, hasIDF)

	// train on train dataset
	// for _, doc := range train {
	// 	c.Train(doc.sentiment, doc.text)
	// }
	// if hasIDF {
	// 	for _, doc := range train {
	// 		c.WordFreq(doc.sentiment, doc.text)
	// 	}
	// }

	// for _, doc := range train {
	// 	classify.RedisTrain(doc.sentiment, doc.text)
	// }

	// fmt.Println("Finished training")

	// Test individual user entered sentences
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("Enter sentence: ")
	// 	text, _ := reader.ReadString('\n')
	// 	sentiment := c.Classify(text)
	// 	fmt.Printf("This sentence is %v\n", sentiment)
	// }

	// validate dataset
	wg.Add(len(test))
	for _, doc := range test {
		count++
		validate(doc.text, doc.sentiment, &c)
	}

	wg.Wait()
	fmt.Printf("Accuracy on TEST dataset is %2.1f%% With %2.1f%% as neutral\n", float64(accurates)*100/float64(count), float64(neutral)*100/float64(count))
	elapsed := time.Since(start)
	log.Printf("runtime was%s", elapsed)
}

func validate(text string, docSent string, c *classify.Classifier) {
	defer wg.Done()
	sentiment := c.Classify(text)
	if sentiment == docSent {
		accurates++
	}
	if sentiment == "neutral" {
		neutral++
	}
}

func setupData(file string) {
	rand.Seed(time.Now().UTC().UnixNano())
	data, err := readLines(file)
	if err != nil {
		fmt.Println("Cannot read file", err)
		os.Exit(1)
	}
	for _, line := range data {
		s := strings.Split(line, "|")
		//fmt.Println(s)
		doc, sentiment := s[3], s[1]

		if rand.Float64() > testPercentage {
			train = append(train, document{sentiment, doc})
		} else {
			test = append(test, document{sentiment, doc})
		}
	}
}

// read the file line by line
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
