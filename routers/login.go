package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/jwt"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.RestApi {
	var t models.Usuario
	var r models.RestApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Usuario y/o Contraseña Inválidos " + err.Error()
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "El email del usuario es requerido"
		return r
	}

	userData, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario y/o Contraseña Inválidos"
		return r
	}

	fmt.Println("GeneroJWT")
	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Ocurrió un error al intentar generar el token correspondiente > " + err.Error()
		return r
	}

	fmt.Println("Model Resp Login")
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	fmt.Println("Marshal Resp")
	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Ocurrió un error al intentar formatear el token a JSON > " + err2.Error()
		return r
	}

	fmt.Println("Cookie")
	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieString := cookie.String()

	fmt.Println("Preparo Resp")
	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow_Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
