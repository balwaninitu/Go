package services

import (
	"fmt"

	"github.com/balwaninitu/courses_rest_api/src/domain"
	"github.com/balwaninitu/courses_rest_api/src/utils"
)

var (
	CourseService courseServiceInterface = &courseService{}
)

type courseServiceInterface interface {
	Create(course domain.Courses) (*domain.Courses, utils.ApiErr)
	Get(string) (*domain.Courses, utils.ApiErr)
	Update(domain.Courses) (*domain.Courses, utils.ApiErr)
	Delete(string) utils.ApiErr
}

type courseService struct {
}

func (s *courseService) Create(course domain.Courses) (*domain.Courses, utils.ApiErr) {
	if err := course.Create(); err != nil {
		return nil, err
	}
	fmt.Println("from service")
	fmt.Println(course)
	return &course, nil
}

func (s *courseService) Get(id string) (*domain.Courses, utils.ApiErr) {
	course := domain.Courses{Id: id}
	if err := course.Get(); err != nil {
		return nil, err
	}
	return &course, nil
}

func (s *courseService) Update(course domain.Courses) (*domain.Courses, utils.ApiErr) {
	current := &domain.Courses{Id: course.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *courseService) Delete(id string) utils.ApiErr {
	course := &domain.Courses{Id: id}
	return course.Delete()
}
