package domain

import (
	"api-/config"
	"api-/utils"
)

var (
	courseDB = make(map[int64]*Courses)
)

type Courses struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func (course Courses) Get() *utils.ApiErr {
	if err := config.Client.Ping(); err != nil {
		panic(err)
	}

	result := courseDB[course.Id]
	if result == nil {

	}
	course.Id = result.Id
	course.Title = result.Title
	return nil
}
