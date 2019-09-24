package models

//InstagramPayload is the data structure based off the format of incoming JSON payloads
type InstagramPayload struct {
	ID            string `json:"id"`
	Caption       string `json:"caption"`
	CommentsCount int    `json:"comments_count"`
	LikeCount     int    `json:"like_count"`
	MediaType     string `json:"media_type"`
	MediaURL      string `json:"media_url"`
	Permalink     string `json:"permalink"`
}
