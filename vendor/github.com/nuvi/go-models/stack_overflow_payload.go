package models

import (
	"time"
)

//StackOverflowPayload is the data structure based off the format of incoming JSON payloads
type StackOverflowPayload struct {
	Owner struct {
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Silver int `json:"silver"`
			Gold   int `json:"gold"`
		} `json:"badge_counts"`
		Reputation   int    `json:"reputation"`
		UserID       int    `json:"user_id"`
		UserType     string `json:"user_type"`
		AcceptRate   int    `json:"accept_rate"`
		ProfileImage string `json:"profile_image"`
		DisplayName  string `json:"display_name"`
		Link         string `json:"link"`
	} `json:"owner,omitempty"`
	CommentCount     int    `json:"comment_count"`
	DownVoteCount    int    `json:"down_vote_count"`
	UpVoteCount      int    `json:"up_vote_count"`
	Score            int    `json:"score"`
	LastActivityDate int    `json:"last_activity_date"`
	CreationDate     int    `json:"creation_date"`
	PostType         string `json:"post_type"`
	PostID           int    `json:"post_id"`
	Title            string `json:"title"`
	ShareLink        string `json:"share_link"`
	BodyMarkdown     string `json:"body_markdown"`
	Link             string `json:"link"`
	Body             string `json:"body"`
	LastEditor       struct {
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Silver int `json:"silver"`
			Gold   int `json:"gold"`
		} `json:"badge_counts"`
		Reputation   int    `json:"reputation"`
		UserID       int    `json:"user_id"`
		UserType     string `json:"user_type"`
		AcceptRate   int    `json:"accept_rate"`
		ProfileImage string `json:"profile_image"`
		DisplayName  string `json:"display_name"`
		Link         string `json:"link"`
	} `json:"last_editor,omitempty"`
	LastEditDate int `json:"last_edit_date,omitempty"`
	Comments     []struct {
		Owner struct {
			BadgeCounts struct {
				Bronze int `json:"bronze"`
				Silver int `json:"silver"`
				Gold   int `json:"gold"`
			} `json:"badge_counts"`
			Reputation   int    `json:"reputation"`
			UserID       int    `json:"user_id"`
			UserType     string `json:"user_type"`
			AcceptRate   int    `json:"accept_rate"`
			ProfileImage string `json:"profile_image"`
			DisplayName  string `json:"display_name"`
			Link         string `json:"link"`
		} `json:"owner"`
		CanFlag      bool   `json:"can_flag"`
		Edited       bool   `json:"edited"`
		Score        int    `json:"score"`
		PostType     string `json:"post_type"`
		CreationDate int    `json:"creation_date"`
		PostID       int    `json:"post_id"`
		CommentID    int    `json:"comment_id"`
		BodyMarkdown string `json:"body_markdown"`
		Link         string `json:"link"`
		Body         string `json:"body"`
	} `json:"comments,omitempty"`
}

// CreationTime returns creation time of the post
func (stackOverflow *StackOverflowPayload) CreationTime() (time.Time, error) {
	return time.Unix(int64(stackOverflow.CreationDate), 0).UTC(), nil
}
