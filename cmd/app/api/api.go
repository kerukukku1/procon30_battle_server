package api

import (
	"fmt"
	"net/http"
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
	"github.com/gorilla/mux"
)

func RegistHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println("regist", vars["userName"])

	// read salt
	salt := []byte("some salt")

	// convert hash
	converted, _ := scrypt.Key([]byte(vars["userName"]), salt, 32768, 8, 1, 32)
	token := hex.EncodeToString(converted[:])

	// return JWT
	fmt.Fprintf(w, token)
}