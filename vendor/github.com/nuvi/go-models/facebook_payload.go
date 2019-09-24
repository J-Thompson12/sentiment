package models

import (
	"fmt"
	"time"
)

//FacebookProfile is based off of the "profile" data type from Facebooks Graph API
type FacebookProfile interface{}

//FacebookAttachment is based off of the "attachment" data type from Facebooks Graph API
type FacebookAttachment struct {
	Description string `json:"description"`
	Media       struct {
		Image struct {
			Height int    `json:"height"`
			Src    string `json:"src"`
			Width  int    `json:"width"`
		} `json:"image"`
	} `json:"media"`
	Target struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"target"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

// FacebookWithTags is only used by go-social-activity-parser
type FacebookWithTags struct {
		Data []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}


//FacebookMentionTag represents a single mention tag from facebook
type FacebookMentionTag struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

// FacebookReaction represents all the info attached to a facebook reaction, including the reaction itself.
type FacebookReaction struct {
	CanPost bool   `json:"can_post"`
	ID      string `json:"id"`
	Link    string `json:"link"`
	Name    string `json:"name"`
	Pic     string `json:"pic"`
	PicCrop struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Left   int    `json:"left"`
		Top    int    `json:"top"`
		Right  int    `json:"right"`
		Bottom int    `json:"bottom"`
		URL    string `json:"url"`
	} `json:"pic_crop"`
	PicLarge    string `json:"pic_large"`
	PicSquare   string `json:"pic_square"`
	ProfileType string `json:"profile_type"`
	PicSmall    string `json:"pic_small"`
	Type        string `json:"type"`
}

