package models

import (
	"time"
)

//TumblrPayload is the data structure based off the format of incoming JSON payloads
type TumblrPayload struct {
	Version         string `json:"version"`
	ID              string `json:"id"`
	Timestamp       int64  `json:"timestamp"`
	ActivityType    string `json:"activity_type"`
	ActivityPrivacy string `json:"activity_privacy"`
	Activity        struct {
		Action            string `json:"action"`
		UserPrimaryBlogID int    `json:"user_primary_blog_id"`
		Post              struct {
			Type         string `json:"type"`
			ID           int64  `json:"id"`
			ShortURL     string `json:"short_url"`
			Created      int64  `json:"created"`
			Updated      int64  `json:"updated"`
			IsSubmission bool   `json:"is_submission"`
			Blog         struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				URL         string `json:"url"`
				Title       string `json:"title"`
				Description string `json:"description"`
				IsGroupBlog bool   `json:"is_group_blog"`
			} `json:"blog"`
			ParentPost struct {
				ID       int64  `json:"id"`
				ShortURL string `json:"short_url"`
				Blog     struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"blog"`
			} `json:"parent_post"`
			RootPost struct {
				ID       int64  `json:"id"`
				ShortURL string `json:"short_url"`
				Blog     struct {
					ID int `json:"id"`
				} `json:"blog"`
			} `json:"root_post"`
			Content struct {
				NoteCount      int      `json:"note_count"`
				Tags           []string `json:"tags"`
				Slug           string   `json:"slug"`
				Format         string   `json:"format"`
				SourceTitle    string   `json:"source_title"`
				SourceURL      string   `json:"source_url"`
				Mentions       []string `json:"mentions"` // array of useless numeric strings
				Name           string   `json:"name"`
				Title          string   `json:"title"`
				Body           string   `json:"body"`
				AcceptsAnswers bool     `json:"accepts_answers"`
				Text           string   `json:"text"`
				Source         string   `json:"source"`
				Asking         struct {
					IsAnonymous       bool `json:"is_anonymous"`
					UserPrimaryBlogID int  `json:"user_primary_blog_id"`
				}
				Question string `json:"question"`
				Answer   string `json:"Answer"`
				Photos   []struct {
					IsPanorama bool   `json:"is_panorama"`
					Caption    string `json:"caption"`
					AltSizes   []struct {
						Width  int    `json:"width"`
						Height int    `json:"height"`
						URL    string `json:"url"`
					} `json:"alt_sizes"`
					OriginalSize struct {
						URL    string `json:"url"`
						Height int    `json:"height"`
						Width  int    `json:"width"`
					} `json:"original_size"`
					Exif struct {
						FocalLength string `json:"focal_length"`
						Exposure    string `json:"exposure"`
						Camera      string `json:"camera"`
						Iso         string `json:"iso"`
						Aperture    string `json:"aperture"`
					} `json:"exif"`
				} `json:"photos"`
				AudioType  string `json:"audio_type"`
				AudioURL   string `json:"audio_url"`
				IsExternal bool   `json:"is_external"`
				Caption    string `json:"caption"`
				Plays      int    `json:"plays"`
				Additional struct {
					Artist    string      `json:"artist"`
					Album     string      `json:"album"`
					Year      int         `json:"year"`
					Track     interface{} `json:"track"`
					TrackName string      `json:"track_name"`
					AlbumArt  string      `json:"album_art"`
				} `json:"additional"`
				VideoType string `json:"video_type"`
				VideoURL  string `json:"video_url"`
				Duration  int    `json:"duration"`
				Thumbnail struct {
					URL    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				}
				Dialogue []struct {
					Name   string `json:"name"`
					Label  string `json:"label"`
					Phrase string `json:"phrase"`
				}
				Link        string `json:"link"`
				LinkImage   string `json:"link_image"`
				Description string `json:"description"`
			} `json:"content"`
		} `json:"post"`
	} `json:"activity"`
	ReblogKey string `json:"reblog_key"`
}

// IsReshare checks if the payload represents a repost
func (tumblrPayload *TumblrPayload) IsReshare() bool {
	if tumblrPayload.Activity.Post.ParentPost.ID > 0 || tumblrPayload.Activity.Post.RootPost.ID == 0 {
		return true
	}
	return false
}

// CreationTime returns creation time of the post
func (tumblrPayload *TumblrPayload) CreationTime() (time.Time, error) {
	parsedTime := time.Unix(tumblrPayload.Timestamp, 0).UTC()

	return parsedTime, nil
}
