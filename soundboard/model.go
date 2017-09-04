package soundboard

import (
	"gopkg.in/mgo.v2/bson"
)

type Soundboard struct {
	ID     bson.ObjectId `bson:"_id"`
	Title  string        `json:"title"`
	Sounds []Sound       `json:"sounds"`
}

type Sound struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string        `json:"title"`
	Character string        `json:"character"`
	Episode   string        `json:"episode"`
	File      string        `json:"file"`
}

type Sounds []Sound