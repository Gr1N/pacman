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

	Logger logger
	Server server
	DB     db
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
