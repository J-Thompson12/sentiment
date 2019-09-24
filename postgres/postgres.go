package postgres

import (
	"fmt"
	"os"

	logit "github.com/brettallred/go-logit"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // This is needed for gorm to know to use postgres
)

//Classifications is a database model
type Classifications struct {
	Text      string `json:"text"`
	Sentiment int    `json:"sentiment"`
}

// DB is a db connection
type DB struct {
	DB *gorm.DB
}

//InitPostgres initializes postgres
func InitPostgres() DB {
	postgresqlURL := os.Getenv("POSTGRES_URL")
	if postgresqlURL == "" {
		logit.Fatal("No POSTGRES_URL set in .env")
	}

	db, err := gorm.Open("postgres", postgresqlURL)
	if err != nil {
		logit.Fatal(err.Error())
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	return DB{
		DB: db,
	}
}

//EnsurePostgresTables ensures the table
func (db DB) EnsurePostgresTables() {
	if os.Getenv("PLATFORM_ENV") == "test" {
		db.DB.DropTableIfExists(&Classifications{})
	}
	db.DB.CreateTable(&Classifications{})
	db.DB.AutoMigrate(&Classifications{})
}

//GetFaschetData gets all entries in the table where a specific Faschet is not null
func (db DB) GetFaschetData(faschet string) (allData []Classifications) {
	query := fmt.Sprintf("SELECT * FROM classifications WHERE %v IS NOT NULL", faschet)
	db.DB.Raw(query).Scan(&allData)
	return allData
}
