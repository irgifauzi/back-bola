package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Club struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Club     string             `bson:"nama_club,omitempty" json:"nama_club,omitempty"`
	Liga          string             `bson:"liga,omitempty" json:"liga,omitempty"`
	Tahun_Berdiri int                `bson:"tahun_berdiri,omitempty" json:"tahun_berdiri,omitempty"`
	Stadion       string             `bson:"stadion,omitempty" json:"stadion,omitempty"`
	Manajer       string             `bson:"manajer,omitempty" json:"manajer,omitempty"`
	Jumlah_Pemain int                `bson:"jumlah_pemain,omitempty" json:"jumlah_pemain,omitempty"`
	Logo          string             `bson:"logo,omitempty" json:"logo,omitempty"`
}

type Pemain struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Pemain   string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty"`
	Tim           Club               `bson:"tim,omitempty" json:"tim,omitempty"` 
	Posisi        string             `bson:"posisi,omitempty" json:"posisi,omitempty"`
	Tinggi        float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty"`
	Berat         float64            `bson:"berat,omitempty" json:"berat,omitempty"`
	Tanggal_Lahir string 			 `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Negara        string             `bson:"negara,omitempty" json:"negara,omitempty"`
	No_Punggung   int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty"`
}


type Admin struct{
	ID   		primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username	string 		       `bson:"username,omitempty" json:"username,omitempty"`
	Password 	string			   `bson:"password,omitempty" json:"password,omitempty"`
}