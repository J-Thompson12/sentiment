package models

import (
	"time"
)

//WordpressFirehosePayload is the data structure based off the format of incoming JSON payloads
type WordpressFirehosePayload struct {
	Verb        string `json:"verb"`
	ID          string `json:"id"`
	PostedTime  string `json:"postedTime"`
	DisplayName string `json:"displayName"`
	Body        string `json:"body"`
	Provider    struct {
		ObjectType  string `json:"objectType"`
		DisplayName string `json:"displayName"`
		Link        string `json:"link"`
	} `json:"provider"`
	Actor struct {
		ObjectType  string `json:"objectType"`
		DisplayName string `json:"displayName"`
		ID          string `json:"id"`
		WpEmailMd5  string `json:"wpEmailMd5"`
		Link        string `json:"link"`
	} `json:"actor"`
	Target struct {
		ObjectType     string `json:"objectType"`
		DisplayName    string `json:"displayName"`
		Link           string `json:"link"`
		WpBlogID       int    `json:"wpBlogId"`
		WpPostID       int    `json:"wpPostId"`
		WpCommentCount int    `json:"wpCommentCount"`
		Feed           string `json:"feed"`
		Summary        string `json:"summary"`
		Author         struct {
			ObjectType  string `json:"objectType"`
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			WpEmailMd5  string `json:"wpEmailMd5"`
			Link        string `json:"link"`
		} `json:"author"`
		ID string `json:"id"`
	} `json:"target"`
	Object struct {
		ObjectType  string `json:"objectType"`
		DisplayName string `json:"displayName"`
		Link        string `json:"link"`
		WpPostID    int    `json:"wpPostId"`
		Summary     string `json:"summary"`
		Content     string `json:"content"`
		WpBlogID    int    `json:"wpBlogId"`
		ID          string `json:"id"`
		Tags        []struct {
			ObjectType  string `json:"objectType"`
			DisplayName string `json:"displayName"`
			Link        string `json:"link"`
		} `json:"tags"`
		InReplyTo struct {
			Link string `json:"link"`
		} `json:"inReplyTo"`
	} `json:"object"`
}

// CreationTime returns creation time of the post
func (wordpressFirehosePayload *WordpressFirehosePayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RFC3339, wordpressFirehosePayload.PostedTime, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
