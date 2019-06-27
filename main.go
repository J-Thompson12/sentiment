package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	classify "github.com/nuvi/justin-sentiment/classify"
)

// parameters
var dataFile = "all.txt"
var train []document
var test []document
var testPercentage = 0.1

// datasets
type document struct {
	sentiment string
	text      string
}

func main() {
	setupData(dataFile)
	fmt.Println("number of docs in TRAIN dataset:", len(train))
	fmt.Println("number of docs in TEST dataset:", len(test))
	c := classify.CreateClassifier()

	// train on train dataset
	for _, doc := range train {
		c.Train(doc.sentiment, doc.text)
	}

	// validate dataset
	count, accurates, neutral := 0, 0, 0
	for _, doc := range test {
		count++
		sentiment := c.Classify(doc.text)
		if sentiment == doc.sentiment {
			accurates++
		}
		if sentiment == "neutral" {
			neutral++
		}
		// fmt.Println(doc.text + " |" + doc.sentiment + " --------- " + sentiment)
	}
	fmt.Printf("Accuracy on TEST dataset is %2.1f%% With %2.1f%% as neutral\n", float64(accurates)*100/float64(count), float64(neutral)*100/float64(count))
	// fmt.Println(c.Words)
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
		doc, sentiment := s[0], s[1]

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
