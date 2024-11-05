package routers

import (
	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
)

func BajaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var r models.RestApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es requerido"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar borrar la relación > " + err.Error()
		return r
	}
	if !status {
		r.Message = "No se ha logrado borrar la relación"
		return r
	}

	r.Status = 200
	r.Message = "Baja de relación satisfactoria"
	return r
}
