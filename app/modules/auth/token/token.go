package token

import (
	"errors"

	"github.com/pborman/uuid"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/models"
	"github.com/Gr1N/pacman/app/modules/helpers"
)

const (
	audienceMaxLength = 255
	audienceMinLength = 1
)

var (
	ErrAudienceRequired = errors.New("Audience does not match requirements")
)

func ValidateToken() {

}

func ValidateTokenRequest(audience string, v *revel.Validation) error {
	v.Required(audience)
	v.MaxSize(audience, audienceMaxLength)
	v.MinSize(audience, audienceMinLength)

	if v.HasErrors() {
		return ErrAudienceRequired
	}

	return nil
}

func FinishTokenRequest(userId int64, audience string) *models.Token {
	value := uuid.NewRandom().String()
	value = helpers.EncodeSha1(value)

	token, _ := models.CreateUserToken(userId, audience, value)
	return token
}
