package jwt

import (
	"context"
	"time"

	"github.com/Hoeru23/twittergo/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GeneroJWT(ctx context.Context, t models.Usuario) (string, error) {

	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	miClave := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
