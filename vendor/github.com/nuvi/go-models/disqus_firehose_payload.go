package models

import "time"

//DisqusFirehosePayload is the data structure based off the format of incoming JSON payloads
type DisqusFirehosePayload struct {
	ID         string `json:"id"`
	Verb       string `json:"verb"`
	PostedTime string `json:"postedTime"`
	Body       string `json:"body"`
	Actor      struct {
		ObjectType        string `json:"objectType"`
		ID                string `json:"id"`
		DisqusID          string `json:"disqusId"`
		PreferredUsername string `json:"preferredUsername"`
	} `json:"actor"`
	Target struct {
		ObjectType string `json:"objectType"`
		ID         string `json:"id"`
		Link       string `json:"link"`
		Website    struct {
			ID       string `json:"id"`
			DisqusID string `json:"disqusId"`
		} `json:"website"`
		DisqusID   string `json:"disqusId"`
		PostedTime string `json:"postedTime"`
	} `json:"target"`
	Object struct {
		ObjectType string `json:"objectType"`
		ID         string `json:"id"`
		Link       string `json:"link"`
		DisqusID   string `json:"disqusId"`
	} `json:"object"`
	DisqusType             string `json:"disqusType"`
	DisqusTypePrev         string `json:"disqusTypePrev"`
	DisqusMessageTimestamp string `json:"disqusMessageTimestamp"`
	InReplyTo              struct {
		ObjectType string `json:"objectType"`
		ID         string `json:"id"`
		Author     struct {
			ID       string `json:"id"`
			DisqusID string `json:"disqusId"`
		} `json:"author"`
		DisqusID string `json:"disqusId"`
	} `json:"inReplyTo,omitempty"`
}

// IsReshare checks if the payload represents a repost
func (disqusFirehose *DisqusFirehosePayload) IsReshare() bool {
	return false
}

// CreationTime returns creation time of the post
func (disqusFirehose *DisqusFirehosePayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RFC3339, disqusFirehose.PostedTime, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

// AuthorProfileURL returns the post author's profile URL
func (disqusFirehose *DisqusFirehosePayload) AuthorProfileURL() (URL string) {
	if disqusFirehose.Actor.PreferredUsername != "" {
		URL = "https://disqus.com/" + disqusFirehose.Actor.PreferredUsername
	}
	return
}
