package main

    In order to make id private
type User struct {
    ID string `bson:"_id" json:"id"` //this is public
    ID string `bson:"_id" json:"_"` //private
    ID string `bson:"_id" json:"omitempty"` //private
}
