package models

import (
	"time"
)

//SocialGistReddit firehose payload for reddit
type SocialGistReddit struct {
	Report struct {
		Boardname     string `json:"boardname"`
		Siteid        string `json:"siteid"`
		Forumid       string `json:"forumid"`
		Forumname     string `json:"forumname"`
		Forumurl      string `json:"forumurl"`
		ThreadID      string `json:"ThreadID"`
		Threadurl     string `json:"threadurl"`
		Threadtitle   string `json:"threadtitle"`
		PostID        string `json:"postid"`
		Sitepostid    string `json:"sitepostid"`
		ParentID      string `json:"parentid"`
		URL           string `json:"Url"`
		Urlwithanchor string `json:"Urlwithanchor"`
		Anchor        string `json:"anchor"`
		MainURL       string `json:"MainUrl"`
		Content       struct {
			Date       string `json:"date"`
			Subject    string `json:"subject"`
			Author     string `json:"author"`
			AuthorURL  string `json:"AuthorUrl"`
			Avatarurl  string `json:"avatarurl"`
			Registered string `json:"registered"`
			Location   string `json:"location"`
			Age        string `json:"age"`
			Sex        string `json:"sex"`
			Text       string `json:"text"`
			Htmltext   string `json:"htmltext"`
		} `json:"content"`
		Topics         string `json:"topics"`
		Categories     string `json:"categories"`
		Crawled        string `json:"Crawled"`
		Language       string `json:"language"`
		LanguageCode   string `json:"languageCode"`
		Threadstarter  int    `json:"threadstarter"`
		Sentiment      string `json:"sentiment"`
		Recommendation string `json:"recommendation"`
		Ticker         string `json:"ticker"`
		Gmt            string `json:"gmt"`
		Signature      string `json:"signature"`
		Countrycode    string `json:"countrycode"`
		Providerid     string `json:"providerid"`
	} `json:"report"`
}

// CreationTime parses the creation time of reddit posts from social gist
func (redditFirehose *SocialGistReddit) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", redditFirehose.Report.Content.Date, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
