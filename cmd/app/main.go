package main

import (
	"log"
	"net/http"
	ini "uovh/pkg/initroutes"
)

func main() {
	mux := http.NewServeMux()
	ini.InitRotes(mux)

	srv := http.Server{
		Addr: ":8000",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	
	if err != nil{
		log.Fatal(err.Error())
	}
}