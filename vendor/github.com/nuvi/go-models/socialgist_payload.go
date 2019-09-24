package models

import (
	"time"
)

// SocialgistBlogsFirehosePayload is the data structure for the Socialgist Blog firehose
type SocialgistBlogsFirehosePayload struct {
	Post    *SocialgistBlogsFirehosePost    `json:"post"`
	Comment *SocialgistBlogsFirehoseComment `json:"comment"`
}

// SocialgistBlogsFirehosePost represents a social gist element received from a firehose
type SocialgistBlogsFirehosePost struct {
	Link struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		Type string `json:"type"`
	} `json:"Link"`
	AuthorEmail string `json:"authoremail"`
	AuthorName  string `json:"authorname"`
	AuthorURL   string `json:"authorurl"`
	BloghostID  string `json:"bloghostid"`
	BlogID      string `json:"blogid"`
	BlogTitle   string `json:"blogtitle"`
	Category    string `json:"category"`
	Content     string `json:"content"`
	Country     string `json:"country"`
	Generator   string `json:"generator"`
	GUID        string `json:"guid"`
	Lang        string `json:"lang"`
	MainURL     string `json:"mainurl"`
	ParsedDate  string `json:"parseddate"`
	PostID      string `json:"postid"`
	PostLink    string `json:"postlink"`
	Published   string `json:"pubdate"`
	Source      string `json:"source"`
	Title       string `json:"title"`
	Updated     string `json:"updated"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialgistBlogsFirehosePost) CreationTime() (time.Time, error) {
	if socialgistPayload.Published == "" {
		return time.Now().In(time.UTC), nil
	}
	parsedTime, err := time.ParseInLocation("2006-01-02T15:04:05", socialgistPayload.Published, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

// SocialgistBlogsFirehoseComment represents a comment received from the socialgist firehose
type SocialgistBlogsFirehoseComment struct {
	AuthorEmail     string `json:"authoremail"`
	AuthorName      string `json:"authorname"`
	AuthorURL       string `json:"authorurl"`
	BlogHostID      string `json:"bloghostid"`
	CommentID       string `json:"commentid"`
	CommentLink     string `json:"commentlink"`
	Content         string `json:"content"`
	Country         string `json:"country"`
	DiscoveryMethod string `json:"discoveryMethod"`
	Generator       string `json:"generator"`
	Lang            string `json:"lang"`
	ParsedDate      string `json:"parseddate"`
	PostGUID        string `json:"postguid"`
	PostLink        string `json:"postlink"`
	PostPublished   string `json:"postpublished"`
	PostTitle       string `json:"posttitle"`
	Provider        string `json:"provider"`
	Published       string `json:"pubdate"`
	SourceCrawled   string `json:"sourcecrawled"`
	SourceGUID      string `json:"sourceguid"`
	SourceID        string `json:"sourceid"`
	SourceLanguage  string `json:"sourcelanguage"`
	SourceTitle     string `json:"sourcetitle"`
	SourceType      string `json:"sourcetype"`
	SourceURL       string `json:"sourceurl"`
	Title           string `json:"title"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialgistBlogsFirehoseComment) CreationTime() (time.Time, error) {
	if socialgistPayload.Published == "" {
		return time.Now().In(time.UTC), nil
	}
	parsedTime, err := time.ParseInLocation("2006-01-02T15:04:05", socialgistPayload.Published, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

// SocialgistNewsFirehosePayload is the data structure for the Socialgist News firehose
type SocialgistNewsFirehosePayload struct {
	Article struct {
		CaNationalots string `json:"ca_nationalots"`
		CaNewsrank    string `json:"ca_newsrank"`
		CaPrintcir    string `json:"ca_printcir"`
		CaReachpermil string `json:"ca_reachpermil"`
		CaWebrank     string `json:"ca_webrank"`
		Description   struct {
			Author        string `json:"author"`
			Charset       string `json:"charset"`
			Content       string `json:"content"`
			HltextDisplay string `json:"hltext_display"`
			Language      string `json:"language"`
			LanguageCode  string `json:"language_code"`
		} `json:"description"`
		DocURL      string `json:"docurl"`
		HarvestTime string `json:"harvest_time"`
		ID          string `json:"id"`
		Location    struct {
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			ZipCode     string `json:"zip_code"`
		} `json:"location"`
		OriginalLink   string `json:"original_link"`
		ProviderID     string `json:"providerid"`
		Source         string `json:"source"`
		SourceCategory string `json:"source_category"`
	} `json:"article"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialgistNewsFirehosePayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", socialgistPayload.Article.HarvestTime, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

//SocialgistRSSPayload is the data structure for Socialgist Blogs and News responses
type SocialgistRSSPayload struct {
	AuthorInfo struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
		Nick string `json:"Nick"`
		URL  string `json:"Url"`
	} `json:"AuthorInfo"`
	CommentsInThread interface{} `json:"CommentsInThread"`
	ContentSource    string      `json:"ContentSource"`
	Country          string      `json:"Country"`
	Crawled          string      `json:"Crawled"`
	FeedInfo         struct {
		ExtKey string `json:"ExtKey"`
		ID     string `json:"Id"`
		Title  string `json:"Title"`
		URL    string `json:"Url"`
	} `json:"FeedInfo"`
	ID              string `json:"Id"`
	Inserted        string `json:"Inserted"`
	IsAdult         string `json:"IsAdult"`
	IsComment       int    `json:"IsComment"`
	Language        string `json:"Language"`
	MatchingQueries []struct {
		Query   string `json:"query"`
		Queryid string `json:"queryid"`
	} `json:"MatchingQueries"`
	PostSize    int    `json:"PostSize"`
	PostTitle   string `json:"PostTitle"`
	Published   string `json:"Published"`
	Subject     string `json:"Subject"`
	SubjectHTML string `json:"SubjectHtml"`
	Tags        string `json:"Tags"`
	Text        string `json:"Text"`
	TextHTML    string `json:"TextHtml"`
	ThreadID    string `json:"ThreadId"`
	Type        string `json:"Type"`
	URL         string `json:"Url"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialgistRSSPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", socialgistPayload.Published, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

//SocialgistBoardPayload is the data structure for Socialgist Reddit responses.
type SocialgistBoardPayload struct {
	Author          string `json:"Author"`
	AuthorAvatatURL string `json:"AuthorAvatarUrl"`
	AuthorID        string `json:"AuthorId"`
	AuthorInfo      struct {
		Age        string `json:"Age"`
		KloutScore string `json:"KloutScore"`
		Location   string `json:"Location"`
		Name       string `json:"Name"`
		Registered string `json:"Registered"`
		Sex        string `json:"Sex"`
		URL        string `json:"Url"`
	} `json:"AuthorInfo"`
	BBSType         string `json:"BBSType"` //Socialgist Internal Use Only
	Country         string `json:"Country"`
	ContentSource   string `json:"ContentSource"`
	Crawled         string `json:"Crawled"`
	Domain          string `json:"Domain"`
	ExtKey          string `json:"ExtKey"` //Socialgist Internal Use Only
	Forum           string `json:"Forum"`
	ForumID         string `json:"ForumKey"`
	ForumURL        string `json:"FormUrl"`
	TimezoneOffset  string `json:"Gmt"`
	Grouped         int    `json:"Grouped"`
	Icon            string `json:"Icon"`
	ID              string `json:"Id"`
	Inserted        string `json:"Inserted"`
	IsAdult         string `json:"IsAdult"`
	IsThread        string `json:"IsThread"`
	IsThreadStart   string `json:"IsThreadStart"`
	Language        string `json:"Language"`
	MatchingQueries []struct {
		Query   string `json:"query"`
		Queryid string `json:"queryid"`
	} `json:"MatchingQueries"`
	PostSize      int    `json:"PostSize"`
	PostsInThread string `json:"PostsInThread"`
	Published     string `json:"Published"`
	Signature     string `json:"Signature"`
	SignatureHTML string `json:"SignatureHtml"`
	SiteID        string `json:"SiteId"`
	SiteKey       string `json:"SiteKey"`
	SiteTitle     string `json:"SiteTitle"`
	SiteURL       string `json:"SiteUrl"`
	Subject       string `json:"Subject"`
	SubjectHTML   string `json:"SubjectHtml"`
	Text          string `json:"Text"`
	TextHTML      string `json:"TextHtml"`
	ThreadID      string `json:"ThreadId"`
	ThreadTitle   string `json:"ThreadTitle"`
	URL           string `json:"Url"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialgistBoardPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", socialgistPayload.Published, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

//SocialgistVideoPayload is the data structure for Socialgist Youtube responses
type SocialgistVideoPayload struct {
	AuthorInfo struct {
		ID           string `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Name         string `json:"Name"`
		Nick         string `json:"Nick"`
		SiteAuthorID string `json:"SiteAuthorID"`
		URL          string `json:"Url"`
	} `json:"AuthorInfo"`
	Category         string `json:"Category"`
	CommentsInThread string `json:"CommentsInThread"`
	ContentSource    string `json:"ContentSource"`
	Crawled          string `json:"Crawled"`
	Duration         int    `json:"Duration"`
	ID               string `json:"Id"`
	Inserted         string `json:"Inserted"`
	IsComment        int    `json:"IsComment"`
	Language         string `json:"Language"`
	MatchingQueries  []struct {
		Query   string `json:"query"`
		Queryid string `json:"queryid"`
	} `json:"MatchingQueries"`
	PostSize  int    `json:"PostSize"`
	Published string `json:"Published"`
	SiteInfo  struct {
		ID      string `json:"Id"`
		Name    string `json:"Name"`
		SiteURL string `json:"SiteUrl"`
	} `json:"SiteInfo"`
	Tags         string `json:"Tags"`
	ThumbnailURL string `json:"ThumbnailUrl"`
	Title        string `json:"Title"`
	TitleHTML    string `json:"TitleHtml"`
	Text         string `json:"Text"`
	TextHTML     string `json:"TextHtml"`
	ThreadID     string `json:"ThreadId"`
	URL          string `json:"Url"`
	VideoTitle   string `json:"VideoTitle"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialgistVideoPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", socialgistPayload.Published, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

//SocialgistReviewPayload is the data structure for Socialgist Review responses.
type SocialgistReviewPayload struct {
	AttributeRatings interface{} `json:"AttributeRatings"` //AttributeRatings includes an arbitrary number of arbitrarily named fields
	AuthorInfo       struct {
		Age         string `json:"Age"`
		AvatarURL   string `json:"AvatarUrl"`
		Description string `json:"Description"`
		ID          string `json:"Id"`
		Location    string `json:"Location"`
		Name        string `json:"Name"`
		Sex         string `json:"Sex"`
		URL         string `json:"Url"`
		UserName    string `json:"UserName"`
	} `json:"AuthorInfo"`
	ContentSource string `json:"ContentSource"`
	Crawled       string `json:"Crawled"`
	ID            string `json:"Id"`
	Inserted      string `json:"Inserted"`
	ItemInfo      struct {
		ID          string `json:"Id"`
		Category    string `json:"Category"`
		CategoryURL string `json:"CategoryUrl"`
		Item        string `json:"Item"`
		ItemHTML    string `json:"ItemHtml"`
		ItemURL     string `json:"ItemURL"`
	} `json:"ItemInfo"`
	Language        string `json:"Language"`
	MatchingQueries []struct {
		Query   string `json:"query"`
		Queryid string `json:"queryid"`
	} `json:"MatchingQueries"`
	OriginalSiteURL string `json:"OriginalSiteUrl"`
	OriginalURL     string `json:"OriginalUrl"`
	PostSize        int    `json:"PostSize"`
	Published       string `json:"Published"`
	Recommend       string `json:"Recommend"`
	ReviewRating    string `json:"ReviewRating"`
	ReviewType      string `json:"ReviewType"`
	SiteInfo        struct {
		Country string `json:"Country"`
		ExtKey  string `json:"ExtKey"` //Socialgist Internal Use Only
		ID      string `json:"Id"`
		Name    string `json:"Name"`
		SiteURL string `json:"SiteUrl"`
	} `json:"SiteInfo"`
	Text         string `json:"Text"`
	TextCons     string `json:"TextCons"`
	TextConsHTML string `json:"TextConsHtml"`
	TextHTML     string `json:"TextHtml"`
	TextPros     string `json:"TextPros"`
	TextProsHTML string `json:"TextProsHtml"`
	Title        string `json:"Title"`
	TitleHTML    string `json:"TitleHtml"`
	URL          string `json:"Url"`
}

// CreationTime returns creation time of the post
func (socialgistPayload *SocialgistReviewPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", socialgistPayload.Published, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

// SocialgistVideoPayloadV2 is the new data structure for Socialgist youtube videos, comments, and stats.
// Socialgist sends three different message types on the single video connection. The can be decoded into
// this single type.
type SocialgistVideoPayloadV2 struct {
	Comment SocialgistVideoCommentsV2 `json:"comment"`
	Video   SocialgistVideoV2         `json:"video"`
}

// SocialgistVideoV2 carries information about a video
type SocialgistVideoV2 struct {
	SocialgistVideoCommonV2

	ID          string   `json:"site_videoid"`
	Description string   `json:"description"`
	ThumbURL    string   `json:"thumb_url_code"`
	Tags        []string `json:"tags"`
	Category    string   `json:"category"`
	Duration    int      `json:"duration,string"`
	URL         string   `json:"videourl"`
}

// SocialgistVideoCommentsV2 carries information about a comment made on a video
type SocialgistVideoCommentsV2 struct {
	SocialgistVideoCommonV2

	ID        string `json:"site_commentid"`
	ParentID  string `json:"videoid"`
	Content   string `json:"content"`
	ParentURL string `json:"videourl"`
	URL       string `json:"commentsurl"`
}

// SocialgistVideoCommonV2 carries info common to comments and videos
type SocialgistVideoCommonV2 struct {
	Network   string    `json:"siteid"`
	Title     string    `json:"title"`
	Published time.Time `json:"published"`
	Crawled   time.Time `json:"crawled"`
	Language  string    `json:"lang"`
	Author    struct {
		ID   string `json:"site_authorid"`
		Name string `json:"name"`
		URL  string `json:"authorurl"`
	}
}

// SocialgistVideoPayloadV2Type represents the payload type for a SocialgistVideoPayloadV2
type SocialgistVideoPayloadV2Type int

// The SocialgistVideoPayloadV2Types
const (
	SocialgistVideoPayloadV2TypeVideo = iota
	SocialgistVideoPayloadV2TypeComment
	SocialgistVideoPayloadV2TypeStat
)

// Type returns the type of payload
func (p SocialgistVideoPayloadV2) Type() SocialgistVideoPayloadV2Type {
	switch {
	case p.Comment.ID != "":
		return SocialgistVideoPayloadV2TypeComment
	case p.Video.ID != "":
		return SocialgistVideoPayloadV2TypeVideo
	default:
		return SocialgistVideoPayloadV2TypeStat
	}
}
