package action

import (
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {

}

func PublishList(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	userID := c.DefaultQuery("user_id", "")
	if VerifyToken(token, userID) {
		// resp, err := domain.PublishList()
	}

}

func VerifyToken(token string, userID string) bool {
	return true
}
