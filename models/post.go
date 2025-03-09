package models

type Post struct {
	ID	  string `bson:"_id,omitempty" json:"id"`
	Title string `bson:"title" json:"title"`
	Body  string `bson:"body" json:"body"`
	UserID string `bson:"user_id" json:"user_id"`
}