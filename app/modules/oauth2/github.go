package oauth2

var (
	GitHub = &Config{
		ClientID:     "fake",
		ClientSecret: "fake",
		RedirectUrl:  "http://localhost",
		Scopes:       []string{},

		Endpoint: Endpoint{
			AuthUrl:  "https://github.com/login/oauth/authorize",
			TokenUrl: "https://github.com/login/oauth/access_token",
		},
	}
)
