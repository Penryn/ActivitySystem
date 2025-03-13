package user

import (
	"activitySystem/internal/service"
	"activitySystem/pkg/jwt"
	"activitySystem/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 业务逻辑
	// 1.用户名是否存在
	user, err := service.GetUserByUsername(c, data.Username)
	if err != nil {
		utils.JsonErrorResponse(c, 200504, "用户不存在")
		return
	}
	// 2.密码是否正确
	if user.Password != data.Password {
		utils.JsonErrorResponse(c, 200512, "密码错误")
		return
	}

	token := jwt.NewJWT(user.ID)
	utils.JsonSuccess(c, gin.H{
		"token": token,
	})
}
