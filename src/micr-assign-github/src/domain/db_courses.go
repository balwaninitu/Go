package domain

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/balwaninitu/courses_rest_api/src/config"
	"github.com/balwaninitu/courses_rest_api/src/utils"
)

const (
	queryGetCourse    = "SELECT id, title FROM gocourses WHERE id=?;"
	queryCreateCourse = "INSERT INTO gocourses(id, title) VALUES(? , ?);"
	queryUpdateCourse = "UPDATE gocourses SET title=? WHERE Id=?;"
	queryDeleteCourse = "DELETE FROM gocourses WHERE id=?;"
)

type Courses struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func (course *Courses) Get() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryGetCourse)
	if err != nil {
		return utils.NewInternalServerError("error when trying to get course", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(course.Id)

	if getErr := result.Scan(&course.Id, &course.Title); getErr != nil {
		return utils.NewInternalServerError("error when trying to get course", errors.New("database error"))
	}
	return nil
}

func (course *Courses) Create() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryCreateCourse)
	if err != nil {
		return utils.NewInternalServerError("error when trying to create course", errors.New("database error"))
	}
	defer stmt.Close()

	insertCourse, createErr := stmt.Exec(course.Title)
	if createErr != nil {
		return utils.NewInternalServerError("error when trying to create course", errors.New("database error"))
	}

	courseId, err := insertCourse.LastInsertId()
	if err != nil {
		return utils.NewInternalServerError("error when trying to create course", errors.New("database error"))
	}
	course.Id = strconv.FormatInt(courseId, 10)
	fmt.Println("from db_courses")
	fmt.Println(course.Id)
	fmt.Println(courseId)
	return nil
}

func (course *Courses) Update() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryUpdateCourse)
	if err != nil {
		return utils.NewInternalServerError("error when trying to update course", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Title, course.Id)
	if err != nil {
		return utils.NewInternalServerError("error when trying to update course", errors.New("database error"))
	}
	return nil
}

func (course *Courses) Delete() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryDeleteCourse)
	if err != nil {
		return utils.NewInternalServerError("error when trying to delete course", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Title, course.Id)
	if err != nil {
		return utils.NewInternalServerError("error when trying to update course", errors.New("database error"))
	}
	return nil
}
