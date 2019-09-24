package models

import (
	"math"
	"time"
)

//RedditPayload is the data structure based off the format of incoming JSON payloads
type RedditPayload struct {
	Kind string `json:"kind"`
	Data struct {
		ApprovedBy          string      `json:"approved_by"`
		Archived            bool        `json:"archived"`
		Author              string      `json:"author"`
		AuthorFlairCSSClass string      `json:"author_flair_css_class"`
		AuthorFlairText     string      `json:"author_flair_text"`
		BannedBy            string      `json:"banned_by"`
		Clicked             bool        `json:"clicked"`
		ContestMode         bool        `json:"contest_mode"`
		Created             float64     `json:"created"`
		CreatedUtc          float64     `json:"created_utc"`
		Distinguished       string      `json:"distinguished"`
		Domain              string      `json:"domain"`
		Downs               int         `json:"downs"`
		Edited              interface{} `json:"edited"`
		Gilded              int         `json:"gilded"`
		Hidden              bool        `json:"hidden"`
		HideScore           bool        `json:"hide_score"`
		ID                  string      `json:"id"`
		IsSelf              bool        `json:"is_self"`
		Likes               bool        `json:"likes"`
		LinkFlairCSSClass   string      `json:"link_flair_css_class"`
		LinkFlairText       string      `json:"link_flair_text"`
		Locked              bool        `json:"locked"`
		Media               interface{} `json:"media"`
		MediaEmbed          struct {
			Content   string `json:"content"`
			Width     int    `json:"width"`
			Scrolling bool   `json:"scrolling"`
			Height    int    `json:"height"`
		} `json:"media_embed"`
		ModReports  []interface{} `json:"mod_reports"`
		Name        string        `json:"name"`
		NumComments int           `json:"num_comments"`
		NumReports  int           `json:"num_reports"`
		Over18      bool          `json:"over_18"`
		ParentID    string        `json:"parent_id"`
		Permalink   string        `json:"permalink"`
		PostHint    string        `json:"post_hint"`
		Preview     struct {
			Images []struct {
				ID          string `json:"id"`
				Resolutions []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"resolutions"`
				Source struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"source"`
				Variants struct {
				} `json:"variants"`
			} `json:"images"`
		} `json:"preview"`
		Quarantine       bool        `json:"quarantine"`
		RemovalReason    interface{} `json:"removal_reason"`
		ReportReasons    interface{} `json:"report_reasons"`
		Saved            bool        `json:"saved"`
		Score            int         `json:"score"`
		SecureMedia      interface{} `json:"secure_media"`
		SecureMediaEmbed struct {
			Content   string `json:"content"`
			Width     int    `json:"width"`
			Scrolling bool   `json:"scrolling"`
			Height    int    `json:"height"`
		} `json:"secure_media_embed"`
		Selftext      string        `json:"selftext"`
		SelftextHTML  string        `json:"selftext_html"`
		Stickied      bool          `json:"stickied"`
		Subreddit     string        `json:"subreddit"`
		SubredditID   string        `json:"subreddit_id"`
		SuggestedSort string        `json:"suggested_sort"`
		Thumbnail     string        `json:"thumbnail"`
		Title         string        `json:"title"`
		Ups           int           `json:"ups"`
		URL           string        `json:"url"`
		UserReports   []interface{} `json:"user_reports"`
		Visited       bool          `json:"visited"`
	} `json:"data"`
}

// CreationTime returns creation time of the post
func (redditPayload *RedditPayload) CreationTime() (time.Time, error) {
	parsedTime := time.Unix(int64(redditPayload.Data.CreatedUtc),
		int64((redditPayload.Data.CreatedUtc-math.Floor(redditPayload.Data.CreatedUtc))*1000000000)).UTC()

	return parsedTime, nil
}
