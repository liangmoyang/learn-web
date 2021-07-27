package main

import (
	"log"
	"net/http"
	"web/server"
)

func main() {

	svr := server.NewHttpServer("test-server")

	//svr.Route("/home",home)
	svr.Route("POST", "/user/signup", server.Signup)

	log.Fatalln(svr.Start(":8080"))

}

func home(w http.ResponseWriter, r *http.Request) {

}
