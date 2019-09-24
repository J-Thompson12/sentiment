package models

//CombinedResult combines SocialActivity and Enriched
type CombinedResult struct {
	Activity SocialActivity `json:"activity"`
	Enriched Enriched       `json:"enriched"`
}

//Enriched contains the enriched data for a SocialActivity
type Enriched struct {
	AuthorFollowersCount int         `json:"author_followers_count"`
	AuthorGender         Gender      `json:"author_gender"`
	Categories           []Category  `json:"categories"`
	Entities             []Entity    `json:"entities"`
	SentimentScore       float64     `json:"sentiment_score"`
	Themes               []Theme     `json:"themes"`
	Intentions           []Intention `json:"intentions"`
	UniqueHash           string      `json:"unique_hash"`
}

//Category contains information about the NLP Categories for a SocialActivity
type Category struct {
	Children       []Category `json:"children"`
	SentimentScore float64    `json:"sentiment_score"`
	Title          string     `json:"title"`
}

//Entity contains information about the entities found in a SocialActivity
type Entity struct {
	SentimentScore float64 `json:"sentiment_score"`
	Title          string  `json:"title"`
	Label          string  `json:"label"`
}

//Theme contains information about the themes found in a SocialActivity
type Theme struct {
	Title          string  `json:"title"`
	Normalized     string  `json:"normalized"`
	Stemmed        string  `json:"stemmed"`
	SentimentScore float64 `json:"sentiment_score"`
}

// Intention carries nlp derived intent from a body of text
type Intention struct {
	// Who had the intent
	Who string `json:"who"`
	// To What was the intent directed
	What string `json:"what"`
	// Type is the actual intent (buy, sell, recommend, etc)
	Type string `json:"type"`
}

//Gender contains information about the Gender of an author
type Gender struct {
	AuthorRealName string  `json:"author_real_name"`
	Male           float64 `json:"male"`
	Female         float64 `json:"female"`
}
