package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/search/:category", searchHandler)
	router.GET("/result/:category", resultHandler)

	return router
}
