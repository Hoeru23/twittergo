package routers

import (
	"encoding/json"
	"fmt"

	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RestApi {
	var r models.RestApi
	r.Status = 400

	fmt.Println("Entré en VerPerfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Ocurrió un error al intentar buscar el registro " + err.Error()
		return r
	}

	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Ocurrió un error al formatear los datos de los usuarios como JSON " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
