package models

import "time"
import "encoding/json"

// SocialgistWeiboPayload is the general payload wrapper for Weibo data.
// Either Text.Status OR Test.Comment will be populated, not both.
type SocialgistWeiboPayload struct {
	ID        int64 `json:"id"`
	MatchInfo struct {
		Keyword     string `json:"keyword"`
		KeywordFlag string `json:"keyword_flag"`
		Subid       int    `json:"subid"`
		UID         int    `json:"uid"`
	} `json:"match_info"`
	ReceivedAt string `json:"received_at"`
	Text       struct {
		Event   string                  `json:"event"`
		Status  *SocialgistWeiboStatus  `json:"status"`
		Comment *SocialgistWeiboComment `json:"comment"`
		Type    string                  `json:"type"`
	} `json:"text"`
	URLs []interface{} `json:"urls"`
}

// SocialgistWeiboUser contains information about a Weibo user and is contained
// in SocialgistWeiboStatus and SocialgistWeiboComment
type SocialgistWeiboUser struct {
	AllowAllActMsg  bool   `json:"allow_all_act_msg"`
	AvatarHd        string `json:"avatar_hd"`
	AvatarLarge     string `json:"avatar_large"`
	City            string `json:"city"`
	CityCoordinates string `json:"city_coordinates"`
	CityName        string `json:"city_name"`
	CreatedAt       string `json:"created_at"`
	Description     string `json:"description"`
	Domain          string `json:"domain"`
	FavouritesCount int    `json:"favourites_count"`
	FollowersCount  int    `json:"followers_count"`
	Following       bool   `json:"following"`
	FriendsCount    int    `json:"friends_count"`
	Gender          string `json:"gender"`
	GeoEnabled      bool   `json:"geo_enabled"`
	HasServiceTel   bool   `json:"has_service_tel"`
	ID              int64  `json:"id"`
	Insecurity      struct {
		SexualContent bool `json:"sexual_content"`
	} `json:"insecurity"`
	Location               string `json:"location"`
	Name                   string `json:"name"`
	PagefriendsCount       int    `json:"pagefriends_count"`
	ProfileImageURL        string `json:"profile_image_url"`
	ProfileURL             string `json:"profile_url"`
	Province               string `json:"province"`
	ProvinceCoordinates    string `json:"province_coordinates"`
	ProvinceName           string `json:"province_name"`
	ScreenName             string `json:"screen_name"`
	StatusesCount          int    `json:"statuses_count"`
	Urank                  int    `json:"urank"`
	URL                    string `json:"url"`
	Verified               bool   `json:"verified"`
	VerifiedContactEmail   string `json:"verified_contact_email"`
	VerifiedContactMobile  string `json:"verified_contact_mobile"`
	VerifiedContactName    string `json:"verified_contact_name"`
	VerifiedLevel          int    `json:"verified_level"`
	VerifiedReasonModified string `json:"verified_reason_modified"`
	VerifiedState          int    `json:"verified_state"`
	VerifiedType           int    `json:"verified_type"`
	VerifiedTypeExt        int    `json:"verified_type_ext"`
}

// SocialgistWeiboStatus contains information about a status. Can be contained
// in SocialgistWeiboPayload or as the parent status in SocialgistWeiboComment
type SocialgistWeiboStatus struct {
	AdTags struct {
		Adv []string `json:"adv"`
	} `json:"ad_tags"`
	Annotations []struct {
		ClientMblogid string `json:"client_mblogid"`
		MapiRequest   bool   `json:"mapi_request"`
	} `json:"annotations"`
	APIState      int    `json:"api_state"`
	BizFeature    int    `json:"biz_feature"`
	BmiddlePic    string `json:"bmiddle_pic"`
	CommentsCount int    `json:"comments_count"`
	CreatedAt     string `json:"created_at"`
	ExpireTime    int    `json:"expire_time"`
	ExtendInfo    struct {
		Ad struct {
			ClientMblogid string          `json:"client_mblogid"`
			MapiRequest   bool            `json:"mapi_request"`
			URLMarked     json.RawMessage `json:"url_marked"`
		} `json:"ad"`
		WeiboCamera struct {
			C []string `json:"c"`
		} `json:"weibo_camera"`
	} `json:"extend_info"`
	Favorited           bool   `json:"favorited"`
	ID                  int64  `json:"id"`
	InReplyToScreenName string `json:"in_reply_to_screen_name"`
	InReplyToStatusID   string `json:"in_reply_to_status_id"`
	InReplyToUserID     string `json:"in_reply_to_user_id"`
	IsLongText          bool   `json:"isLongText"`
	IsShowBulletin      int    `json:"is_show_bulletin"`
	IsWeiboVideo        bool   `json:"is_weibo_video"`
	Language            string `json:"language"`
	MblogVipType        int    `json:"mblog_vip_type"`
	Mid                 string `json:"mid"`
	ObjectType          string `json:"objectType"`
	OriginalPic         string `json:"original_pic"`
	PageMarkLevel       int    `json:"pageMarkLevel"`
	ParentRtIDDB        int    `json:"parent_rt_id_db"`
	PicStatus           string `json:"picStatus"`
	PicUrls             []struct {
		ThumbnailPic string `json:"thumbnail_pic"`
	} `json:"pic_urls"`
	RepostsCount    int                    `json:"reposts_count"`
	RetweetedStatus *SocialgistWeiboStatus `json:"retweeted_status"`
	Similarity      string                 `json:"similarity"`
	Source          string                 `json:"source"`
	SourceID        int64                  `json:"source_id"`
	State           int                    `json:"state"`
	StatusURL       string                 `json:"statusurl"`
	Text            string                 `json:"text"`
	ThumbnailPic    string                 `json:"thumbnail_pic"`
	Truncated       bool                   `json:"truncated"`
	User            *SocialgistWeiboUser   `json:"user"`
	Visible         struct {
		ListID int `json:"list_id"`
		Type   int `json:"type"`
	} `json:"visible"`
}

// CreationTime returns creation time of the post
func (weiboStatus *SocialgistWeiboStatus) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RubyDate, weiboStatus.CreatedAt, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}

// SocialgistWeiboComment contains information about a weibo comment. Contains
// the parent SocialgistWeiboStatus in Status
type SocialgistWeiboComment struct {
	CommentURL  string                 `json:"commenturl"`
	CreatedAt   string                 `json:"created_at"`
	FloorNumber int                    `json:"floor_number"`
	ID          int64                  `json:"id"`
	IsTrash     bool                   `json:"isTrash"`
	Language    string                 `json:"language"`
	Mid         string                 `json:"mid"`
	Mode        int                    `json:"mode"`
	Rootid      int                    `json:"rootid"`
	Source      string                 `json:"source"`
	State       int                    `json:"state"`
	Status      *SocialgistWeiboStatus `json:"status"`

	Text string              `json:"text"`
	User SocialgistWeiboUser `json:"user"`
}

// CreationTime returns creation time of the post
func (weiboComment *SocialgistWeiboComment) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RubyDate, weiboComment.CreatedAt, time.UTC)
	if err != nil {
		return invalidDate()
	}
	return parsedTime, nil
}
