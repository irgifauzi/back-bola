package _bola

import (
	"fmt"
	"testing"
	"time"

	"github.com/irgifauzi/back-bola/model"
	"github.com/irgifauzi/back-bola/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertPemain(t *testing.T) {
	nama_pemain := "Ronaldo"
	tim := "Al-nasr"
	posisi := "Striker"
	tinggi := 187.0
	berat := 83.0
	tanggalLahir := time.Date(1985, time.February, 5, 0, 0, 0, 0, time.UTC)
	tanggalLahirPrimitive := primitive.NewDateTimeFromTime(tanggalLahir)
	negara := "Portugal"
	no_punggung := 7

	insertedID, err := module.InsertPemain(module.MongoConn, "pemain", nama_pemain, tim, posisi, tinggi, berat, tanggalLahirPrimitive, negara, no_punggung)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
		return
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetPemainFromID(t *testing.T) {
	id := "6678f4d89ee2c6d9da780380"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPemainFromID(objectID, module.MongoConn, "pemain")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllDataPemain(module.MongoConn, "pemain")
	fmt.Println(data)
}

func TestDeletePemainByID(t *testing.T) {
	id := "6678f4d89ee2c6d9da780380" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePemainByID(objectID, module.MongoConn, "pemain")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	_, err = module.GetPemainFromID(objectID, module.MongoConn, "pemain")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
