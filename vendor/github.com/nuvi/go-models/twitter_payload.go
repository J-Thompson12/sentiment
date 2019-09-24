package models

import (
	"time"
)

/* see http://support.gnip.com/sources/twitter/data_format.html
   http://support.gnip.com/doing-more-with-140.html
   http://support.gnip.com/articles/identifying-and-understanding-retweets.html (for twitter_quoted_status)
*/

//TwitterPayload is the data structure based off the format of incoming JSON payloads
type TwitterPayload struct {
	ID    string `json:"id"`
	Actor struct {
		ObjectType  string `json:"objectType"`
		ID          string `json:"id"`
		Link        string `json:"link"`
		DisplayName string `json:"displayName"`
		PostedTime  string `json:"postedTime"`
		Image       string `json:"image"`
		Summary     string `json:"summary"`
		Links       []struct {
			Href string `json:"href"` // null is also accepted
			Rel  string `json:"rel"`
		} `json:"links"`
		FriendsCount      float64  `json:"friendsCount"`
		FollowersCount    float64  `json:"followersCount"`
		ListedCount       float64  `json:"listedCount"`
		StatusesCount     float64  `json:"statusesCount"`
		TwitterTimeZone   string   `json:"twitterTimeZone"`
		Verified          bool     `json:"verified"`
		UtcOffset         string   `json:"utcOffset"`
		PreferredUsername string   `json:"preferredUsername"`
		Languages         []string `json:"languages"`
		FavoritesCount    float64  `json:"favoritesCount"`
		Location          struct {
			ObjectType  string `json:"objectType"`
			DisplayName string `json:"displayName"`
		} `json:"location"`
	} `json:"actor"`
	Verb      string `json:"verb"`
	Generator struct {
		DisplayName string `json:"displayName"`
		Link        string `json:"link"`
	} `json:"generator"`
	Provider struct {
		ObjectType  string `json:"objectType"`
		DisplayName string `json:"displayName"`
		Link        string `json:"link"`
	} `json:"provider"`
	InReplyTo struct {
		Link string `json:"link"`
	} `json:"inReplyTo"`
	Location TwitterLocation `json:"location"`
	Geo      struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}
	TwitterEntities         TwitterEntities `json:"twitter_entities"`
	TwitterExtendedEntities struct {
		Media []TwitterMedia `json:"media"`
	} `json:"twitter_extended_entities"`
	Link               string            `json:"link"`
	Body               string            `json:"body"`
	ObjectType         string            `json:"objectType"`
	Object             *TwitterPayload   `json:"object"`
	PostedTime         string            `json:"postedTime"`
	LongObject         TwitterLongObject `json:"long_object"`
	DisplayTextRange   []int             `json:"display_text_range"`
	FavoritesCount     float64           `json:"favoritesCount"`
	TwitterFilterLevel string            `json:"twitter_filter_level"`
	TwitterLang        string            `json:"twitter_lang"`
	RetweetCount       float64           `json:"retweetCount"`
	Gnip               struct {
		MatchingRules []struct {
			Tag string `json:"tag"`
			ID  int64  `json:"id"`
		} `json:"matching_rules"`
		Urls []struct {
			URL                    string `json:"url"`
			ExpandedURL            string `json:"expanded_url"`
			ExpandedStatus         int    `json:"expanded_status"`
			ExpandedURLTitle       string `json:"expanded_url_title"`
			ExpandedURLDescription string `json:"expanded_url_description"`
		} `json:"urls"`
		KloutScore int `json:"klout_score"`
		Language   struct {
			Value string `json:"value"`
		} `json:"language"`
		ProfileLocations []struct {
			Address struct {
				Country     string `json:"country"`
				CountryCode string `json:"countryCode"`
				Locality    string `json:"locality"`
				Region      string `json:"region"`
				SubRegion   string `json:"subRegion"`
			} `json:"address"`
			DisplayName string `json:"displayName"`
			Geo         struct {
				Coordinates []float64 `json:"coordinates"`
				Type        string    `json:"type"`
			} `json:"geo"`
			ObjectType string `json:"objectType"`
		} `json:"profileLocations"`
	} `json:"gnip"`
	TwitterQuotedStatus *TwitterPayload `json:"twitter_quoted_status"`
}

// IsReshare checks if the payload represents a repost
func (twitterPayload *TwitterPayload) IsReshare() bool {
	if twitterPayload.Object.ObjectType == "activity" {
		return true
	}
	return false
}

// IsComment checks if the payload represents a comment
func (twitterPayload *TwitterPayload) IsComment() bool {
	return twitterPayload.InReplyTo.Link != ""
}

// Hashtags returns hashtags of the post
func (twitterPayload *TwitterPayload) Hashtags() (hashtags []string) {
	hashtagsArray := twitterPayload.TwitterEntities.Hashtags
	if twitterPayload.LongObject.Body != "" {
		hashtagsArray = twitterPayload.LongObject.TwitterEntities.Hashtags
	}
	for _, v := range hashtagsArray {
		hashtags = append(hashtags, v.Text)
	}
	return
}

// CreationTime returns creation time of the post
func (twitterPayload *TwitterPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RFC3339, twitterPayload.PostedTime, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
