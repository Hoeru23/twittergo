package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
)

func ListarUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var r models.RestApi
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUsuario := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}
	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Debe enviar el parámetro 'page' como entero mayor a 0 > " + err.Error()
		return r
	}

	usuarios, status := bd.LeoUsuariosTodos(IDUsuario, int64(pageTemp), search, typeUser)
	if !status {
		r.Message = "Error al leer los usuarios"
		return r
	}

	respJson, err := json.Marshal(usuarios)
	if err != nil {
		r.Message = "Error al formatear los datos de los usuarios en tipo JSON > " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
