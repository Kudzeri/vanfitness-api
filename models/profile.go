package models

type Profile struct {
	ID     string `bson:"_id,omitempty" json:"id"`
	Prefix string `bson:"prefix" json:"prefix"`
	Level  string `bson:"level" json:"level"`
	Height string `bson:"height" json:"height"`
	Weight string `bson:"weight" json:"weight"`
	Age    string `bson:"age" json:"age"`
	Sex    string `bson:"sex" json:"sex"`

	UserID string `bson:"user_id" json:"user_id"`
}
