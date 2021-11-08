package handler

import (
	"Jumia_todoList/usecase"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Matcher struct {
	RoutingGroups RoutingGroup
	Method        string
}

type DefaultHandler struct {
	ServiceBuilder usecase.DefaultService
	Logger         *zerolog.Logger
	Router         *gin.Engine
	Matcher        []Matcher
}

func NewDefaultHandler(Logger *zerolog.Logger, serviceBuilder usecase.DefaultService) *DefaultHandler {
	return &DefaultHandler{
		ServiceBuilder: serviceBuilder,
		Logger:         Logger,
		Router:         gin.Default(),
		Matcher:        make([]Matcher, 0),
	}
}

func (s *DefaultHandler) Use(method string, gp RoutingGroup) *DefaultHandler {
	newMatcher := Matcher{RoutingGroups: gp, Method: method}
	s.Matcher = append(s.Matcher, newMatcher)
	return s
}

func (s DefaultHandler) Serve() {

	for _, m := range s.Matcher {
		m.RoutingGroups.Routes(s.Router, s.ServiceBuilder.Resolve(m.Method))
	}
}