//FacebookPayload is the data structure based off the format of incoming JSON payloads
type FacebookPayload struct {
	QueryMetadata struct {
		Query string
	}
	ID           string `json:"id"`
	AdminCreator []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"admin_creator"`
	Application struct {
		Category string `json:"category"`
		Link     string `json:"link"`
		Name     string `json:"name"`
		ID       string `json:"id"`
	} `json:"application,omitempty"`
	CallToAction struct {
		Context interface{} `json:"context"`
	} `json:"call_to_action"`
	Caption       string `json:"caption"`
	CreatedTime   string `json:"created_time"`
	Description   string `json:"description"`
	FeedTargeting struct {
		AgeMax               int         `json:"age_max"`
		AgeMin               int         `json:"age_min"`
		Cities               []int       `json:"cities"`
		CollegeYears         []int       `json:"college_years"`
		Countries            []string    `json:"countries"`
		EducationStatuses    []int       `json:"education_statuses"`
		Genders              []int       `json:"genders"`
		InterestedIn         []int       `json:"interested_in"`
		Locales              []int       `json:"locales"`
		Regions              interface{} `json:"regions"`
		RelationshipStatuses []int       `json:"relationship_statuses"`
	} `json:"feed_targeting"`
	From struct {
		Username string `json:"username"`
		FanCount int64  `json:"fan_count"`
		Name     string `json:"name"`
		ID       string `json:"id"`
	} `json:"from"`
	Icon                 string               `json:"icon"`
	InstagramEligibility string               `json:"instagram_eligibility"`
	IsHidden             bool                 `json:"is_hidden"`
	IsInstagramEligible  bool                 `json:"is_instagram_eligible"`
	IsPublished          bool                 `json:"is_published"`
	Link                 string               `json:"link"`
	Message              string               `json:"message"`
	MessageTags          []FacebookMentionTag `json:"message_tags"`
	Name                 string               `json:"name"`
	ObjectID             string               `json:"object_id"`
	ParentID             string               `json:"parent_id"`
	PermalinkURL         string               `json:"permalink_url"`
	Picture              string               `json:"picture"`
	Place                struct {
		ID       string `json:"id"`
		Location struct {
			City      string  `json:"city"`
			Country   string  `json:"country"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			State     string  `json:"state"`
			Street    string  `json:"street"`
			Zip       string  `json:"zip"`
		} `json:"location"`
		Name          string  `json:"name"`
		OverallRating float64 `json:"overall_rating"`
	} `json:"place"`
	Privacy struct {
		Description string `json:"description"`
		Value       string `json:"value"`
		Friends     string `json:"friends"`
		Allow       string `json:"allow"`
		Deny        string `json:"deny"`
	} `json:"privacy"`
	Properties []struct {
		Name string `json:"name"`
		Text string `json:"text"`
		HREF string `json:"href"`
	} `json:"properties"`
	Shares struct {
		Count int `json:"count"`
	} `json:"shares"`
	Source     string `json:"source"`
	StatusType string `json:"status_type"`
	Story      string `json:"story"`
	StoryTags  []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Length int    `json:"length"`
	} `json:"story_tags"`
	Targeting struct {
		Countries []string `json:"countries"`
		Locales   []int    `json:"locales"`
		Regions   []uint32 `json:"regions"`
		Cities    []uint32 `json:"cities"`
	} `json:"targeting"`
	To          []FacebookProfile `json:"to"`
	Type        string            `json:"type"`
	UpdatedTime string            `json:"updated_time"`
	WithTags    struct {
		Data []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	} `json:"with_tags"`
	Actions []struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"actions"`
	FullPicture         string `json:"full_picture"`
	InstreamEligibility string `json:"instream_eligibility"`
	IsExpired           bool   `json:"is_expired"`
	PromotionStatus     string `json:"promotion_status"`
	Subscribed          bool   `json:"subscribed"`
	TimelineVisibility  string `json:"timeline_visibility"`
	IsSpherical         bool   `json:"is_spherical"`
	Reactions           struct {
		Data   []FacebookReaction `json:"data"`
		Paging FacebookPaging     `json:"paging"`
	} `json:"reactions"`
	Sharedposts struct {
		Data []struct {
			Actions []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"actions"`
			Type        string `json:"type"`
			Description string `json:"description"`
			CreatedTime string `json:"created_time"`
			ParentID    string `json:"parent_id"`
			ObjectID    string `json:"object_id"`
			ID          string `json:"id"`
			Story       string `json:"story"`
		} `json:"data"`
		Paging FacebookPaging `json:"paging"`
	} `json:"sharedposts"`
	Likes struct {
		Data    []FacebookReaction `json:"data"`
		Paging  FacebookPaging     `json:"paging"`
		Summary struct {
			TotalCount int  `json:"total_count"`
			CanLike    bool `json:"can_like"`
			HasLiked   bool `json:"has_liked"`
		} `json:"summary"`
	} `json:"likes"`
	Comments struct {
		Data   []FacebookComment `json:"data"`
		Paging FacebookPaging    `json:"paging"`
	} `json:"comments"`
	Attachments struct {
		Data []struct {
			FacebookAttachment
			Subattachments struct {
				Data []FacebookAttachment `json:"data"`
			} `json:"subattachments"`
		} `json:"data"`
	} `json:"attachments"`
}

// FacebookPaging is necessary for navigating paged documents received from facebook
type FacebookPaging struct {
	Cursors struct {
		Before string `json:"before"`
		After  string `json:"after"`
	} `json:"cursors"`
	Next string `json:"next"`
}

type FacebookClientCode struct {
	Code string `json:"code"`
}

// IsReshare checks if the payload represents a repost
func (facebookPayload *FacebookPayload) IsReshare() bool {
	return false
}

// RealPermalinkURL returns a permanent link to the post (even if Facebook didn't provide it)
func (facebookPayload *FacebookPayload) RealPermalinkURL() string {
	if facebookPayload.PermalinkURL != "" {
		return facebookPayload.PermalinkURL
	}
	if facebookPayload.ID != "" {
		return fmt.Sprintf("https://www.facebook.com/%s", facebookPayload.ID)
	}
	return ""
}

// CreationTime returns creation time of the post
func (facebookPayload *FacebookPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02T15:04:05-0700", facebookPayload.CreatedTime, time.UTC)

	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

// LocationDisplayName returns a name of the post location (if present)
func (facebookPayload *FacebookPayload) LocationDisplayName() string {
	return facebookPayload.Place.Name
}
