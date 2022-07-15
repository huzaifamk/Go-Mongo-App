package models

import (
	"gopkg.in/mgo.v2"
)

type Subject struct {
	ID   int `bson:"_id"`
	Name string `bson:"name"`
	Board string `bson:"board"`
	Paper string `bson:"paper"`
}