package controllers

import (
	"net/http"

	models "github.com/huzaifamk/Go-Mongo-App/models"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SubjectController struct {
	s *mgo.Session
}

func NewSubjectController(s *mgo.Session) *SubjectController {
	return &SubjectController{s}
}

func (c *SubjectController) GetOne(e echo.Context) error {
	id := e.Param("id")
	session := c.s.Copy()
	defer session.Close()
	collection := session.DB("test").C("subjects")
	var subject models.Subject
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&subject)
	if err != nil {
		return e.JSON(http.StatusNotFound, map[string]string{"result": "not found"})
	}
	return e.JSON(http.StatusOK, subject)
}

func (c *SubjectController) GetAll(e echo.Context) error {
	session := c.s.Copy()
	defer session.Close()
	collection := session.DB("test").C("subjects")
	var subjects []models.Subject
	err := collection.Find(nil).All(&subjects)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, subjects)
}

func (c *SubjectController) Create(e echo.Context) error {
	var subject models.Subject
	subject.ID = bson.NewObjectId()
	err := e.Bind(&subject)
	if err != nil {
		return err
	}
	session := c.s.Copy()
	defer session.Close()
	collection := session.DB("test").C("subjects")
	err = collection.Insert(subject)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusCreated, subject)
}

func (c *SubjectController) Update(e echo.Context) error {
	id := e.Param("id")
	var subject models.Subject
	subject.ID = bson.ObjectIdHex(id)
	err := e.Bind(&subject)
	if err != nil {
		return err
	}
	session := c.s.Copy()
	defer session.Close()
	collection := session.DB("test").C("subjects")
	err = collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &subject)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, subject)
}

func (c *SubjectController) Delete(e echo.Context) error {
	id := e.Param("id")
	session := c.s.Copy()
	defer session.Close()
	collection := session.DB("test").C("subjects")
	err := collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, map[string]string{"result": "success"})
}

