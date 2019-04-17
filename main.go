package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers () *httprouter.Router {
	router := httprouter.New()

	router.GET("/dnmonster/:name", dnmonsterHandeler)

	//router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router

}

func dnmonsterHandeler (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	//get name from params
	name := string(params.ByName("name"))

	//get sha1(Hexadecimal) of name
	i16 := Sha256(name)

	//change Hexadecimal to binary
	i2 := ChanHexToBin(i16)

	//create png of i2
	CreatePng(i2)

	imagePath := "./out.png"

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)

}

func main(){
	r := RegisterHandlers ()
	http.ListenAndServe(":9000", r)
}
