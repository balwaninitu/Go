package services

import (
	"ms/domain"
	"ms/utils"
)

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

// func GetKey(course domain.Courses) (*domain.Courses, utils.ApiErr) {
// 	accessKey := &domain.Courses{
// 		Key: course.Key,
// 	}
// 	if err := accessKey.GetKey(); err != nil {
// 		return nil, err
// 	}
// 	return accessKey, nil
// }
