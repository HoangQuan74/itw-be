package routers

import (
	"itw-be/src/modules/student"
	userAccounts "itw-be/src/modules/user-account"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		userAccounts.RouterUserAccount(v1.Group("/user-accounts"))
		student.RouterStudent(v1.Group("/students"))
	}
	return r
}
