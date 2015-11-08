package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/helpers"
	"github.com/Gr1N/pacman/modules/settings"
)

const (
	sessionIDKey = "session_id"
)

// Init initializes application session.
func Init(g *gin.Engine) {
	store := sessions.NewCookieStore([]byte(settings.S.Secret))
	g.Use(sessions.Sessions("pman_session", store))
}

// Get returns session object (simple shortcut).
func Get(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}

// ID returns unique session ID.
func ID(s sessions.Session) string {
	if id := s.Get(sessionIDKey); id != nil {
		return id.(string)
	}

	id := helpers.RandomString(24)

	s.Set(sessionIDKey, id)
	s.Save()

	return id
}
