package models

import (
	"time"
)

//TwitterDirectPayload is the data structure based off the format of incoming JSON payloads
type TwitterDirectPayload struct {
	CreatedAt        string          `json:"created_at"`
	ID               int64           `json:"id"`
	IDStr            string          `json:"id_str"`
	Text             string          `json:"text"`
	Truncated        bool            `json:"truncated"`
	Entities         TwitterEntities `json:"entities"`
	ExtendedEntities struct {
		Media []struct {
			ID            int64  `json:"id"`
			IDStr         string `json:"id_str"`
			Indices       []int  `json:"indices"`
			MediaURL      string `json:"media_url"`
			MediaURLHTTPS string `json:"media_url_https"`
			URL           string `json:"url"`
			DisplayURL    string `json:"display_url"`
			ExpandedURL   string `json:"expanded_url"`
			Type          string `json:"type"`
			Sizes         struct {
				Medium Size `json:"medium"`
				Large  Size `json:"large"`
				Small  Size `json:"small"`
				Thumb  Size `json:"thumb"`
			} `json:"sizes"`
			VideoInfo struct {
				AspectRatio    []int `json:"aspect_ratio"`
				DurationMillis int   `json:"duration_millis"`
				Variants       []struct {
					ContentType string `json:"content_type"`
					URL         string `json:"url"`
					Bitrate     int    `json:"bitrate,omitempty"`
				} `json:"variants"`
			} `json:"video_info"`
			AdditionalMediaInfo struct {
				Monetizable bool `json:"monetizable"`
			} `json:"additional_media_info"`
		} `json:"media"`
	} `json:"extended_entities"`
	Source               string `json:"source"`
	InReplyToStatusID    int64  `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int    `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  string `json:"in_reply_to_screen_name"`
	User                 struct {
		ID          int64       `json:"id"`
		IDStr       string      `json:"id_str"`
		Name        string      `json:"name"`
		ScreenName  string      `json:"screen_name"`
		Location    string      `json:"location"`
		Description string      `json:"description"`
		URL         interface{} `json:"url"`
		Entities    struct {
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		Protected                      bool   `json:"protected"`
		FollowersCount                 int    `json:"followers_count"`
		FriendsCount                   int    `json:"friends_count"`
		ListedCount                    int    `json:"listed_count"`
		CreatedAt                      string `json:"created_at"`
		FavouritesCount                int    `json:"favourites_count"`
		UtcOffset                      int    `json:"utc_offset"`
		TimeZone                       string `json:"time_zone"`
		GeoEnabled                     bool   `json:"geo_enabled"`
		Verified                       bool   `json:"verified"`
		StatusesCount                  int    `json:"statuses_count"`
		Lang                           string `json:"lang"`
		ContributorsEnabled            bool   `json:"contributors_enabled"`
		IsTranslator                   bool   `json:"is_translator"`
		IsTranslationEnabled           bool   `json:"is_translation_enabled"`
		ProfileBackgroundColor         string `json:"profile_background_color"`
		ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool   `json:"profile_background_tile"`
		ProfileImageURL                string `json:"profile_image_url"`
		ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
		ProfileLinkColor               string `json:"profile_link_color"`
		ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
		HasExtendedProfile             bool   `json:"has_extended_profile"`
		DefaultProfile                 bool   `json:"default_profile"`
		DefaultProfileImage            bool   `json:"default_profile_image"`
		Following                      bool   `json:"following"`
		FollowRequestSent              bool   `json:"follow_request_sent"`
		Notifications                  bool   `json:"notifications"`
	} `json:"user"`
	Geo         interface{} `json:"geo"`
	Coordinates struct {
		Coordinates [][][]float64 `json:"coordinates"`
		Type        string        `json:"type"`
	} `json:"coordinates"`
	Place struct {
		ID              string `json:"id"`
		URL             string `json:"url"`
		PlaceType       string `json:"place_type"`
		Name            string `json:"name"`
		FullName        string `json:"full_name"`
		CountryCode     string `json:"country_code"`
		Country         string `json:"country"`
		ContainedWithin []struct {
			Attributes struct {
			} `json:"attributes"`
		} `json:"contained_within"`
		BoundingBox struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"bounding_box"`
		Attributes struct {
		} `json:"attributes"`
	} `json:"place"`
	Contributors                interface{} `json:"contributors"`
	IsQuoteStatus               bool        `json:"is_quote_status"`
	RetweetCount                int         `json:"retweet_count"`
	FavoriteCount               int         `json:"favorite_count"`
	Favorited                   bool        `json:"favorited"`
	Retweeted                   bool        `json:"retweeted"`
	PossiblySensitive           bool        `json:"possibly_sensitive"`
	PossiblySensitiveAppealable bool        `json:"possibly_sensitive_appealable"`
	Lang                        string      `json:"lang"`
}

// Hashtags returns hashtags of the post
func (twitterDirectPayload *TwitterDirectPayload) Hashtags() (hashtags []string) {
	for _, v := range twitterDirectPayload.Entities.Hashtags {
		hashtags = append(hashtags, v.Text)
	}
	return
}

// CreationTime returns creation time of the post
func (twitterDirectPayload *TwitterDirectPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("Mon Jan 02 15:04:05 +0000 2006", twitterDirectPayload.CreatedAt, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
