package deps

import (
	"regexp"
)

const (
	depsLocations = []string{
		// "",
		"requirements",
	}
	depsPatterns = []*regexp.Regexp{
		regexp.MustCompile("^requiremets\\/.+\\.(?:txt|pip)$"),
	}
)
