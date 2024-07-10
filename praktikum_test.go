package _714220035

import (
	"fmt"
	"testing"

	"github.com/irgifauzi/back-bola/model"
	"github.com/irgifauzi/back-bola/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertPemain(t *testing.T) {
	nama_pemain := "Ronaldo"
	tim := model.Club{
		Nama_Club: "Barcelona",
		Liga: "Liga eropa",
		Tahun_Berdiri: 1899,
		Stadion: "Olimpiade lluis Companys",
		Manajer: "Hansi Flick",
		Jumlah_Pemain: 11,
		Logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa1EFk-Of9eicV4Bfo-6lwP83tuSoy7eW-BQ&s",
	}
	posisi := "Striker"
	tinggi := 187.0
	berat := 83.0
	tanggal_lahir := "09-06-1985"
	negara := "Portugal"
	no_punggung := 7

	insertedID, err := module.InsertPemain(module.MongoConn, "pemain", nama_pemain, tim, posisi, tinggi, berat, tanggal_lahir, negara, no_punggung)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
		return
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetPemainFromID(t *testing.T) {
	id := "668e488005df6fa1b2719599"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	pema, err := module.GetPemainFromID(objectID, module.MongoConn, "pemain")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(pema)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllDataPemain(module.MongoConn, "pemain")
	fmt.Println(data)
}

func TestDeletePemainByID(t *testing.T) {
	id := "668e488005df6fa1b2719599" 
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePemainByID(objectID, module.MongoConn, "pemain")
	if err != nil {
		t.Fatalf("error calling DeletePemainByID: %v", err)
	}

	_, err = module.GetPemainFromID(objectID, module.MongoConn, "pemain")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}