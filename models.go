package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type Soundboard struct {
	Title  string
	Sounds []Sound
}

type Sound struct {
	Title     string
	Character string
	Episode   string
	File      string
}