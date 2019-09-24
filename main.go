package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nuvi/justin-sentiment/postgres"
	"github.com/nuvi/justin-sentiment/rabbit"
	"github.com/nuvi/justin-sentiment/tfidf"
	"github.com/nuvi/justin-sentiment/train"
)

// parameters
var c train.Classifier

func main() {
	db := postgres.InitPostgres()
	db.EnsurePostgresTables()

	sentimentData := db.GetFaschetData("sentiment")

	sentimentCategories := []string{"positive", "neutral", "negative"}

	sentiment := train.CreateClassifier(sentimentCategories)

	//trainFile on trainFile dataset
	for _, data := range sentimentData {
		sent := ""
		switch data.Sentiment {
		case 1:
			sent = "negative"
		case 2:
			sent = "neutral"
		case 3:
			sent = "positive"
		}
		train.Train(&sentiment, sent, data.Text)
	}
	if os.Getenv("CLASSIFICATION_METHOD") == "tfidf" {
		for _, data := range sentimentData {
			sent := ""
			switch data.Sentiment {
			case 1:
				sent = "negative"
			case 2:
				sent = "neutral"
			case 3:
				sent = "positive"
			}
			tfidf.TfFreq(&sentiment, sent, data.Text)
		}
	}

	rabbit.Consume(&c)
}
