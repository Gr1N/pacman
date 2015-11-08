package oauth2

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
