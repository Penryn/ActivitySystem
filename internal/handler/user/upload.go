package user

import (
	"activitySystem/internal/global"
	"activitySystem/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path/filepath"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.JsonErrorResponse(c, 400, "获取文件失败")
	}
	// 检查文件大小是否超出限制，限制为50MB
	if file.Size > 50<<20 {
		utils.JsonErrorResponse(c, 400, "文件大小超出限制")
	}
	// 保存文件
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	dst := "./public/" + filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "保存文件失败: "+err.Error())
	}
	urlHost := global.Config.GetString("url.host")
	url := urlHost + "/public/" + filename
	utils.JsonSuccess(c, url)
}
