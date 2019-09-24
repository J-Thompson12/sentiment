package models

import (
	"time"
)

//SocialGistReview review structure from socialgist firehose
type SocialGistReview struct {
	Review struct {
		ID           string `json:"id"`
		Subject      string `json:"subject"`
		Text         string `json:"text"`
		TextPros     string `json:"textPros"`
		TextCons     string `json:"textCons"`
		ReviewDate   string `json:"reviewDate"`
		Crawled      string `json:"crawled"`
		ReviewType   string `json:"reviewType"`
		ReviewRating struct {
			RawScore  string `json:"RawScore"`
			RatingNum string `json:"ratingNum"`
		} `json:"reviewRating"`
		Recommend         string `json:"recommend"`
		NumOfHelpfulVotes string `json:"numOfHelpfulVotes"`
		AttributeRatings  struct {
		} `json:"attributeRatings"`
		Author                 string `json:"author"`
		AuthorName             string `json:"authorName"`
		AuthorURL              string `json:"authorURL"`
		AvatarURL              string `json:"avatarURL"`
		Location               string `json:"location"`
		Age                    string `json:"age"`
		Sex                    string `json:"sex"`
		AuthorDetails          string `json:"authorDetails"`
		URL                    string `json:"url"`
		SiteID                 string `json:"siteID"`
		SiteAlias              string `json:"siteAlias"`
		SiteURL                string `json:"siteURL"`
		SiteCountry            string `json:"siteCountry"`
		OrigSiteURL            string `json:"origSiteURL"`
		OrigURL                string `json:"origURL"`
		Language               string `json:"language"`
		ItemCategory           string `json:"itemCategory"`
		ItemCategoryURL        string `json:"itemCategoryURL"`
		Item                   string `json:"item"`
		ItemURL                string `json:"itemURL"`
		ItemOverallRating      string `json:"itemOverallRating"`
		ItemReviewCount        string `json:"itemReviewCount"`
		ItemPercentRecommended string `json:"itemPercentRecommended"`
		ItemProperties         struct {
			Address string `json:"address"`
			Author  string `json:"author"`
		} `json:"itemProperties"`
	} `json:"review"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialGistReview) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", socialgistPayload.Review.ReviewDate, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
