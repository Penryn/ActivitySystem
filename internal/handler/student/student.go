package student

import (
	"activitySystem/internal/model"
	"activitySystem/internal/service"
	"activitySystem/pkg/utils"
	"errors"
	"gorm.io/gorm"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateActivityData struct {
	Title     string   `json:"title" binding:"required"`
	Content   string   `json:"content" binding:"required"`
	Location  string   `json:"location" binding:"required"`
	Category  string   `json:"category" binding:"required"`
	Img       []string `json:"img"`
	Deadline  string   `json:"deadline" binding:"required,datetime=2006-01-02T15:04:05Z"`
	StartTime string   `json:"start_time" binding:"required,datetime=2006-01-02T15:04:05Z"`
}

func CreateActivity(c *gin.Context) {
	var data CreateActivityData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user := c.MustGet("user").(*model.User)
	deadline, err := time.Parse(time.RFC3339, data.Deadline)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "时间格式错误")
		return
	}
	startTime, err := time.Parse(time.RFC3339, data.StartTime)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "时间格式错误")
		return
	}
	img := strings.Join(data.Img, ",")
	// 2.创建活动
	err = service.CreateActivity(c, user.ID, data.Title, data.Content, data.Category, data.Location, img, deadline, startTime)
	if err != nil {
		utils.JsonErrorResponse(c, 200508, "创建失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetActivityListData struct {
	Category string `form:"category"`
	PageNum  int    `form:"page_num" binding:"required"`
	PageSize int    `form:"page_size" binding:"required"`
}

type GetActivityListResponse struct {
	ActivityList []model.Activity `json:"Activity_list"`
	Num          int              `json:"num"`
}

func GetNewestActivityList(c *gin.Context) {
	var data GetActivityListData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 业务逻辑
	// 1.获取活动列表
	var activityList []model.Activity
	activityList, num, err := service.GetNewestActivityList(c, data.Category, data.PageNum, data.PageSize)
	if err != nil {
		utils.JsonErrorResponse(c, 200511, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"Activity_list": activityList,
		"num":           math.Ceil(float64(num) / float64(data.PageSize)),
	})
}

func GetLatestActivityList(c *gin.Context) {
	// 业务逻辑
	// 1.获取活动列表
	var activityList []model.Activity
	activityList, _, err := service.GetLatestActivityList(c, 1, 5)
	if err != nil {
		utils.JsonErrorResponse(c, 200511, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"Activity_list": activityList,
	})
}

func GetHottestActivityList(c *gin.Context) {
	// 业务逻辑
	// 1.获取活动列表
	var activityList []model.Activity
	activityList, _, err := service.GetHottestActivityList(c, 1, 10)
	if err != nil {
		utils.JsonErrorResponse(c, 200511, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"Activity_list": activityList,
	})
}

type GetActivityData struct {
	ActivityID int `form:"activity_id" binding:"required"`
}

func GetActivity(c *gin.Context) {
	var data GetActivityData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user := c.MustGet("user").(*model.User)
	// 1.获取活动
	activity, err := service.GetActivityByID(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "活动不存在")
		return
	}
	record, err := service.GetRecordByActivityIDAndUserID(c, user.ID, activity.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.JsonErrorResponse(c, 200512, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"activity": activity,
		"status":   record != nil,
	})

}

