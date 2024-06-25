package module

import (
	"context"
	"errors"
	"fmt"

	"github.com/irgifauzi/back-bola/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPemain(db *mongo.Database, col string, nama_pemain string, tim string, posisi string, tinggi float64, berat float64, tanggal_lahir primitive.DateTime, negara string, no_punggung int) (insertedID primitive.ObjectID, err error) {
	pemain := model.Pemain{
		Nama_Pemain:   nama_pemain,
		Tim:           tim,
		Posisi:        posisi,
		Tinggi:        tinggi,
		Berat:         berat,
		Tanggal_Lahir: tanggal_lahir,
		Negara:        negara,
		No_Punggung:   no_punggung,
	}

	result, err := db.Collection(col).InsertOne(context.Background(), pemain)
	if err != nil {
		fmt.Printf("InsertPemain: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllDataPemain(db *mongo.Database, col string) (data []model.Pemain) {
	gem := db.Collection(col)
	filter := bson.M{}
	cursor, err := gem.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllDataPemain: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetPemainFromID(_id primitive.ObjectID, db *mongo.Database, col string) (pemain model.Pemain, errs error) {
	bols := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := bols.FindOne(context.TODO(), filter).Decode(&pemain)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return pemain, fmt.Errorf("no data found for ID %s", _id)
		}
		return pemain, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return pemain, nil
}

func UpdatePemain(db *mongo.Database, col string, id primitive.ObjectID, nama_pemain string, tim string, posisi string, tinggi float64, berat float64, tanggal_lahir primitive.DateTime, negara string, no_punggung int) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_pemain":   nama_pemain,
			"tim":           tim,
			"posisi":        posisi,
			"tinggi":        tinggi,
			"berat":         berat,
			"tanggal_lahir": tanggal_lahir,
			"negara":        negara,
			"no_punggung":   no_punggung,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePemain: %v\n", err)
		return err
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return err
	}
	return nil
}

func DeletePemainByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := karyawan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
