package routers

import (
	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var r models.RestApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	err := bd.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrió un error al intentar eliminar el tweet > " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Tweet eliminado correctamente."
	return r
}
