package api

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func RegistHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Fprintf(w, "<h1>regist</h1>")
	fmt.Println("regist", vars["userName"])
}