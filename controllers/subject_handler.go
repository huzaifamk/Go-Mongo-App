package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type SubjectController struct {
	s *mgo.Session
}

func NewSubjectController(s *mgo.Session) *SubjectController {
	return &SubjectController{s}
}