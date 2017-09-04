package soundboard

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type Repository struct{}

const SERVER = "localhost:27017"

const DBNAME = "soundboard"
const DOCNAME = "sounds"

func (r Repository) GetSounds() Sounds {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to MongoDB server: ", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Sounds{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

func (r Repository) AddSound(sound Sound) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	sound.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(sound)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) UpdateSound(sound Sound) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	sound.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(sound.ID, sound)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) DeleteSound(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify ID
	if !bson.IsObjectIdHex(id) {
		return "Not Found"
	}

	oid := bson.ObjectIdHex(id)

	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "Internal Error"
	}

	return "OK"
}
