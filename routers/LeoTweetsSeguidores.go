package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
)

func LeoTweetsSeguidores(request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var r models.RestApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()
	pagina := request.QueryStringParameters["pagina"]

	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil || pag < 1 {
		r.Message = "Debe enviar el enviar parámetro como un valor númerico mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweetsSeguidores(IDUsuario, pag)
	if !correcto {
		r.Message = "Error al leer los tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de tweets de los seguidores"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
