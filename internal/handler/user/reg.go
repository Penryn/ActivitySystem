package user

import (
	"activitySystem/internal/service"
	"activitySystem/pkg/utils"
	"github.com/gin-gonic/gin"
)

type RegDate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	StuID    string `json:"stu_id" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func Reg(c *gin.Context) {
	var data RegDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 参数校验
	//用户名只有数字
	if !utils.IsNumber(data.StuID) {
		utils.JsonErrorResponse(c, 200502, "学号必须为纯数字")
		return
	}
	//学号长度
	if len(data.StuID) != 12 {
		utils.JsonErrorResponse(c, 200503, "学号长度必须12位")
		return
	}

	// 业务逻辑
	// 1.用户名是否存在
	_, err = service.GetUserByUsername(c, data.Username)
	if err == nil {
		utils.JsonErrorResponse(c, 200505, "用户名已存在")
		return
	}

	// 2.创建用户
	err = service.CreateUser(c, data.Username, data.Password, data.StuID, data.Email)
	if err != nil {
		utils.JsonErrorResponse(c, 200506, "注册失败")
		return
	}

	utils.JsonSuccess(c, nil)

}
