package settings

import (
	"github.com/BurntSushi/toml"
	"github.com/imdario/mergo"
)

var (
	// S represents application settings object.
	S *s
)

type s struct {
	RunMode string `toml:"run_mode"`
	Secret  string

	Logger logger
	Server server
	DB     db
	Cache  cache
	Auth   auth
}

type logger struct {
	MinLevel string `toml:"min_level"`
}

type server struct {
	Port string
}

type db struct {
	Driver       string
	Spec         string
	MaxIdleConns int  `toml:"max_idle_conns"`
	MaxOpenConns int  `toml:"max_open_conns"`
	LogMode      bool `toml:"log_mode"`
}

type cache struct {
	Host              string
	Password          string
	DefaultExpiration string `toml:"default_expiration"`
}

type auth struct {
	EnabledServices []string `toml:"enabled_services"`

	GitHubClientID     string   `toml:"github_client_id"`
	GitHubClientSecret string   `toml:"github_client_secret"`
	GitHubRedirectURL  string   `toml:"github_redirect_url"`
	GitHubScopes       []string `toml:"github_scopes"`

	StateCacheTimeout string `toml:"state_cache_timeout"`
}

// Init initializes application settings.
func Init() {
	var defaults s
	if _, err := toml.DecodeFile("conf/app.toml", &defaults); err != nil {
		panic(err)
	}

	var locals s
	if _, err := toml.DecodeFile("conf/app_local.toml", &locals); err == nil {
		mergo.Merge(&locals, defaults)
		S = &locals
	} else {
		S = &defaults
	}
}
