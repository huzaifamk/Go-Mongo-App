package controllers

import (
	"net/http"

	models "github.com/huzaifamk/Go-Mongo-App/models"
	helpers "github.com/huzaifamk/Go-Mongo-App/helper"
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

func (c *SubjectController) PostForm(e echo.Context) error {
	name := e.FormValue("name")
	session := c.s.Copy()
	defer session.Close()
	collection := session.DB("test").C("info")
	err := collection.Insert(bson.M{"name": name})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func (c *SubjectController) UploadFile(e echo.Context) error {
	file, err := e.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	result, err := helpers.UploadToS3(e, file.Filename, src)
	if err != nil {
		return err
	}
	data := &models.UploadResult{
		Path: result,
	}
	return e.JSON(http.StatusOK, data)
}

