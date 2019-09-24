package models

// NecessaryEnvVars probably should go away at some point
var NecessaryEnvVars = map[string][]string{
	"google_plus":    {"GOOGLE_PLUS_AUTH_TOKEN_CLIENT_ID", "GOOGLE_PLUS_AUTH_TOKEN_CLIENT_SECRET"},
	"instagram":      {},
	"reddit":         {},
	"stack_overflow": {},
	"vk":             {},
	"youtube":        {"YOUTUBE_AUTH_TOKEN_CLIENT_ID", "YOUTUBE_AUTH_TOKEN_CLIENT_SECRET"},
	"pinterest":      {},
	"pinterest_v1":   {},
	"instagram_v2":   {},
	"facebook":       {"TOKEN_APP_URL", "FACEBOOK_AUTH_TOKEN_CLIENT_ID", "FACEBOOK_AUTH_TOKEN_CLIENT_SECRET"},
	"twitter":        {},
	"lexis_nexis":    {"LEXIS_NEXIS_KEY"},
}
