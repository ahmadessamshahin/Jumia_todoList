package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Matcher struct {
	RoutingGroups RoutingGroup
	UseCases      interface{}
}

type DefaultHandler struct {
	Logger  *zerolog.Logger
	Router  *gin.Engine
	Matcher []Matcher
}

func NewDefaultHandler(Logger *zerolog.Logger) *DefaultHandler {
	return &DefaultHandler{
		Logger:  Logger,
		Router:  gin.Default(),
		Matcher: make([]Matcher, 0),
	}
}

func (s *DefaultHandler) Use(gp RoutingGroup, useCase interface{}) *DefaultHandler {
	newMatcher := Matcher{RoutingGroups: gp, UseCases: useCase}
	s.Matcher = append(s.Matcher, newMatcher)
	return s
}

func (s DefaultHandler) Serve() {
	
	for _, m := range s.Matcher {
		m.RoutingGroups.Routes(s.Router, m.UseCases)
	}
}
