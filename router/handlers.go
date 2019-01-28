package router

import (
	"github.com/JohnnyWester/Grabber/news"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Category = "category"
const SearchCategoryResponseMessage = "Search is initialized"

func indexHandler(context *gin.Context) {
	context.String(http.StatusOK, "Hello World!")
}

func searchHandler(context *gin.Context) {
	category := context.Param(Category)
	news.Collect(category)

	context.String(http.StatusOK, SearchCategoryResponseMessage)
}

func resultHandler(context *gin.Context) {
	category := context.Param(Category)
	topics := news.Result(category)

	context.JSON(http.StatusOK, topics)
}
