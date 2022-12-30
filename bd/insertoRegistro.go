package bd

import (
	"context"
	"fmt"
	"github.com/nmorales1991/go-first-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	//creo un contexto nuevo de 15 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	fmt.Println(ctx)
	defer cancel()

	db := MongoClient.Database("sample_api")
	coll := db.Collection("users")
	// encripto la clave
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Clave), 8)
	u.Clave = string(bytes)
	//inserto
	result, err := coll.InsertOne(ctx, u)
	fmt.Println(result)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//retorno el id del usuario insertado
	return ObjID.String(), true, nil
}
