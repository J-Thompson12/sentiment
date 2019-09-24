package models

import (
	"time"
)

//GooglePlusPayload is the data structure based off the format of incoming JSON payloads
type GooglePlusPayload struct {
	Verb      string `json:"verb"`
	URL       string `json:"url"`
	Updated   string `json:"updated"`
	Title     string `json:"title"`
	Published string `json:"published"`
	Provider  struct {
		Title string `json:"title"`
	} `json:"provider"`
	Object struct {
		URL       string `json:"url"`
		Resharers struct {
			TotalItems int    `json:"totalItems"`
			SelfLink   string `json:"selfLink"`
		} `json:"resharers"`
		Replies struct {
			TotalItems int    `json:"totalItems"`
			SelfLink   string `json:"selfLink"`
		} `json:"replies"`
		Plusoners struct {
			TotalItems int    `json:"totalItems"`
			SelfLink   string `json:"selfLink"`
		} `json:"plusoners"`
		ObjectType      string       `json:"objectType"`
		ID              string       `json:"id"`
		Content         string       `json:"content"`
		OriginalContent string       `json:"originalContent"`
		Attachments     []Attachment `json:"attachments"`
		Actor           struct {
			Verification struct {
				AdHocVerified string `json:"adHocVerified"`
			} `json:"verification"`
			URL   string `json:"url"`
			Image struct {
				URL string `json:"url"`
			} `json:"image"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"actor"`
	} `json:"object"`
	Kind       string `json:"kind"`
	ID         string `json:"id"`
	Etag       string `json:"etag"`
	Annotation string `json:"annotation"`
	Actor      struct {
		Verification struct {
			AdHocVerified string `json:"adHocVerified"`
		} `json:"verification"`
		URL   string `json:"url"`
		Image struct {
			URL string `json:"url"`
		} `json:"image"`
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Name        struct {
			GivenName  string `json:"givenName"`
			FamilyName string `json:"familyName"`
		} `json:"name"`
	} `json:"actor"`
	Access struct {
		Kind  string `json:"kind"`
		Items []struct {
			Type string `json:"type"`
		} `json:"items"`
		Description string `json:"description"`
	} `json:"access"`
	Location struct {
		Position struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
		DisplayName string `json:"displayName"`
	} `json:"location"`
}

// Attachment represents an attachment in a post, comment, etc.
type Attachment struct {
	URL   string `json:"url"`
	Embed struct {
		URL  string `json:"url"`
		Type string `json:"type"`
	}
	Thumbnails []struct {
		URL   string `json:"url"`
		Image struct {
			Width  int    `json:"width"`
			URL    string `json:"url"`
			Type   string `json:"type"`
			Height int    `json:"height"`
		} `json:"image"`
		Description string `json:"description"`
	} `json:"thumbnails"`
	Image struct {
		Width  int    `json:"width"`
		URL    string `json:"url"`
		Type   string `json:"type"`
		Height int    `json:"height"`
	} `json:"image"`
	ObjectType  string `json:"objectType"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

// IsReshare checks if the payload represents a repost
func (googleplusPayload *GooglePlusPayload) IsReshare() bool {
	if googleplusPayload.Verb == "share" {
		return true
	}
	return false
}

// FullRealName returns the author's full real name
func (googleplusPayload *GooglePlusPayload) FullRealName() (fullName string) {
	if len(googleplusPayload.Actor.Name.GivenName) != 0 && len(googleplusPayload.Actor.Name.FamilyName) != 0 {
		fullName = googleplusPayload.Actor.Name.GivenName + " " + googleplusPayload.Actor.Name.FamilyName
	}
	return
}

// CreationTime returns creation time of the post
func (googleplusPayload *GooglePlusPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RFC3339, googleplusPayload.Published, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, err
}
