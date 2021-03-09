package models

import "gopkg.in/mgo.v2/bson"

type Movie struct {
	ID 					bson.ObjectID 'bson:"_id" json:"id"'
	Name 				string 'bson:"name" json:"name"'
	ThubImage 	string 'bson:"thub_image" json:"thub"'
	Description string        `bson:"description" json:"description"`
	Active      bool          `bson:"active" json:"active"`
}