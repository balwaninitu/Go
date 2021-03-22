package controllers

import (
	"Api/utils"
	"net/http"
)

type Controller struct{}

func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseJSON(w, "Yes")
	}
}
