package middleware

import (
	"activitySystem/internal/service"
	"activitySystem/pkg/jwt"
	"activitySystem/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.JsonErrorResponse(c, http.StatusUnauthorized, "token无效")
			c.Abort()
			return
		}
		uid, err := jwt.ParseJWT(token)
		if err != nil {
			utils.JsonErrorResponse(c, 200500, "token无效")
			c.Abort()
			return
		}
		user, err := service.GetUserByID(c, uid)
		if err != nil {
			utils.JsonErrorResponse(c, 200507, "用户不存在")
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
