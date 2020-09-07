package middleware

import (
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("goshop_member_id", "1")
	}
}
