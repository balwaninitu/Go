package domain

import (
	"errors"

	"courses_api/src/config"

	"courses_api/src/utils"
)

const (
	queryCreateCourse = "INSERT INTO courses(key, title) VALUES(?, ?);"
	queryGetCourse    = "SELECT id, key, title FROM courses WHERE id=?;"
	queryUpdateCourse = "UPDATE courses SET key=?, title=? WHERE id=?;"
	queryDeleteCourse = "DELETE FROM courses WHERE id=?;"
)

type Courses struct {
	Id    int64  `json:"id"`
	Key   string `json:"key"`
	Title string `json:"title"`
}

func (course *Courses) Get() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryGetCourse)
	if err != nil {
		return utils.NewInternalServerError("error when trying to get course", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(course.Id)

	if getErr := result.Scan(&course.Id, &course.Key, &course.Title); getErr != nil {
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

	//result, err := config.Client.Exec(queryCreateCourse,course.Title )

	insertResult, createErr := stmt.Exec(course.Key, course.Title)
	if createErr != nil {
		return utils.NewInternalServerError("error when trying to create course", errors.New("database error"))
	}
	courseId, err := insertResult.LastInsertId()
	if err != nil {
		return utils.NewInternalServerError("error when trying to create course", errors.New("database error"))
	}
	course.Id = courseId
	return nil
}

func (course *Courses) Update() utils.ApiErr {
	stmt, err := config.Client.Prepare(queryUpdateCourse)
	if err != nil {
		return utils.NewInternalServerError("error when trying to update course", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Key, course.Title, course.Id)
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

	_, err = stmt.Exec(course.Id)
	if err != nil {
		return utils.NewInternalServerError("error when trying to delete course", errors.New("database error"))
	}
	return nil
}
