package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"ms/domain"

	"ms/services"
	"ms/utils"

	"github.com/gin-gonic/gin"
)

func getCourseId(courseIdParam string) (int64, utils.ApiErr) {
	courseId, useErr := strconv.ParseInt(courseIdParam, 10, 64)
	if useErr != nil {
		return 0, utils.NewBadRequestError("course id should be a number")
	}
	return courseId, nil
}

func GetByKey(accessTokenId string) (string, error) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return "", utils.NewBadRequestError("invalid access token id")
	}
	return "", nil
}

// func GetToken() {
// 	accessToken, err := GetByKey(c.Param("access_token_id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, accessToken)
// }
// func GetKey(c *gin.Context) {
// 	var course domain.Courses
// 	if err := c.ShouldBindJSON(&course); err != nil {
// 		getErr := utils.NewBadRequestError("invalid json body")
// 		c.JSON(getErr.Status(), getErr)
// 		return
// 	}
// 	result, err := services.GetKey(course)
// 	if err != nil {
// 		c.JSON(err.Status(), err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, result)
// }

func Create(c *gin.Context) {
	var course domain.Courses
	if err := c.ShouldBindJSON(&course); err != nil {
		createErr := utils.NewBadRequestError("invalid json body")
		c.JSON(createErr.Status(), createErr)
		return
	}
	result, saveErr := services.Create(course)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func Get(c *gin.Context) {
	// str, err := GetByKey(c.Param("access_token_id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err)
	// 	return
	// }
	//c.JSON(http.StatusOK, str)
	courseId, idErr := getCourseId(c.Param("course_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	course, getErr := services.Get(courseId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, course)
}

func Update(c *gin.Context) {
	courseId, idErr := getCourseId(c.Param("course_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var course domain.Courses
	if err := c.ShouldBind(&course); err != nil {
		upErr := utils.NewBadRequestError("invalid json body")
		c.JSON(upErr.Status(), upErr)
		return
	}
	course.Id = courseId

	result, err := services.Update(course)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	courseId, idErr := getCourseId(c.Param("course_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	if err := services.Delete(courseId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
