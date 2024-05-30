package student

import "github.com/gin-gonic/gin"

// RouterStudent sets up the routes for student-related endpoints
func RouterStudent(router *gin.RouterGroup) {
	router.GET("/", CTGetStudents)
	router.POST("/", CTCreateStudent)
}
