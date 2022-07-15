package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Subject struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Board string `bson:"board" json:"board"`
	Paper string `bson:"paper" json:"paper"`
}