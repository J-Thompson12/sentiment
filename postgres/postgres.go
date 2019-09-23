package postgres

import (
	"os"

	logit "github.com/brettallred/go-logit"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // This is needed for gorm to know to use postgres
)

// SentimentData will contain all the data for the trainer
type SentimentData struct {
	Text      string `json:"text"`
	Sentiment string `json:"sentiment"`
	Category  string `json:"category"`
	Sarcastic string `json:"sarcastic"`
	Language  string
}

var db *gorm.DB

//InitPostgres initializes postgres
func InitPostgres() {
	logit.Info("Initializing postgres")
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		logit.Fatal("No POSTGRES_URL set in .env")
	}

	var err error
	db, err = gorm.Open("postgres", postgresURL)
	if err != nil {
		logit.Fatal(err.Error())
	}

	err = db.DB().Ping()
	if err != nil {
		logit.Fatal(err.Error())
	}
	logit.Info("Postgres is connected")
	ensurePostgresTables()
}

func ensurePostgresTables() {
	if os.Getenv("PLATFORM_ENV") == "test" {
		db.DropTableIfExists(&SentimentData{})
	}
	db.CreateTable(&SentimentData{})
	db.AutoMigrate(&SentimentData{})
}

//GetData gets all entries in the table
func GetData() (allData []SentimentData) {
	db.Find(&allData)
	return allData
}
