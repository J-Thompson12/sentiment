package models

import (
	"errors"
	"time"
)

//ErrInvalidDate is returned when an incoming social activity model fails to
//parse the date string.
var ErrInvalidDate = errors.New("models: cannot parse invalid date in activity payload")

func invalidDate() (time.Time, error) {
	return time.Now(), ErrInvalidDate
}

// SocialActivity contains information parsed from payloads that will be used in Nuvis
type SocialActivity struct {
	// NORMALIZED DATA
	ActivityURL            string `json:"activity_url"`
	ActivityTitle          string `json:"activity_title"`
	AddressRegion          string
	AuthorFavoritesCount   int    `json:"author_favorites_count"`
	AuthorFriendsCount     int    `json:"author_friends_count"`
	AuthorFollowersCount   int    `json:"author_followers_count"`
	AuthorKloutScore       int    `json:"author_klout_score"`
	AuthorPictureURL       string `json:"author_picture_url"`
	AuthorPostsCount       int    `json:"author_posts_count"`
	AuthorProfileURL       string `json:"author_profile_url"`
	AuthorRealName         string `json:"author_real_name"`
	AuthorUsername         string `json:"author_username"`
	AuthorID               string `json:"author_id"`
	Bio                    string `json:"bio"`
	BodyText               string `json:"raw_body_text"`
	CleanedBodyText        string `json:"cleaned_body_text"` //Body text that is clean of all punctuation except for hashtags (#)
	CompanyUID             string `json:"company_uid"`
	Country                string `json:"country"`
	CountryCode            string `json:"country_code"`
	CountryGN              string `json:"country_gn"`
	CountryGNCode          string `json:"country_gn_code"`
	CountryISOCode         string
	EmbeddedUrls           []string `json:"embedded_urls"`
	GeoCoordinates         []float64
	Hashtags               []string       `json:"hashtags"`
	HistoricalSearch       bool           `json:"historical_search"`
	ID                     string         `json:"social_source_uid"`
	IsComment              bool           `json:"is_comment"`
	IsReshare              bool           `json:"is_reshare"`
	IsVerified             bool           `json:"is_verified"`
	Keywords               map[string]int `json:"keywords"`
	Language               string         `json:"language"`
	Latitude               float64        `json:"latitude"`
	Longitude              float64        `json:"longitude"`
	LikeCount              int            `json:"like_count"`
	LocationDisplayID      string         `json:"location_display_id"`
	LocationDisplayName    string         `json:"location_display_name"`
	LocationGeoCoordinates [][][]float64
	MatchedByParent        bool          `json:"matched_by_parent"`
	Mentions               []string      `json:"mentions"`
	MentionsExpanded       []UserMention `json:"mentions_expanded"`
	MetaData               []MetaData    `json:"meta_data"`
	Network                string        `json:"network"`
	OriginalNetwork        string        `json:"original_network"`
	NormalizedUrls         []string      `json:"normalized_urls"`
	Parent                 ParentInfo    `json:"parent"`
	PayloadVersion         string        `json:"payload_version"`
	PostMedia              []Media       `json:"post_media"`
	PostCreatedAt          time.Time     `json:"post_created_at"`
	ProfileGeoLocation     []float64
	Region                 string                `json:"region"`
	RegionGN               string                `json:"region_gn"`
	RegionCode             string                `json:"region_code"`
	RegionGNCode           string                `json:"region_gn_code"`
	RetweetCount           int                   `json:"retweet_count"`
	SocialMonitorUID       string                `json:"social_monitor_uid"`
	SocialSourceUids       []string              `json:"social_monitor_uids"`
	SocialMonitorSources   []SocialMonitorSource `json:"social_monitor_sources"`
	Source                 string                `json:"source"`
	SourceID               string                `json:"source_id"`
	TopicMonitorID         string                `json:"topic_monitor_id"`
	Themes                 []string              `json:"themes"` // used for arabic activities
	UrlsMap                map[string]bool
	// used for nuvis migration
	SentitmentScore float64 `json:"sentiment_score"`
}

// Media contains information and urls pertaining to attached media to be put in PostMedia inside the SocialActivity struct
type Media struct {
	// NORMALIZED DATA
	MediaURL   string `json:"media_url"`
	URL        string `json:"url"`
	DisplayURL string `json:"display_url"` // Should be a photo for displaying in the social activity card
	MediaType  string `json:"media_type"`
	VideoURL   string `json:"video_url"`
}

// MetaData is a struct for fields that don't fit into the SocialActivity struct but are still important
type MetaData struct {
	RedditScore             int    `json:"reddit_score"`              //Net of upvotes and downvotes
	StackOverflowReputation int    `json:"stack_overflow_reputation"` //Reputation of the original poster
	PinterestBoardName      string `json:"pinterest_board_name"`      //Name of the authors pinterest board
}

// ParentInfo is a struct that contains all parent information from social activity that gets set when given reshare payloads
type ParentInfo struct {
	ParentAuthor           string    `json:"parent_author"`
	ParentAuthorProfileURL string    `json:"parent_author_profile_url"`
	ParentAuthorReachCount int       `json:"parent_author_reach"`
	ParentBodyText         string    `json:"parent_body_text"`
	ParentCreatedAt        time.Time `json:"parent_created_at"`
	ParentID               string    `json:"parent_social_source_id"`
}

// SocialMonitorSource splitting the source into monitor and company id for better topic matching
type SocialMonitorSource struct {
	CompanyID         string         `json:"company_uid"`
	MonitorID         string         `json:"monitor_uid"`
	Keywords          map[string]int `json:"keywords"`
	SentimentCategory string         `json:"sentiment_category"`
	MatchedByParent   bool           `json:"matched_by_parent"`
}

// UserMention represents a user mentioned in a post
type UserMention struct {
	DisplayName string `json:"display_name"`
	UserID      string `json:"user_id"`
	Network     string `json:"network"`
	ProfileURL  string `json:"profile_url"`
}

// HasGeo checks to see if lat and lon exist and returns true if both exist
func (activity *SocialActivity) HasGeo() bool {
	if activity.Latitude == 0.0 &&
		activity.Longitude == 0.0 {
		return false
	}
	return true
}
