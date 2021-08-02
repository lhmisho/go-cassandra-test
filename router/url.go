package router

import (
	"github.com/julienschmidt/httprouter"
)

func UrlRouter() *httprouter.Router{
	router := httprouter.New()

	router.GET("/handle-user", HandleUser)
	return router
}
