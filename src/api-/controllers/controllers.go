package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// func Create(c *gin.Context) {
// 	var course domain.Courses
// 	if err := c.ShouldBindJSON(&course); err != nil {
// 		return
// 	}
// 	// result, createErr := services.Create(course)
// 	// if createErr != nil {
// 	// 	return
// 	// }
// 	c.JSON(http.StatusCreated, result)
// }
func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
