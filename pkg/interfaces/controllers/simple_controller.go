package controllers

import (
	"net/http"
)

type SimpleController interface {
	Index(c Context)
}

type simpleController struct{}

func NewSimpleController() SimpleController {
	return &simpleController{}
}

func (sc *simpleController) Index(c Context) {
	c.JSON(http.StatusOK, "OK")
}
