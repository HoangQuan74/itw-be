package userAccounts

import "github.com/gin-gonic/gin"

func RouterUserAccount(router *gin.RouterGroup) {
	router.GET("/", GetUserAccounts)
}