package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/balwaninitu/courses_rest_api/src/domain"
	"github.com/balwaninitu/courses_rest_api/src/services"
	"github.com/balwaninitu/courses_rest_api/src/utils"
	"github.com/gorilla/mux"
)

var (
	CourseController courseControllerInterface = &coursesController{}
)

type courseControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type coursesController struct {
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

func (c *coursesController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := utils.NewBadRequestError("invalid request body")
		utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var courseRequest domain.Courses
	if err := json.Unmarshal(requestBody, &courseRequest); err != nil {
		respErr := utils.NewBadRequestError("invalid json body")
		utils.RespondError(w, respErr)
		return
	}

	result, createErr := services.CourseService.Create(courseRequest)
	if createErr != nil {
		utils.RespondError(w, createErr)
		return
	}
	utils.RespondJson(w, http.StatusCreated, result)
	fmt.Println("from controllers")
	fmt.Println(result)

}
func (c *coursesController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	courseId := strings.TrimSpace(params["id"])

	course, err := services.CourseService.Get(courseId)
	if err != nil {
		utils.RespondJson(w, http.StatusOK, course)
	}

}

func (c *coursesController) Update(w http.ResponseWriter, r *http.Request) {
	var course domain.Courses
	params := mux.Vars(r)

	course.Id = strings.TrimSpace(params["id"])
	//isPartial := c.Request.Method == http.MethodPut

	result, err := services.CourseService.Update(course)
	if err != nil {
		utils.RespondJson(w, http.StatusOK, result)
	}

}

func (c *coursesController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	courseId := strings.TrimSpace(params["id"])

	if err := services.CourseService.Delete(courseId); err != nil {
		utils.RespondJson(w, http.StatusOK, map[string]string{"status": "Deleted"})
	}

}
