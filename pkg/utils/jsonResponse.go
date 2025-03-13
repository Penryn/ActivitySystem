package utils

import "github.com/gin-gonic/gin"

func JsonResponse(c *gin.Context,httpStatus int,code int,msg string,data interface{}){
	c.JSON(httpStatus,gin.H{
		"code":code,
		"msg":msg,
		"data":data,
	})
}

func JsonSuccess(c *gin.Context,data interface{}){
	JsonResponse(c,200,200,"success",data)
}

func JsonErrorResponse(c *gin.Context,code int,msg string){
	JsonResponse(c,200,code,msg,nil)
}