package session

import (
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/helpers"
	"github.com/Gr1N/pacman/modules/settings"
)

const (
	sessionIDKey = "session_id"
	userIDKey    = "user_id"
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

// UserID returns User ID saved in session.
func UserID(s sessions.Session) (int64, bool) {
	if rawID := s.Get(userIDKey); rawID != nil {
		if userID, err := strconv.ParseInt(rawID.(string), 10, 64); err == nil {
			return userID, true
		}
	}

	return -1, false
}

// SetUserID saves User ID in session.
func SetUserID(s sessions.Session, userID int64) {
	s.Set(userIDKey, strconv.FormatInt(userID, 10))
	s.Save()
}
