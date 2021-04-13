package domain

import (
	"errors"

	"courses_api/src/config"

	"courses_api/src/utils"

	"github.com/balwaninitu/courses_rest_api/src/logger"
)

const (
	queryCreateCourse = "INSERT INTO courses(title) VALUES(?);"
	queryGetCourse    = "SELECT id, title FROM courses WHERE id=?;"
	queryUpdateCourse = "UPDATE courses SET title=? WHERE id=?;"
	queryDeleteCourse = "DELETE FROM courses WHERE id=?;"
)

//transferring data from persistent layer(database) to application
type Courses struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func (course *Courses) Get() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryGetCourse)
	if err != nil {
		logger.Error("error when trying to prepare get course statement", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(course.Id)

	if getErr := result.Scan(&course.Id, &course.Title); getErr != nil {
		logger.Error("error when trying to get course", getErr)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}

func (course *Courses) Create() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryCreateCourse)
	if err != nil {
		logger.Error("error when trying to create course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	//statement get close after operation or when no longer needed
	defer stmt.Close()

	insertResult, createErr := stmt.Exec(course.Title)
	if createErr != nil {
		logger.Error("error when trying to create course", createErr)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	courseId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to create course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	course.Id = courseId
	return nil
}

func (course *Courses) Update() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryUpdateCourse)
	if err != nil {
		logger.Error("error when trying to update course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Title, course.Id)
	if err != nil {
		logger.Error("error when trying to update course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}

func (course *Courses) Delete() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryDeleteCourse)
	if err != nil {
		logger.Error("error when trying to delete course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Id)
	if err != nil {
		logger.Error("error when trying to delete course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}
