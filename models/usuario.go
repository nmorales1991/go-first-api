package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuario struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre    string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos string             `bson:"apellidos,omitempty" json:"apellidos,omitempty"`
	Clave     string             `bson:"clave,omitempty" json:"clave,omitempty"`
	Correo    string             `bson:"correo,omitempty" json:"correo,omitempty"`
}
