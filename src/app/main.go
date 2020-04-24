package main

import (
	"app/al/routes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	routes.SetUpRoutes(router)
	fmt.Println("Set up finished")
	log.Fatal(http.ListenAndServe(":8000", router))
}
