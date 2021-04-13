package services

import (
	"courses_api/src/domain"
	"courses_api/src/utils"
)

//all business logic are in services package
//Id input by user will get sent over if available in database once get controller get hit
func Get(courseId int64) (*domain.Courses, utils.ApiErr) {
	result := &domain.Courses{Id: courseId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func Create(course domain.Courses) (*domain.Courses, utils.ApiErr) {
	if err := course.Create(); err != nil {
		return nil, err
	}
	return &course, nil
}

func Update(course domain.Courses) (*domain.Courses, utils.ApiErr) {
	current := &domain.Courses{Id: course.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	current.Id = course.Id
	current.Title = course.Title
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func Delete(courseId int64) utils.ApiErr {
	course := &domain.Courses{Id: courseId}
	return course.Delete()
}
