package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"minigo/module/response"
	"net/http"
)

func SetUpRoutes(router *httprouter.Router) {
	router.GET("/", ApiInfo)
}

func ApiInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response.Data(w, http.StatusOK, fmt.Sprintf(`{"apiVersion":"%v"}`, "0.1.0"))
}
