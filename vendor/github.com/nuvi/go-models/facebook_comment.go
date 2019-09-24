package models

import (
	"fmt"
	"time"
)

// FacebookComment represents a comment from facebook
type FacebookComment struct {
	Attachment   FacebookAttachment `json:"attachment,omitempty"`
	CanComment   bool               `json:"can_comment"`
	CanLike      bool               `json:"can_like"`
	CanRemove    bool               `json:"can_remove"`
	CommentCount int                `json:"comment_count"`
	CreatedTime  string             `json:"created_time"`
	From         struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"from"`
	ID        string `json:"id"`
	IsHidden  bool   `json:"is_hidden"`
	IsPrivate bool   `json:"is_private"`
	LikeCount int    `json:"like_count"`
	Likes     struct {
		Data   []FacebookReaction `json:"data"`
		Paging FacebookPaging     `json:"paging"`
	} `json:"likes"`
	Message      string               `json:"message"`
	MessageTags  []FacebookMentionTag `json:"message_tags,omitempty"`
	PermalinkURL string               `json:"permalink_url"`
	Reactions    struct {
		Data   []FacebookReaction `json:"data"`
		Paging FacebookPaging     `json:"paging"`
	} `json:"reactions"`
	Subcomments struct {
		Data   []FacebookComment `json:"data"`
		Paging FacebookPaging    `json:"paging"`
	} `json:"comments,omitempty"`
	UserLikes bool `json:"user_likes"`
}

// IsReshare checks if the payload represents a repost
func (facebookPayload *FacebookComment) IsReshare() bool {
	return false
}

// RealPermalinkURL returns a permanent link to the post (even if Facebook didn't provide it)
func (facebookPayload *FacebookComment) RealPermalinkURL() string {
	if facebookPayload.PermalinkURL != "" {
		return facebookPayload.PermalinkURL
	}
	if facebookPayload.ID != "" {
		return fmt.Sprintf("https://www.facebook.com/%s", facebookPayload.ID)
	}
	return ""
}

// CreationTime returns creation time of the post
func (facebookPayload *FacebookComment) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02T15:04:05-0700", facebookPayload.CreatedTime, time.UTC)

	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
