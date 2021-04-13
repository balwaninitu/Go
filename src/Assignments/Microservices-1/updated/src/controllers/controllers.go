package controllers

import (
	"net/http"
	"strconv"

	"courses_api/src/domain"

	"courses_api/src/services"
	"courses_api/src/utils"

	"github.com/gin-gonic/gin"
)

/*checkCourseId will check for primitive type of id and provide error
if incorrect value enter by user*/
func checkCourseId(courseIdParam string) (int64, utils.ApiErr) {
	courseId, useErr := strconv.ParseInt(courseIdParam, 10, 64)
	if useErr != nil {
		return 0, utils.NewBadRequestError("course id should be a number")
	}
	return courseId, nil
}

//controller package will be first or intenal layer of mvc pattern
//all request get handle by controllers through end points
/*when func create get invoke from handler it will map json data provided by user
from body,display error in case of incorrect json data*/
func Create(c *gin.Context) {
	var course domain.Courses
	if err := c.ShouldBindJSON(&course); err != nil {
		createErr := utils.NewBadRequestError("invalid json body")
		c.JSON(createErr.Status(), createErr)
		return
	}
	/*saving in database,controller is no incharge of databse, it all take care by services
	if json data in body isvalid create func in services package get invoke and if err nil
	data will be added in persistent storage which is mysql database*/
	result, saveErr := services.Create(course)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

/*when Get func invoke from handler it will check input type of id by calling func
checkcourseId and error if not correct */
func Get(c *gin.Context) {
	courseId, idErr := checkCourseId(c.Param("course_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	//if input id is okay get func from services package get called
	course, getErr := services.Get(courseId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, course)
}

func Update(c *gin.Context) {
	courseId, idErr := checkCourseId(c.Param("course_id"))
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
	courseId, idErr := checkCourseId(c.Param("course_id"))
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
