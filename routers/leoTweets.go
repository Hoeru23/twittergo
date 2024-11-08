package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RestApi {
	var r models.RestApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil || pag < 1 {
		r.Message = "Debe enviar el enviar parámetro como un valor númerico mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweets(ID, int64(pag))
	if !correcto {
		r.Message = "Error al leer los Tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
