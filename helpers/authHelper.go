package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err := errors.New("Unauthorized access")
		return err
	}
	return err
}
func MatchID(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err := errors.New("Unauthorized access")
		return err
	}
	err := CheckUserType(c, userType)
	return err
}
