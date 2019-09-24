package models

import (
	"time"
)

//WebhosePayload is the data structure based off the format of incoming JSON payloads
type WebhosePayload struct {
	Forum           string `xml:"forum"`
	DiscussionTitle string `xml:"discussion_title"`
	Signature       string `xml:"signature"`
	PostNum         string `xml:"post_num"`
	ExternalLinks   string `xml:"external_links"`
	Post            string `xml:"post"`
	Country         string `xml:"country"`
	SpamScore       string `xml:"spam_score"`
	PostDate        string `xml:"post_date"`
	GmtOffset       string `xml:"gmt_offset"`
	TopicText       string `xml:"topic_text"`
	Username        string `xml:"username"`
	ForumTitle      string `xml:"forum_title"`
	PostTime        string `xml:"post_time"`
	Language        string `xml:"language"`
	PostURL         string `xml:"post_url"`
	Type            string `xml:"type"`
	PostID          string `xml:"post_id"`
	TopicURL        string `xml:"topic_url"`
	MainImage       string `xml:"main_image"`
}

// CreationTime returns creation time of the post
func (webhosePayload *WebhosePayload) CreationTime() (time.Time, error) {
	if (webhosePayload.PostDate != "") && (webhosePayload.PostTime == "") {
		webhosePayload.PostTime = "0000"
	}
	parsedTime, err := time.ParseInLocation("200601021504", webhosePayload.PostDate+webhosePayload.PostTime, time.UTC)

	if err != nil {
		return invalidDate()
	}

	return parsedTime, nil
}
