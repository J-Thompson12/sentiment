package models

//Contains all of the structures that are shared between twitter and twitter
//direct payloads.

//TwitterEntities is the data structure based off the "entities" format of incoming JSON payloads
type TwitterEntities struct {
	Hashtags []struct {
		Text    string `json:"text"`
		Indices []int  `json:"indices"`
	} `json:"hashtags"`
	Symbols []struct {
		Text    string `json:"text"`
		Indices []int  `json:"indices"`
	} `json:"symbols"`
	Urls []struct {
		URL         string `json:"url"`
		ExpandedURL string `json:"expanded_url"`
		DisplayURL  string `json:"display_url"`
		Indices     []int  `json:"indices"`
	} `json:"urls"`
	UserMentions []struct {
		ScreenName string  `json:"screen_name"`
		Name       string  `json:"name"`
		ID         float64 `json:"id"`
		IDStr      string  `json:"id_str"`
		Indices    []int   `json:"indices"`
	} `json:"user_mentions"`
	Media []TwitterMedia `json:"media"`
}

//TwitterLocation is the data structure based off the "location" format of incoming JSON payloads
type TwitterLocation struct {
	CountryCode string `json:"country_code"`
	DisplayName string `json:"displayName"`
	Geo         struct {
		Coordinates [][][]float64 `json:"coordinates"`
		Type        string        `json:"type"`
	} `json:"geo"`
	Link               string `json:"link"`
	Name               string `json:"name"`
	ObjectType         string `json:"objectType"`
	TwitterCountryCode string `json:"twitter_country_code"`
	TwitterPlaceType   string `json:"twitter_place_type"`
}

//TwitterLongObject is the data structure based off the "long object" format of incoming JSON payloads
type TwitterLongObject struct {
	Body             string          `json:"body"`
	DisplayTextRange []int           `json:"display_text_range"`
	TwitterEntities  TwitterEntities `json:"twitter_entities"`
	ExtendedEntities struct {
		Media []TwitterMedia `json:"media"`
	} `json:"extended_entities"`
}

//TwitterMedia is the data structure based off the "media" format of incoming JSON payloads
type TwitterMedia struct {
	ID            float64 `json:"id"`
	IDStr         string  `json:"id_str"`
	Indices       []int   `json:"indices"`
	MediaURL      string  `json:"media_url"`
	MediaURLHTTPS string  `json:"media_url_https"`
	URL           string  `json:"url"`
	DisplayURL    string  `json:"display_url"`
	ExpandedURL   string  `json:"expanded_url"`
	Type          string  `json:"type"`
	Sizes         struct {
		Small  Size `json:"small"`
		Medium Size `json:"medium"`
		Large  Size `json:"large"`
		Thumb  Size `json:"thumb"`
	} `json:"sizes"`
}

//Size contains the size information for a TwitterMedia structure
type Size struct {
	W      int    `json:"w"`
	H      int    `json:"h"`
	Resize string `json:"resize"`
}
