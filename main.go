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
	uc := controllers.NewSubjectController(getSession())
	e.GET("/subjects", uc.GetAll)
	e.POST("/subjects", uc.Create)
	e.GET("/subjects/:id", uc.GetOne)
	e.PUT("/subjects/:id", uc.Update)
	e.DELETE("/subjects/:id", uc.Delete)
	e.Logger.Fatal(e.Start(":1323"))
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return session
}