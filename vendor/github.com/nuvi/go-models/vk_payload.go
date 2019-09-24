package models

import (
	"fmt"
	"time"
)

//VkPayload is the data structure based off the format of incoming JSON payloads
type VkPayload struct {
	ID         int64       `json:"id"`
	Date       int64       `json:"date"`
	OwnerID    int64       `json:"owner_id"`
	FromID     int64       `json:"from_id"`
	PostID     int64       `json:"post_id"`
	PostType   string      `json:"post_type"`
	Text       string      `json:"text"`
	PostSource interface{} `json:"post_source"`
	Comments   interface{} `json:"comments"`
	Likes      struct {
		Count      int `json:"count"`
		UserLikes  int `json:"user_likes"`
		CanLike    int `json:"can_like"`
		CanPublish int `json:"can_publish"`
	} `json:"likes"`
	Reposts struct {
		Count        int `json:"count"`
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Attachments []VkAttachment `json:"attachments"`
	CopyHistory []VkPayload    `json:"copy_history"`
	Geo         struct {
		Type        string `json:"type"` // point/place
		Coordinates string `json:"coordinates"`
		Place       struct {
			ID        int64       `json:"id"`
			Title     string      `json:"title"`
			Latitude  float64     `json:"latitude"`
			Longitude float64     `json:"longitude"`
			Created   int64       `json:"created"`
			Icon      string      `json:"icon"`
			Country   interface{} `json:"country"`
			City      interface{} `json:"city"`
		} `json:"place"`
	} `json:"geo"`
	Profiles []VkProfile `json:"profiles"`
	Groups   []VkGroup   `json:"groups"`
}

//VkProfile is the data structure based off the "profile" format of incoming JSON payloads
type VkProfile struct {
	ID             int64       `json:"id"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Deactivated    string      `json:"deactivated"`
	Hidden         int64       `json:"hidden"`
	Domain         string      `json:"domain"`
	FollowersCount int64       `json:"followers_count"`
	Photo50        interface{} `json:"photo_50"`
	Photo100       string      `json:"photo_100"`
	Counters       interface{} `json:"counters"`
	//Counters       VkGroupProfileCounters `json:"counters"`
}

//VkGroup is the data structure based off the "group" format of incoming JSON payloads
type VkGroup struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	ScreenName   string      `json:"screen_name"`
	IsClosed     int64       `json:"is_closed"`
	Deactivated  string      `json:"deactivated"`
	Type         string      `json:"type"`
	HasPhoto     int64       `json:"has_photo"`
	Photo50      interface{} `json:"photo_50"`
	Photo100     string      `json:"photo_100"`
	Photo200     string      `json:"photo_200"`
	MembersCount int64       `json:"members_count"`
	Counters     interface{} `json:"counters"`
	//Counters     VkGroupProfileCounters `json:"counters"`
}

//VkGroupProfileCounters is the data structure based off the "group profile counters" format of incoming JSON payloads
type VkGroupProfileCounters struct {
	Albums        int `json:"albums"`
	Videos        int `json:"videos"`
	Audios        int `json:"audios"`
	Photos        int `json:"photos"`
	Notes         int `json:"notes"`
	Friends       int `json:"friends"`
	Groups        int `json:"groups"`
	OnlineFriends int `json:"online_friends"`
	MutualFriends int `json:"mutual_friends"`
	UserVideos    int `json:"user_videos"`
	Followers     int `json:"followers"`
	Pages         int `json:"pages"`
}

//VkAttachment is the data structure based off the "attachment" format of incoming JSON payloads
type VkAttachment struct {
	Type        string        `json:"type"`
	Photo       VkPhoto       `json:"photo"`
	Video       VkVideo       `json:"video"`
	Audio       VkAudio       `json:"audio"`
	Doc         VkDoc         `json:"doc"`
	Link        VkLink        `json:"link"`
	Note        VkNote        `json:"note"`
	App         VkApp         `json:"app"`
	Poll        VkPoll        `json:"poll"`
	Page        VkPage        `json:"page"`
	Album       VkAlbum       `jsno:"album"`
	PhotosList  []string      `json:"photos_list"`
	Market      VkMarketItem  `json:"market"`
	MarketAlbum VkMarketAlbum `json:"market_album"`
}

//VkPhoto is the data structure based off the "photo" format of incoming JSON payloads
type VkPhoto struct {
	ID        int64  `json:"id"`
	AlbumID   int64  `json:"album_id"`
	OwnerID   int64  `json:"owner_id"`
	UserID    int64  `json:"user_id"`
	Text      string `json:"text"`
	Date      int64  `json:"date"`
	Photo75   string `json:"photo_75"`
	Photo130  string `json:"photo_130"`
	Photo604  string `json:"photo_604"`
	Photo807  string `json:"photo_807"`
	Photo1280 string `json:"photo_1280"`
	Photo2560 string `json:"photo_2560"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

//VkVideo is the data structure based off the "video" format of incoming JSON payloads
type VkVideo struct {
	ID          int64  `json:"id"`
	OwnerID     int64  `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Photo130    string `json:"photo_130"`
	Photo320    string `json:"photo_320"`
	Photo640    string `json:"photo_640"`
	Photo800    string `json:"photo_800"`
	Date        int64  `json:"date"`
	AddingDate  int64  `json:"adding_date"`
	Views       int    `json:"views"`
	Comments    int    `json:"comments"`
	Player      string `json:"player"`
	AccessKey   string `json:"access_key"`
	Processing  int    `json:"processing"`
	Live        int    `json:"live"`
}

//VkAudio is the data structure based off the "audio" format of incoming JSON payloads
type VkAudio struct {
	ID       int64  `json:"id"`
	OwnerID  int64  `json:"owner_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	URL      string `json:"url"`
	LyricsID int64  `json:"lyrics_id"`
	AlbumID  int64  `json:"album_id"`
	GenreID  int64  `json:"genre_id"`
	Date     int64  `json:"date"`
	NoSearch int    `json:"no_search"`
}

//VkDoc is the data structure based off the "doc" format of incoming JSON payloads
type VkDoc struct {
	ID      int64  `json:"id"`
	OwnerID int64  `json:"owner_id"`
	Title   string `json:"title"`
	Size    int    `json:"size"`
	Ext     string `json:"ext"`
	URL     string `json:"url"`
	Date    int64  `json:"date"`
	Type    int    `json:"type"`
	Preview struct {
		Photo struct {
			Sizes []VkImageSize `json:"sizes"`
		} `json:"photo"`
		Graffity VkImageSize `json:"graffity"`
		AudioMsg struct {
			Duration int    `json:"duration"`
			Waveform []int  `json:"waveform"`
			LinkOgg  string `json:"link_ogg"`
			LinkMp3  string `json:"link_mp3"`
		} `json:"audio_msg"`
	} `json:"preview"`
}

//VkImageSize is the data structure based off the "Image Size" format of incoming JSON payloads
type VkImageSize struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
}

//VkLink is the data structure based off the "link" format of incoming JSON payloads
type VkLink struct {
	URL         string      `json:"url"`
	Title       string      `json:"title"`
	Caption     string      `json:"caption"`
	Description string      `json:"description"`
	Photo       VkPhoto     `json:"photo"`
	IsExternal  int         `json:"is_external"`
	Product     interface{} `json:"product"`
	Rating      VkRating    `json:"rating"`
	Application interface{} `json:"application"`
	Button      VkButton    `json:"button"`
	PreviewPage string      `json:"preview_page"`
	PreviewURL  string      `json:"preview_url"`
}

//VkRating is the data structure based off the "rating" format of incoming JSON payloads
type VkRating struct {
	Stars        float64 `json:"stars"`
	ReviewsCount float64 `json:"reviews_count"`
}

//VkButton is the data structure based off the "button" format of incoming JSON payloads
type VkButton struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

//VkNote is the data structure based off the "note" format of incoming JSON payloads
type VkNote struct {
	ID       int64  `json:"id"`
	OwnerID  int64  `json:"owner_id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Date     int64  `json:"date"`
	Comments int    `json:"comments"`
	ViewURL  string `json:"view_url"`
}

//VkApp is the data structure based off the "app" format of incoming JSON payloads
type VkApp struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

//VkPoll is the data structure based off the format of incoming JSON payloads
type VkPoll struct {
	ID       int64  `json:"id"`
	OwnerID  int64  `json:"owner_id"`
	Created  int64  `json:"created"`
	Question string `json:"question"`
	Votes    int    `json:"votes"`
	AnswerID int64  `json:"answer_id"`
	Answers  []struct {
		ID    int64   `json:"id"`
		Text  string  `json:"text"`
		Votes int     `json:"votes"`
		Rate  float64 `json:"rate"`
	} `json:"answers"`
}

//VkPage is the data structure based off the "page" format of incoming JSON payloads
type VkPage struct {
	ID                       int64  `json:"id"`
	GroupID                  int64  `json:"group_id"`
	CreatorID                int64  `json:"creator_id"`
	Title                    string `json:"title"`
	CurrentUserCanEdit       int    `json:"current_user_can_edit"`
	CurrentUserCanEditAccess int    `json:"current_user_can_edit_access"`
	WhoCanView               int    `json:"who_can_view"`
	WhoCanEdit               int    `json:"who_can_edit"`
	Edited                   int64  `json:"edited"`
	Created                  int64  `json:"created"`
	EditorID                 int64  `json:"editor_id"`
	Views                    int    `json:"views"`
	Parent                   string `json:"parent"`
	Parent2                  string `json:"parent2"`
	Source                   string `json:"source"`
	HTML                     string `json:"html"`
	ViewURL                  string `json:"view_url"`
}

//VkAlbum is the data structure based off the "alubm" format of incoming JSON payloads
type VkAlbum struct {
	ID          string  `json:"id"`
	Thumb       VkPhoto `json:"thumb"`
	OwnerID     int64   `json:"owner_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Created     int64   `json:"created"`
	Updated     int64   `json:"updated"`
	Size        int     `json:"size"`
}

//VkMarketItem is the data structure based off the "market item" format of incoming JSON payloads
type VkMarketItem struct {
	ID           int64       `json:"id"`
	OwnerID      int64       `json:"owner_id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Price        interface{} `json:"price"`
	Category     interface{} `json:"category"`
	ThumbPhoto   string      `json:"thumb_photo"`
	Date         int64       `json:"date"`
	Availability int         `json:"availability"`
	Photos       []VkPhoto   `json:"photos"`
	CanComment   int         `json:"can_comment"`
	CanRepost    int         `json:"can_repost"`
	Likes        struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
}

//VkMarketAlbum is the data structure based off the "Market Album" format of incoming JSON payloads
type VkMarketAlbum struct {
	ID          int64   `json:"id"`
	OwnerID     int64   `json:"owner_id"`
	Title       string  `json:"title"`
	Photo       VkPhoto `json:"photo"`
	Count       int     `json:"count"`
	UpdatedTime int64   `json:"updated_time"`
	Size        int     `json:"size"`
}

// PermalinkURL returns a permanent link to the post
func (vkPayload *VkPayload) PermalinkURL() string {
	return "https://vk.com/" + vkPayload.ActivityID()
}

// ActivityID returns a unique ID of the post on VK
func (vkPayload *VkPayload) ActivityID() string {
	if vkPayload.PostType == "reply" {
		return fmt.Sprintf("wall%d_%d?reply=%d", vkPayload.OwnerID, vkPayload.PostID, vkPayload.ID)
	}
	if vkPayload.PostType == "photo" {
		return fmt.Sprintf("photo%d_%d", vkPayload.OwnerID, vkPayload.ID)
	}
	if vkPayload.PostType == "video" {
		return fmt.Sprintf("video%d_%d", vkPayload.OwnerID, vkPayload.ID)
	}
	return fmt.Sprintf("wall%d_%d", vkPayload.OwnerID, vkPayload.ID)
}

// ParentID returns a unique ID of the parent post
func (vkPayload *VkPayload) ParentID() string {
	if vkPayload.PostID != 0 && vkPayload.PostType == "reply" {
		return fmt.Sprintf("wall%d_%d", vkPayload.OwnerID, vkPayload.PostID)
	}
	if len(vkPayload.CopyHistory) > 0 {
		return vkPayload.CopyHistory[0].ActivityID()
	}
	return ""
}

// CreationTime returns creation time of the post
func (vkPayload *VkPayload) CreationTime() (time.Time, error) {
	return time.Unix(vkPayload.Date, 0).UTC(), nil
}
