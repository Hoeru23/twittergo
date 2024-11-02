package bd

import (
	"context"

	"github.com/Hoeru23/twittergo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	collection := db.Collection("usuarios") // Colecciones son similares a tablas

	condition := bson.M{"email": email} // funciona como un where/find

	var resultado models.Usuario
	err := collection.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex() // Convierte el hex en string
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
