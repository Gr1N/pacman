package settings

import (
	"github.com/BurntSushi/toml"
	"github.com/imdario/mergo"
)

var (
	S *s
)

type s struct {
	RunMode string `toml:"run_mode"`

	Server server
}

type server struct {
	Port string
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
