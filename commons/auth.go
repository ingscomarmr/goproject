package commons

import (
	"crypto/rsa"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"os"
	"../configuration"
	"../model"
	"time"
)

//declaramos unas variables globales
var (
	privateKey *rsa.PrivateKey
	//PublicKey para validar el token
	PublicKey *rsa.PublicKey
)

//esta funcion se inicializa al importar el paquete, es como parte de una interfaz
func init(){
	conf := configuration.GetConfiguration()


	privateBytes, err := ioutil.ReadFile(conf.PathPrivateKeyRsa)
	if err != nil{
		log.Fatal("No se logro leer el archivo key privado")
	}

	publicBytes, err := ioutil.ReadFile(conf.PathPublicKeyRsa)
	if err != nil{
		log.Fatal("No se logro leer el archivo key publico")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil{
		log.Fatal("No se logro parsear la los bites de rsa")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil{
		log.Fatal("No se pudo hacer el parse de los bites rsa")
	}
	
}

//GenerateJWT: Genera el token pasando el usuario
func GenerateJWT(user model.User)  string{
	claims := model.Claim{
		User:user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:time.Now().Add(time.Hour * 2).Unix(), //dura dos horas la sesion
			Issuer: "goproyect",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	result, err := token.SignedString(privateKey) //firmamos el token con nuestra rsa privado
	if err != nil {
		log.Fatal("Error al firmar el string")
	}
	return  result
}
