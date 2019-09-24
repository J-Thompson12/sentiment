package models

import (
	"time"
)

//YoutubePayload is the data structure based off the format of incoming JSON payloads
type YoutubePayload struct {
	Snippet struct {
		Title      string `json:"title"`
		Thumbnails struct {
			Medium struct {
				Width  int    `json:"width"`
				URL    string `json:"url"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				Width  int    `json:"width"`
				URL    string `json:"url"`
				Height int    `json:"height"`
			} `json:"high"`
			Default struct {
				Width  int    `json:"width"`
				URL    string `json:"url"`
				Height int    `json:"height"`
			} `json:"default"`
		} `json:"thumbnails"`
		PublishedAt          string `json:"publishedAt"`
		LiveBroadcastContent string `json:"liveBroadcastContent"`
		Description          string `json:"description"`
		ChannelTitle         string `json:"channelTitle"`
		ChannelID            string `json:"channelId"`
	} `json:"snippet"`
	Kind string `json:"kind"`
	ID   struct {
		VideoID string `json:"videoId"`
		Kind    string `json:"kind"`
	} `json:"id"`
	Etag string `json:"etag"`
}

// CreationTime returns creation time of the post
func (youtubePayload *YoutubePayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RFC3339Nano, youtubePayload.Snippet.PublishedAt, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
