package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Soundboard structure contains multiple sounds
type Soundboard struct {
	Title  string
	Sounds []Sound
}

// Sound for Soundboard
type Sound struct {
	Title     string
	Character string
	Episode   string
	File      string
}

func main() {
	fmt.Println("Coucou")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("moncul").C("sounds")
	err = c.Insert(&Sound{"Test", "https://google.com/Prout"},
		&Sound{"Test2", "https://google.com/Bite"})
	if err != nil {
		log.Fatal(err)
	}

	result := Sound{}

	err = c.Find(bson.M{"name": "Test2"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("URL: ", result.URL)
}