type UpdateActivityData struct {
	ActivityID int    `json:"Activity_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Location   string `json:"location" binding:"required"`
	Category   string `json:"category" binding:"required"`
	Deadline   string `json:"deadline" binding:"required,datetime=2006-01-02T15:04:05Z"`
	StartTime  string `json:"start_time" binding:"required,datetime=2006-01-02T15:04:05Z"`
}

func UpdateActivity(c *gin.Context) {
	var data UpdateActivityData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 业务逻辑
	// 1.用户是否存在
	user := c.MustGet("user").(*model.User)

	// 2.活动是否存在
	activity, err := service.GetActivityByID(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "活动不存在")
		return
	}

	// 3.活动是否属于用户
	if activity.UserID != user.ID {
		utils.JsonErrorResponse(c, 200510, "无权限")
		return
	}

	deadline, err := time.Parse(time.RFC3339, data.Deadline)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "时间格式错误")
		return
	}
	startTime, err := time.Parse(time.RFC3339, data.StartTime)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "时间格式错误")
		return
	}
	// 4.更新活动
	err = service.UpdateActivity(c, data.ActivityID, data.Title, data.Content, data.Category, data.Location, deadline, startTime)
	if err != nil {
		utils.JsonErrorResponse(c, 200513, "更新失败")
		return
	}
	utils.JsonSuccess(c, nil)

}

type DeleteActivityData struct {
	ActivityID int `form:"Activity_id" binding:"required"`
}

func DeleteActivity(c *gin.Context) {
	var data DeleteActivityData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 业务逻辑
	// 1.用户是否存在
	user := c.MustGet("user").(*model.User)
	// 2.活动是否存在
	activity, err := service.GetActivityByID(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "活动不存在")
		return
	}
	// 3.活动是否属于用户
	if user.ID != activity.UserID {
		utils.JsonErrorResponse(c, 200510, "无权限")
		return
	}
	// 4.删除活动和更新报名记录
	err = service.DeleteActivityAndRecordByActivityID(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200511, "删除失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type UpvoteActivityData struct {
	ActivityID int `json:"activity_id" binding:"required"`
}

func UpvoteActivity(c *gin.Context) {
	var data UpvoteActivityData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 业务逻辑
	// 2.活动是否存在
	_, err := service.GetActivityByID(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "活动不存在")
		return
	}
	// 3.更新活动
	err = service.UpvoteActivity(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200514, "点赞失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type SignUpActivityData struct {
	ActivityID int `json:"activity_id" binding:"required"`
}

func SignUpActivity(c *gin.Context) {
	var data SignUpActivityData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 业务逻辑
	// 1.用户是否存在
	user := c.MustGet("user").(*model.User)
	// 2.活动是否存在
	Activity, err := service.GetActivityByID(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "活动不存在")
		return
	}
	// 3.报名活动
	err = service.SignUpActivity(c, user.ID, Activity.ID)
	if err != nil {
		utils.JsonErrorResponse(c, 200515, "报名失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type CancelSignUpActivityData struct {
	ActivityID int `json:"activity_id" binding:"required"`
}

func CancelSignUpActivity(c *gin.Context) {
	var data CancelSignUpActivityData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 业务逻辑
	// 1.用户是否存在
	user := c.MustGet("user").(*model.User)
	// 2.活动是否存在
	activity, err := service.GetActivityByID(c, data.ActivityID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "活动不存在")
		return
	}
	// 3.取消报名
	err = service.CancelSignUpActivity(c, user.ID, activity.ID)
	if err != nil {
		utils.JsonErrorResponse(c, 200516, "取消失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetSignUpListData struct {
	PageNum  int `form:"page_num" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

func GetSignUpList(c *gin.Context) {
	var data GetSignUpListData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 业务逻辑
	// 1.用户是否存在
	user := c.MustGet("user").(*model.User)
	// 2.获取报名列表
	RecordList, n, err := service.GetRecordList(c, user.ID, data.PageNum, data.PageSize)
	if err != nil {
		utils.JsonErrorResponse(c, 200517, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"record_list": RecordList,
		"num":         math.Ceil(float64(n) / float64(data.PageSize)),
	})
}

func GetStudentInfo(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	utils.JsonSuccess(c, gin.H{
		"user": user,
	})
}

type UpdateStudentInfoData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	StuID    string `json:"stu_id" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   string `json:"avatar" binding:"required"`
	Profile  string `json:"profile"`
}

func UpdateStudentInfo(c *gin.Context) {
	var data UpdateStudentInfoData
	if err := c.ShouldBindJSON(&data); err != nil {
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
	user := c.MustGet("user").(*model.User)
	u, err := service.GetUserByUsername(c, data.Username)
	if err == nil && u.ID != user.ID {
		utils.JsonErrorResponse(c, 200505, "用户名已存在")
		return
	}
	err = service.UpdateUser(c, user.ID, data.Username, data.Password, data.StuID, data.Email, data.Avatar, data.Profile)
	if err != nil {
		utils.JsonErrorResponse(c, 200518, "更新失败")
		return
	}
	utils.JsonSuccess(c, nil)
}
