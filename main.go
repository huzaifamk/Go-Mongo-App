package main

import (
	"fmt"

	"github.com/huzaifamk/Go-Mongo-App/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("Go-Mongo-App")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	SubjectController := controllers.NewSubjectController(getSession())
	e.GET("/subjects/:id", SubjectController.GetOne)
	e.GET("/subjects", SubjectController.GetAll)
	e.POST("/subjects", SubjectController.Create)
	e.PUT("/subjects/:id", SubjectController.Update)
	e.DELETE("/subjects/:id", SubjectController.Delete)
	e.POST("/postform", SubjectController.PostForm)
	e.POST(".postfile", SubjectController.PostFile)
	e.Logger.Fatal(e.Start(":1323"))
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return session
}
