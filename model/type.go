package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pemain struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Pemain  string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty"`
	Tim          string             `bson:"tim,omitempty" json:"tim,omitempty"`
	Posisi       string             `bson:"posisi,omitempty" json:"posisi,omitempty"`
	Tinggi       float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty"`
	Berat        float64            `bson:"berat,omitempty" json:"berat,omitempty"`
	Tanggal_Lahir primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Negara       string             `bson:"negara,omitempty" json:"negara,omitempty"`
	No_Punggung   int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty"`
}

type Admin struct{
	ID   		primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username	string 		       `bson:"username,omitempty" json:"username,omitempty"`
	Password 	string			   `bson:"password,omitempty" json:"password,omitempty"`
}