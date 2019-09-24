package models

import "time"

// TwitterUHFPost represents a Twitter post from the Twitter API
type TwitterUHFPost struct {
	Contributors []int64 `json:"contributors"`
	Coordinates  struct {
		Coordinates [2]float64 `json:"coordinates"` // Coordinate always has to have exactly 2 values
		Type        string     `json:"type"`
	} `json:"coordinates"`
	CreatedAt            string             `json:"created_at"`
	Entities             TwitterUHFEntities `json:"entities"`
	ExtendedEntities     TwitterUHFEntities `json:"extended_entities"`
	FavoriteCount        int                `json:"favorite_count"`
	Favorited            bool               `json:"favorited"`
	FilterLevel          string             `json:"filter_level"`
	HasExtendedProfile   bool               `json:"has_extended_profile"`
	ID                   int64              `json:"id"`
	IDStr                string             `json:"id_str"`
	InReplyToScreenName  string             `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64              `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string             `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64              `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string             `json:"in_reply_to_user_id_str"`
	IsTranslationEnabled bool               `json:"is_translation_enabled"`
	Lang                 string             `json:"lang"`
	Place                struct {
		Attributes  map[string]string `json:"attributes"`
		BoundingBox struct {
			Coordinates [][][]float64 `json:"coordinates"`
			Type        string        `json:"type"`
		} `json:"bounding_box"`
		ContainedWithin []struct {
			Attributes  map[string]string `json:"attributes"`
			BoundingBox struct {
				Coordinates [][][]float64 `json:"coordinates"`
				Type        string        `json:"type"`
			} `json:"bounding_box"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			FullName    string `json:"full_name"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			PlaceType   string `json:"place_type"`
			URL         string `json:"url"`
		} `json:"contained_within"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		FullName    string `json:"full_name"`
		Geometry    struct {
			Coordinates [][][]float64 `json:"coordinates"`
			Type        string        `json:"type"`
		} `json:"geometry"`
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		PlaceType string   `json:"place_type"`
		Polylines []string `json:"polylines"`
		URL       string   `json:"url"`
	} `json:"place"`
	QuotedStatusID              int64                  `json:"quoted_status_id"`
	QuotedStatusIDStr           string                 `json:"quoted_status_id_str"`
	QuotedStatus                *TwitterUHFPost        `json:"quoted_status"`
	PossiblySensitive           bool                   `json:"possibly_sensitive"`
	PossiblySensitiveAppealable bool                   `json:"possibly_sensitive_appealable"`
	RetweetCount                int                    `json:"retweet_count"`
	Retweeted                   bool                   `json:"retweeted"`
	RetweetedStatus             *TwitterUHFPost        `json:"retweeted_status"`
	Source                      string                 `json:"source"`
	Scopes                      map[string]interface{} `json:"scopes"`
	Text                        string                 `json:"text"`
	Truncated                   bool                   `json:"truncated"`
	User                        TwitterUHFUser         `json:"user"`
	WithheldCopyright           bool                   `json:"withheld_copyright"`
	WithheldInCountries         []string               `json:"withheld_in_countries"`
	WithheldScope               string                 `json:"withheld_scope"`
}

//TwitterUHFEntities represents Twitter Entities received from the Twitter API
type TwitterUHFEntities struct {
	Hashtags []struct {
		Indices []int  `json:"indices"`
		Text    string `json:"text"`
	} `json:"hashtags"`
	URLs []struct {
		Indices     []int  `json:"indices"`
		URL         string `json:"url"`
		DisplayURL  string `json:"display_url"`
		ExpandedURL string `json:"expanded_url"`
	} `json:"urls"`
	URL struct {
		URLs []struct {
			Indices     []int  `json:"indices"`
			URL         string `json:"url"`
			DisplayURL  string `json:"display_url"`
			ExpandedURL string `json:"expanded_url"`
		} `json:"urls"`
	} `json:"url"`
	UserMentions []struct {
		Name       string `json:"name"`
		Indices    []int  `json:"indices"`
		ScreenName string `json:"screen_name"`
		ID         int64  `json:"id"`
		IDStr      string `json:"id_str"`
	} `json:"user_mentions"`
	Media []struct {
		ID            int64  `json:"id"`
		IDStr         string `json:"id_str"`
		MediaURL      string `json:"media_url"`
		MediaURLHTTPS string `json:"media_url_https"`
		URL           string `json:"url"`
		DisplayURL    string `json:"display_url"`
		ExpandedURL   string `json:"expanded_url"`
		Sizes         struct {
			Medium Size `json:"medium"`
			Thumb  Size `json:"thumb"`
			Small  Size `json:"small"`
			Large  Size `json:"large"`
		} `json:"sizes"`
		SourceStatusID    int64  `json:"source_status_id"`
		SourceStatusIDStr string `json:"source_status_id_str"`
		Type              string `json:"type"`
		Indices           []int  `json:"indices"`
		VideoInfo         struct {
			AspectRatio    []int `json:"aspect_ratio"`
			DurationMillis int64 `json:"duration_millis"`
			Variants       []struct {
				Bitrate     int    `json:"bitrate"`
				ContentType string `json:"content_type"`
				URL         string `json:"url"`
			} `json:"variants"`
		} `json:"video_info"`
	} `json:"media"`
}

// TwitterUHFRelations contains an array of IDs received from the Twitter API
type TwitterUHFRelations struct {
	IDs []int64 `json:"ids"`
}

// TwitterUHFUser represents a Twitter user received from the Twitter API
type TwitterUHFUser struct {
	ContributorsEnabled            bool               `json:"contributors_enabled"`
	CreatedAt                      string             `json:"created_at"`
	DefaultProfile                 bool               `json:"default_profile"`
	DefaultProfileImage            bool               `json:"default_profile_image"`
	Description                    string             `json:"description"`
	Entities                       TwitterUHFEntities `json:"entities"`
	FavouritesCount                int                `json:"favourites_count"`
	FollowRequestSent              bool               `json:"follow_request_sent"`
	FollowersCount                 int                `json:"followers_count"`
	Following                      bool               `json:"following"`
	FriendsCount                   int                `json:"friends_count"`
	GeoEnabled                     bool               `json:"geo_enabled"`
	HasExtendedProfile             bool               `json:"has_extended_profile"`
	ID                             int64              `json:"id"`
	IDStr                          string             `json:"id_str"`
	IsTranslator                   bool               `json:"is_translator"`
	IsTranslationEnabled           bool               `json:"is_translation_enabled"`
	Lang                           string             `json:"lang"` // BCP-47 code of user defined language
	ListedCount                    int64              `json:"listed_count"`
	Location                       string             `json:"location"` // User defined location
	Name                           string             `json:"name"`
	Notifications                  bool               `json:"notifications"`
	ProfileBackgroundColor         string             `json:"profile_background_color"`
	ProfileBackgroundImageURL      string             `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string             `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool               `json:"profile_background_tile"`
	ProfileBannerURL               string             `json:"profile_banner_url"`
	ProfileImageURL                string             `json:"profile_image_url"`
	ProfileImageURLHTTPS           string             `json:"profile_image_url_https"`
	ProfileLinkColor               string             `json:"profile_link_color"`
	ProfileSidebarBorderColor      string             `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string             `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string             `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool               `json:"profile_use_background_image"`
	Protected                      bool               `json:"protected"`
	ScreenName                     string             `json:"screen_name"`
	Status                         *TwitterUHFPost    `json:"status"` // Only included if the user is a friend
	StatusesCount                  int64              `json:"statuses_count"`
	TimeZone                       string             `json:"time_zone"`
	URL                            string             `json:"url"`
	UTCOffset                      int                `json:"utc_offset"`
	Verified                       bool               `json:"verified"`
	WithheldInCountries            []string           `json:"withheld_in_countries"`
	WithheldScope                  string             `json:"withheld_scope"`
}

// CreationTime returns creation time of the post
func (twitterPost *TwitterUHFPost) CreationTime() (time.Time, error) {
	createdAt, err := time.Parse(time.RubyDate, twitterPost.CreatedAt)
	if err != nil {
		return invalidDate()
	}
	return createdAt, nil
}
