package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, code int, *http.Request){
	respondWithJSON(w, 200 , struct {}{})
}