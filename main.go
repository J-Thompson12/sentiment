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

	"github.com/nuvi/justin-sentiment/compliment"
	"github.com/nuvi/justin-sentiment/multinomial"
	"github.com/nuvi/justin-sentiment/train"
)

// parameters
var dataFile = "databig.txt"
var trainFile []document
var testFile []document
var testFilePercentage = 0.2
var wg sync.WaitGroup
var count, compAccurates, multiAccurates, neutral = 0, 0, 0, 0
var c train.Classifier
var hasIDF = false
var hasCompliment = true
var hasMulti = true

// datasets
type document struct {
	sentiment string
	text      string
}

func main() {
	start := time.Now()
	setupData(dataFile)
	fmt.Println("number of docs in trainFile dataset:", len(trainFile))
	fmt.Println("number of docs in testFile dataset:", len(testFile))
	c = train.CreateClassifier()

	//trainFile on trainFile dataset
	for _, doc := range trainFile {
		train.Train(&c, doc.sentiment, doc.text)
	}
	if hasIDF {
		for _, doc := range trainFile {
			multinomial.TfFreq(&c, doc.sentiment, doc.text)
		}
	}

	// validate dataset
	//wg.Add(len(testFile))
	for _, doc := range testFile {
		count++
		comp, multi := classify(doc.text)
		if comp == doc.sentiment {
			compAccurates++
		}
		if multi == doc.sentiment {
			multiAccurates++
		}

	}

	//wg.Wait()
	fmt.Printf("Accuracy on multi dataset is %2.1f%% With %2.1f%% as neutral\n", float64(multiAccurates)*100/float64(count), float64(neutral)*100/float64(count))
	fmt.Printf("Accuracy on comp dataset is %2.1f%% With %2.1f%% as neutral\n", float64(compAccurates)*100/float64(count), float64(neutral)*100/float64(count))
	elapsed := time.Since(start)
	log.Printf("runtime was%s", elapsed)
}

func classify(document string) (string, string) {
	var compResult string
	var multiResult string

	// Mulitnomial naive bayes
	if hasMulti {
		prob := multinomial.ProbMulti(&c, document)
		highNum := 0.0
		for category, probability := range prob {
			if probability >= highNum {
				highNum = probability
				multiResult = category
			}
		}
	}

	//compliment naive bayes uses lowest number to determine category
	if hasCompliment {
		prob := compliment.ProbabilitiesComp(&c, document)
		lowNum := 99999999999990.0
		for category, probability := range prob {
			if probability < lowNum {
				lowNum = probability
				compResult = category
			}
		}
	}

	return compResult, multiResult
}

func setupData(file string) {
	//rand.Seed(time.Now().UTC().UnixNano())
	data, err := readLines(file)
	if err != nil {
		fmt.Println("Cannot read file", err)
		os.Exit(1)
	}
	for _, line := range data {
		s := strings.Split(line, "|")
		//fmt.Println(s)
		doc, sentiment := s[3], s[1]

		if rand.Float64() > testFilePercentage {
			trainFile = append(trainFile, document{sentiment, doc})
		} else {
			testFile = append(testFile, document{sentiment, doc})
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
