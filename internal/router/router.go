package router

import (
	s "activitySystem/internal/handler/student"
	u "activitySystem/internal/handler/user"
	"activitySystem/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "api"

	api := r.Group(pre)
	{
		user := api.Group("/user")
		{
			user.POST("/reg", u.Reg)
			user.POST("/login", u.Login)
			user.POST("/upload", u.Upload)
		}
		student := api.Group("/student").Use(middleware.JwtMid())
		{
			student.GET("/info", s.GetStudentInfo)
			student.PUT("/info", s.UpdateStudentInfo)
			student.POST("/activity", s.CreateActivity)
			student.GET("/activities", s.GetNewestActivityList)
			student.GET("/activity/newest", s.GetLatestActivityList)
			student.GET("/activity/hottest", s.GetHottestActivityList)
			student.GET("/activity", s.GetActivity)
			student.PUT("/activity", s.UpdateActivity)
			student.DELETE("/activity", s.DeleteActivity)
			student.POST("/activity/upvote", s.UpvoteActivity)
			student.POST("/activity/signUp", s.SignUpActivity)
			student.POST("/activity/signUp/cancel", s.CancelSignUpActivity)
			student.GET("/activity/signUp", s.GetSignUpList)
		}
	}
}
