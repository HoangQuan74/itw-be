package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CTGetStudents godoc
// @Summary Get all students
// @Description Get a list of all students
// @Tags students
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Student
// @Failure 500 {object} gin.H
// @Router /students [get]
func CTGetStudents(c *gin.Context) {
	students, err := SVGetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

// CTCreateStudent godoc
// @Summary Create a new student
// @Description Create a new student with the input payload
// @Tags students
// @Accept  json
// @Produce  json
// @Param student body validate.StudentInput true "Student Input"
// @Success 200 {object} entities.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students [post]
func CTCreateStudent(c *gin.Context) {
	var student StudentInput
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Validate.Struct(student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	newStudent, err := SVCreateStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newStudent)
}