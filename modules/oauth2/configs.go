package oauth2

// Config describes a typical 3-legged OAuth2 flow, with both the
// client application information and the server's endpoint URLs.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string

	Endpoint Endpoint
}

// Endpoint contains the OAuth 2.0 provider's authorization, token and user
// endpoint URLs.
type Endpoint struct {
	AuthURL  string
	TokenURL string
	UserURL  string
}

// NewGitHub initializes new config for GitHub service.
func NewGitHub(clientID, clientSecret, redirectURL string,
	scopes []string) *Config {

	return &Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,

		Endpoint: Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
			UserURL:  "https://api.github.com/user",
		},
	}
}
