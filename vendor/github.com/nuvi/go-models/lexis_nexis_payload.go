package models

import "time"

// LexisNexisSemanticsItem represents events and entities in LexisNexisPayload.Semantics
type LexisNexisSemanticsItem struct {
	Properties []struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"properties,omitempty"`
	Provider string `json:"provider"`
}

// LexisNexisPayload is the data structure based off the format of incoming JSON payloads
type LexisNexisPayload struct {
	SequenceID            string   `json:"sequenceId,omitempty"`
	ID                    string   `json:"id"`
	Language              string   `json:"language"` // "English"
	LanguageCode          string   `json:"languageCode,omitempty"`
	Title                 string   `json:"title"`
	SubTitle              string   `json:"subTitle,omitempty"`
	Content               string   `json:"content,omitempty"`
	ContentWithMarkup     string   `json:"contentWithMarkup,omitempty"`
	Tags                  []string `json:"tags,omitempty"`
	PublishedDate         string   `json:"publishedDate"`
	HarvestDate           string   `json:"harvestDate"`
	EmbargoDate           string   `json:"embargoDate,omitempty"`
	LicenseEndDate        string   `json:"licenseEndDate"`
	ContentLicenseEndDate string   `json:"contentLicenseEndDate,omitempty"`
	URL                   string   `json:"url"`
	OriginalURL           string   `json:"originalUrl,omitempty"`
	CommentsURL           string   `json:"commentsUrl"`
	OutboundUrls          []string `json:"outboundUrls,omitempty"`
	WordCount             string   `json:"wordCount,omitempty"`
	DataFormat            string   `json:"dataFormat"` //"text", "audio", "video", and "image"
	Copyright             string   `json:"copyright"`
	LoginStatus           string   `json:"loginStatus"` // "reg" = requires free registration, "sub" = requires paid subscription, "prem" = premium content site
	DuplicateGroupID      string   `json:"duplicateGroupId"`
	ContentGroupIds       string   `json:"contentGroupIds,omitempty"`
	Harvest               struct {
		LegacyNewsFeedID           string   `json:"legacyNewsFeedId,omitempty"`
		LegacyNewsName             string   `json:"legacyNewsName,omitempty"`
		HarvestingMethods          string   `json:"harvestingMethods,omitempty"`
		HarvesterServerID          string   `json:"harvesterServerId,omitempty"`
		ProcessingTime             string   `json:"processingTime,omitempty"`
		StreamReaderName           string   `json:"streamReaderName,omitempty"`
		StreamReaderStartTime      string   `json:"streamReaderStartTime,omitempty"`
		StreamReaderFileName       string   `json:"streamReaderFileName,omitempty"`
		StreamReaderSearchTermName string   `json:"streamReaderSearchTermName,omitempty"`
		CustomerTags               []string `json:"customerTags,omitempty"`
		StreamReaderLNFileNum      string   `json:"streamReaderLNFileNum"`
		RobotStatus                string   `json:"robotStatus,omitempty"`
	} `json:"harvest"`
	Media struct {
		Images []struct {
			URL      string `json:"url"`
			MimeType string `json:"mimeType"`
			Caption  string `json:"caption"`
		} `json:"images"`
		Video []struct {
			Duration string `json:"duration"`
			URL      string `json:"url"`
			MimeType string `json:"mimeType"`
			Caption  string `json:"caption"`
		} `json:"video"`
		Audio []struct {
			Duration string `json:"duration"`
			URL      string `json:"url"`
			MimeType string `json:"mimeType"`
			Caption  string `json:"caption"`
		} `json:"audio"`
	} `json:"media"`
	PublishingPlatform struct {
		ItemID            string   `json:"itemId"`
		OriginalItemID    string   `json:"originalItemId"`
		StatusID          string   `json:"statusId"`
		ItemType          string   `json:"itemType"`
		InReplyToUserID   string   `json:"inReplyToUserId"`
		InReplyToStatusID string   `json:"inReplyToStatusId"`
		UserMentions      []string `json:"userMentions,omitempty"`
		TotalViews        string   `json:"totalViews"`
		ShareCount        string   `json:"shareCount"`
	} `json:"publishingPlatform"`
	AdultLanguage string `json:"adultLanguage"` // "true" or "false"
	Topics        []struct {
		Name  string `json:"name,omitempty"`
		Group string `json:"group,omitempty"`
	} `json:"topics"`
	IndexTerms []struct {
		Domains []string `json:"domains,omitempty"`
		Name    string   `json:"name"`
		Score   string   `json:"score"`
		Code    string   `json:"code,omitempty"`
	} `json:"indexTerms"`
	Companies []struct {
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Exchange     string `json:"exchange"`
		Isin         string `json:"isin"`
		Cusip        string `json:"cusip,omitempty"`
		TitleCount   int    `json:"titleCount"`
		ContentCount int    `json:"contentCount"`
		Primary      bool   `json:"primary"`
	} `json:"companies,omitempty"`
	Locations []struct {
		Name       string `json:"name"`
		Type       string `json:"type"`
		Class      string `json:"class"`
		Mentions   string `json:"mentions"`
		Confidence string `json:"confidence"`
		Country    struct {
			Confidence string `json:"confidence,omitempty"`
			FipsCode   string `json:"fipsCode,omitempty"`
			IsoCode    string `json:"isoCode,omitempty"`
			Name       string `json:"name,omitempty"`
		} `json:"country"`
		Region    string `json:"region"`
		Subregion string `json:"subregion"`
		State     struct {
			Confidence string `json:"confidence,omitempty"`
			FipsCode   string `json:"fipsCode,omitempty"`
			Name       string `json:"name,omitempty"`
		} `json:"state"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
		Provider  string `json:"provider"`
	} `json:"locations,omitempty"`
	Semantics struct {
		Events   []LexisNexisSemanticsItem `json:"events,omitempty"`
		Entities []LexisNexisSemanticsItem `json:"entities,omitempty"`
	} `json:"semantics,omitempty"`
	Sentiment struct {
		Score    string `json:"score,omitempty"`
		Entities []struct {
			Type      string `json:"type,omitempty"`
			Value     string `json:"value,omitempty"`
			Mentions  string `json:"mentions,omitempty"`
			Score     string `json:"score,omitempty"`
			Evidence  string `json:"evidence,omitempty"`
			Confident bool   `json:"confident"`
		} `json:"entities,omitempty"`
	} `json:"sentiment,omitempty"`
	Print struct {
		Supplement         string `json:"supplement"`
		PublicationEdition string `json:"publicationEdition"`
		RegionalEdition    string `json:"regionalEdition"`
		Section            string `json:"section"`
		SectionFull        string `json:"sectionFull"`
		PageNumber         string `json:"pageNumber"`
		PageCount          string `json:"pageCount"`
		SizeCm             string `json:"sizeCm"`
		SizePercentage     string `json:"sizePercentage"`
		OriginLeft         string `json:"originLeft"`
		OriginTop          string `json:"originTop"`
		Width              string `json:"width"`
		Height             string `json:"height"`
		ByLine             string `json:"byLine"`
		Photo              string `json:"photo"`
	} `json:"print,omitempty"`
	Broadcase struct {
		MarketName      string `json:"marketName"`
		NationalNetwork string `json:"nationalNetwork"`
		Title           string `json:"title"`
		ProgramOrigin   string `json:"programOrigin"`
		ProgramCategory string `json:"programCategory"`
		Station         string `json:"station"`
		Lines           []struct {
			Date string `json:"date,omitempty"`
			Text string `json:"text,omitempty"`
		} `json:"lines,omitempty"`
	} `json:"broadcast"`
	Author struct {
		Name               string `json:"name"`
		HomeURL            string `json:"homeUrl"`
		Email              string `json:"email"`
		Description        string `json:"description"`
		DateLastActive     string `json:"dateLastActive"`
		PublishingPlatform struct {
			UserName       string `json:"userName"`
			UserID         string `json:"userId"`
			StatusesCount  string `json:"statusesCount"`
			TotalViews     string `json:"totalViews"`
			FollowingCount string `json:"followingCount"`
			FollowersCount string `json:"followersCount"`
			KloutScore     string `json:"kloutScore"`
		} `json:"publishingPlatform"`
	} `json:"author"`
	Licenses []struct {
		Name string `json:"name,omitempty"`
	} `json:"licenses,omitempty"`
	LinkedArticles []struct {
		Type      string `json:"type,omitempty"`
		ArticleID string `json:"articleId,omitempty"`
	} `json:"linkedArticles,omitempty"`
	AdvertisingValueEquivalency struct {
		Value    string `json:"value,omitempty"`
		Currency string `json:"currency,omitempty"`
	} `json:"advertisingValueEquivalency,omitempty"`
	Source struct {
		ID                    string `json:"id,omitempty"`
		SyndicatorSourceID    string `json:"syndicatorSourceId,omitempty"`
		Name                  string `json:"name"`
		HomeURL               string `json:"homeUrl"`
		Publisher             string `json:"publisher"`
		PrimaryLanguage       string `json:"primaryLanguage"`
		PrimaryMediaType      string `json:"primaryMediaType"`
		Category              string `json:"category"`
		EditorialRank         string `json:"editorialRank"`
		PublicationID         string `json:"publicationId,omitempty"`
		LicensorPublicationID string `json:"licensorPublicationId,omitempty"`
		ChannelCode           string `json:"channelCode,omitempty"`
		AggregatorID          string `json:"aggregatorId,omitempty"`
		AggregatorName        string `json:"aggregatorName,omitempty"`
		Location              struct {
			Country     string `json:"country"`
			CountryCode string `json:"countryCode"`
			Region      string `json:"region"`
			Subregion   string `json:"subregion"`
			State       string `json:"state"`
			County      string `json:"county,omitempty"`
			ZipArea     string `json:"zipArea"`
			ZipCode     string `json:"zipCode"`
		} `json:"location"`
		Metrics struct {
			Mozscape struct {
				MozRank         string `json:"mozRank,omitempty"`
				PageAuthority   string `json:"pageAuthority,omitempty"`
				DomainAuthority string `json:"domainAuthority,omitempty"`
				ExternalLinks   string `json:"externalLinks,omitempty"`
				Links           string `json:"links,omitempty"`
			} `json:"mozscape,omitempty"`
			ReachMetrics []struct {
				Type        string `json:"type,omitempty"`
				Value       string `json:"value,omitempty"`
				Frequency   string `json:"frequency,omitempty"`
				CountryCode string `json:"countryCode,omitempty"`
			} `json:"reachMetrics,omitempty"`
		} `json:"metrics"`
		FeedSource string `json:"feedSource,omitempty"`
		Feed       struct {
			ID                 string   `json:"id"`
			Name               string   `json:"name"`
			URL                string   `json:"url,omitempty"`
			MediaType          string   `json:"mediaType"`
			PublishingPlatform string   `json:"publishingPlatform"`
			IDFromPublisher    string   `json:"idFromPublisher"`
			Generator          string   `json:"generator"`
			Description        string   `json:"description"`
			Tags               []string `json:"tags,omitempty"`
			ImageURL           string   `json:"imageUrl"`
			Copyright          string   `json:"copyright"`
			Language           string   `json:"language"`
			DataFormat         string   `json:"dataFormat"`
			Rank               struct {
				AutoRank         string `json:"autoRank"`
				AutoRankOrder    string `json:"autoRankOrder"`
				InboundLinkCount string `json:"inboundLinkCount"`
			} `json:"rank"`
			InWhiteList     string   `json:"inWhiteList"`
			AutoTopics      []string `json:"autoTopics"`
			EditorialTopics []string `json:"editorialTopics"`
			Genre           string   `json:"genre"`
		} `json:"feed"`
	} `json:"source"`
	MatchingFilters []string `json:"matchingFilters,omitempty"`
}

// CreationTime returns creation time of the post if present or the time of harvest
func (lexisNexisPayload *LexisNexisPayload) CreationTime() (time.Time, error) {
	parsedTime, err := time.ParseInLocation(time.RFC3339, lexisNexisPayload.HarvestDate, time.UTC)
	if lexisNexisPayload.PublishedDate != "" {
		parsedTime, err = time.ParseInLocation(time.RFC3339, lexisNexisPayload.PublishedDate, time.UTC)
	}
	if err != nil {
		return invalidDate()
	}

	return parsedTime, nil
}
