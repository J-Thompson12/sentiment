package models

import (
	"strings"
	"time"
)

//PinterestPayload is the data structure based off the format of incoming JSON payloads
type PinterestPayload struct {
	GravatarID  string `json:"gravatar_id"`
	Attribution struct {
		Title              string `json:"title"`
		URL                string `json:"url"`
		ProviderIconURL    string `json:"provider_icon_url"`
		AuthorName         string `json:"author_name"`
		ProviderFaviconURL string `json:"provider_favicon_url"`
		AuthorURL          string `json:"author_url"`
		ProviderName       string `json:"provider_name"`
	} `json:"attribution"`
	Creator struct {
		URL       string `json:"url"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		ID        string `json:"id"`
	} `json:"creator"`
	URL   string `json:"url"`
	Media struct {
		Type string `json:"type"`
	} `json:"media"`
	CreatedAt    string `json:"created_at"`
	OriginalLink string `json:"original_link"`
	Note         string `json:"note"`
	Color        string `json:"color"`
	Link         string `json:"link"`
	Board        struct {
		URL  string `json:"url"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"board"`
	Image struct {
		Original struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"original"`
	} `json:"image"`
	Counts struct {
		Comments int `json:"comments"`
		Saves    int `json:"saves"`
	} `json:"counts"`
	ID       string `json:"id"`
	Metadata struct {
		Article struct {
			PublishedAt interface{} `json:"published_at"`
			Description string      `json:"description"`
			Name        string      `json:"name"`
			Authors     []struct {
				Name string `json:"name"`
			} `json:"authors"`
		} `json:"article"`
		Link struct {
			Locale      string `json:"locale"`
			Title       string `json:"title"`
			SiteName    string `json:"site_name"`
			Description string `json:"description"`
			Favicon     string `json:"favicon"`
		} `json:"link"`
	} `json:"metadata"`
}

// GenerateFullName returns the author's full name
func (pinterestPayload *PinterestPayload) GenerateFullName() (fullName string) {
	return strings.Trim(pinterestPayload.Creator.FirstName+" "+pinterestPayload.Creator.LastName, " ")
}

// CreationTime returns creation time of the post if present or the time of harvest
func (pinterestPayload *PinterestPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02T15:04:05", pinterestPayload.CreatedAt, time.UTC)

	if err != nil {
		return invalidDate()
	}

	return parsedTime, nil

}
