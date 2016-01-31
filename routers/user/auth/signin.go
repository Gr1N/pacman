package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/oauth2"
)

type signInCompleteBinding struct {
	State string `form:"state" binding:"required,len=32"`
	Code  string `form:"code" binding:"required,len=20"`
}

// SignIn starts authentication process.
func SignIn(c *gin.Context) {
	service := c.Param("service")
	if err := oauth2.HandleService(service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	redirectURL := oauth2.HandleAuthorizeRequest(service)

	c.JSON(http.StatusOK, gin.H{
		"redirectURL": redirectURL,
	})
}

// SignInComplete finishes authentication process.
func SignInComplete(c *gin.Context) {
	service := c.Param("service")
	if err := oauth2.HandleService(service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var b signInCompleteBinding
	if err := c.Bind(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := oauth2.ValidateAuthorizeRequest(service, b.State); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	user, err := oauth2.FinishAuthorizeRequest(service, b.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	token, _ := models.CreateUserToken(user.ID, "Auth "+time.Now().String())

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"token": token.Value,
		},
	})
}
