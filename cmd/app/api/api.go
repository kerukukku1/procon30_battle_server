package api

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	jwt "github.com/dgrijalva/jwt-go"
)

func RegistHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println("regist", vars["userName"])

	// set header
    token := jwt.New(jwt.SigningMethodHS256)

    // set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["name"] = vars["userName"]

	// make JWT
	// TODO: signature
    tokenString, _ := token.SignedString([]byte("signature"))

	// return JWT
	fmt.Fprintf(w, tokenString)
}