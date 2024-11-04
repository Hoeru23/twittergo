package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RestApi {
	var mensaje models.Tweet
	var r models.RestApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)
	if err != nil {
		r.Message = "Ocurrió un error al intentar decodificar el body > " + err.Error()
		return r
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTwwet(registro)
	if err != nil {
		r.Message = "Ocurrió un error al intentar insertar el registro > " + err.Error()
		return r
	}
	if !status {
		r.Message = "NO se ha logrado insertar el registro"
		return r
	}

	r.Status = 200
	r.Message = "Tweet creado correctamente"
	return r
}
