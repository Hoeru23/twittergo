package bd

import (
	"fmt"

	"github.com/Hoeru23/twittergo/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	fmt.Println("> ChequeoYaExisteUsuario")
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	fmt.Println("> Comparo con el bcrypt")
	err := bcrypt.CompareHashAndPassword(passwordBytes, passwordBD)
	if err != nil {
		return usu, false
	}

	return usu, true
}
