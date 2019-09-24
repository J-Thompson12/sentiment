package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Wordpress is going to be used to get stream data
type Wordpress struct {
	Verb      string    `json:"verb"`
	Published time.Time `json:"published"`
	Content   string    `json:"content"`
	Actor     struct {
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
	} `json:"actor"`
	Target struct {
		ObjectType  string      `json:"objectType"`
		DisplayName string      `json:"displayName"`
		URL         string      `json:"url"`
		BlogID      wordpressID `json:"wpcom:blog_id"`
		PostID      wordpressID `json:"wpcom:post_id"`
	} `json:"target"`
	Object struct {
		PermalinkURL string      `json:"permalinkUrl"`
		URL          string      `json:"url"`
		ObjectType   string      `json:"objectType"`
		Published    time.Time   `json:"published"`
		PostID       wordpressID `json:"wpcom:post_id"`
		CommentID    string      `json:"wpCommentId"`
		Content      string      `json:"content"`
	} `json:"object"`
}

// because who knows what the id will be from post to post
type wordpressID string

func (id *wordpressID) UnmarshalJSON(b []byte) error {
	var i int
	err := json.Unmarshal(b, &i)
	if err != nil {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		*id = wordpressID(str)
		return nil
	}
	*id = wordpressID(fmt.Sprintf("%d", i))
	return nil
}
