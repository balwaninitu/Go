package domain

import (
	"errors"

	"course-api/src/config"
	"course-api/src/logger"

	"course-api/src/utils"
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
	stmt, err := config.DB.Prepare(queryGetCourse)
	if err != nil {
		logger.ErrorLog.Println("error when trying to prepare get course statement", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(course.Id)

	if getErr := result.Scan(&course.Id, &course.Title); getErr != nil {
		logger.ErrorLog.Println("error when trying to get course", getErr)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}

func (course *Courses) Create() utils.ApiErr {
	stmt, err := config.DB.Prepare(queryCreateCourse)
	if err != nil {
		logger.ErrorLog.Println("error when trying to create course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	//statement get close after operation or when no longer needed
	defer stmt.Close()

	insertResult, createErr := stmt.Exec(course.Title)
	if createErr != nil {
		logger.ErrorLog.Println("error when trying to create course", createErr)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	courseId, err := insertResult.LastInsertId()
	if err != nil {
		logger.ErrorLog.Println("error when trying to create course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	course.Id = courseId
	return nil
}

func (course *Courses) Update() utils.ApiErr {
	stmt, err := config.DB.Prepare(queryUpdateCourse)
	if err != nil {
		logger.ErrorLog.Println("error when trying to update course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Title, course.Id)
	if err != nil {
		logger.ErrorLog.Println("error when trying to update course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}

func (course *Courses) Delete() utils.ApiErr {
	stmt, err := config.DB.Prepare(queryDeleteCourse)
	if err != nil {
		logger.ErrorLog.Println("error when trying to delete course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Id)
	if err != nil {
		logger.ErrorLog.Println("error when trying to delete course", err)
		return utils.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}
