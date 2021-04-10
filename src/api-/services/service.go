package services

import (
	"api-/domain"
	"api-/utils"
)

func Create(course domain.Courses) (*domain.Courses, *utils.ApiErr) {
	return &course, nil
}
