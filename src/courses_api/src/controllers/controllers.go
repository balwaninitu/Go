package controllers

import (
	"net/http"
	"os"
	"strconv"

	"courses_api/src/domain"

	"courses_api/src/services"
	"courses_api/src/utils"

	"github.com/gin-gonic/gin"
)

func getCourseId(courseIdParam string) (int64, utils.ApiErr) {
	courseId, courseErr := strconv.ParseInt(courseIdParam, 10, 64)
	if courseErr != nil {
		return 0, utils.NewBadRequestError("course id should be a number")
	}
	return courseId, nil
}

func validKey(w http.ResponseWriter, r *http.Request) bool {
	v := r.URL.Query()
	secret := os.Getenv("SECRET")
	if key, ok := v["key"]; ok {
		if key[0] == secret {
			return true
		} else { //invalid key
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("401 - Invalid key"))
			return false
		}
	} else { //key is not provided
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Please supply access key"))
		return false
	}
}

func Create(c *gin.Context) {
	var course domain.Courses
	if err := c.ShouldBindJSON(&course); err != nil {
		createErr := utils.NewBadRequestError("invalid json body")
		c.JSON(createErr.Status(), createErr)
		return
	}
	result, saveErr := services.CourseService.Create(course)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

// func (c *coursesController) Get(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	courseId := strings.TrimSpace(params["id"])

// 	course, err := services.CourseService.Get(courseId)
// 	if err != nil {
// 		utils.RespondJson(w, http.StatusOK, course)
// 	}

// }

// func (c *coursesController) Update(w http.ResponseWriter, r *http.Request) {
// 	var course domain.Courses
// 	params := mux.Vars(r)

// 	course.Id = strings.TrimSpace(params["id"])
// 	//isPartial := c.Request.Method == http.MethodPut

// 	result, err := services.CourseService.Update(course)
// 	if err != nil {
// 		utils.RespondJson(w, http.StatusOK, result)
// 	}

// }

// func (c *coursesController) Delete(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	courseId := strings.TrimSpace(params["id"])

// 	if err := services.CourseService.Delete(courseId); err != nil {
// 		utils.RespondJson(w, http.StatusOK, map[string]string{"status": "Deleted"})
// 	}

// }
