package userAccounts

import "github.com/gin-gonic/gin"

func GetUserAccounts(c *gin.Context) {
	c.JSON(200, nil)
}
