package models

import (
	"time"
)

// Location describes a geo location for elastic search
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// CompactSocialActivity is the structure of an activity as it is inserted into Bigfoot
// This should match the mapping of the index
type CompactSocialActivity struct {
	Timestamp                string    `json:"@timestamp"`
	ActivityURL              string    `json:"activity_url"`
	ActivityTitle            string    `json:"activity_title"`
	AggregatedAt             time.Time `json:"aggregated_at"`
	AuthorFollowersCount     int       `json:"author_followers_count"`
	AuthorGender             Gender    `json:"author_gender"`
	AuthorKloutScore         int       `json:"author_klout_score"`
	AuthorPictureURL         string    `json:"author_picture_url"`
	AuthorProfileURL         string    `json:"author_profile_url"`
	AuthorRealName           string    `json:"author_real_name"`
	AuthorUsername           string    `json:"author_username"`
	Bio                      string    `json:"bio"`
	Categories               []string  `json:"categories"`
	CompanyUID               string    `json:"company_uid"`
	CountryCode              string    `json:"country_code"`
	Emojis                   []string  `json:"emojis"`
	Entities                 []string  `json:"entities"`
	EngagementsLastUpdated   time.Time `json:"engagements_last_updated"`
	EngagementsFavorites     int64     `json:"engagements_favorites"`
	EngagementsReplies       int64     `json:"engagements_replies"`
	EngagementsShares        int64     `json:"engagements_shares"`
	EngagementsVideoViews    int64     `json:"engagements_video_views"`
	Hashtags                 []string  `json:"hashtags"`
	Intents                  []string  `json:"intents"`
	IsComment                string    `json:"is_comment"`
	IsReshare                string    `json:"is_reshare"`
	IsVerified               string    `json:"is_verified"`
	Keywords                 []string  `json:"keywords"`
	Language                 string    `json:"language"`
	Location                 Location  `json:"location"`
	MatchedByParent          string    `json:"matched_by_parent"`
	Mentions                 []string  `json:"mentions"`
	MentionsExpanded         string    `json:"mentions_expanded"`
	Network                  string    `json:"network"`
	NormalizedNegativeThemes []string  `json:"normalized_negative_themes"`
	NormalizedPositiveThemes []string  `json:"normalized_positive_themes"`
	NormalizedThemes         []string  `json:"normalized_themes"`
	NormalizedURLs           []string  `json:"normalized_urls"`
	Notes                    string    `json:"notes"`
	OriginalNetwork          string    `json:"original_network"`
	ParentSocialSourceID     string    `json:"parent_social_source_id"`
	ParentAuthorUsername     string    `json:"parent_author_username"`
	PostCreatedAt            time.Time `json:"post_created_at"`
	PostMedia                string    `json:"post_media"`
	RawBodyText              string    `json:"raw_body_text"`
	Region                   string    `json:"region"`
	RegionCode               string    `json:"region_code"`
	SentimentCategory        string    `json:"sentiment_category"`
	SentimentScore           float64   `json:"sentiment_score"`
	SocialMonitorUID         string    `json:"social_monitor_uid"`
	SocialSourceUID          string    `json:"social_source_uid"`
	SubjectivityScore        float64   `json:"subjectivity_score"`
	Tags                     []string  `json:"tags"`
	Version                  string    `json:"@version"`
	UniqueHash               string    `json:"unique_hash"`
}
