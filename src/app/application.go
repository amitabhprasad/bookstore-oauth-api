package app

import (
	"github.com/gin-gonic/gin"

	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/http"
	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/repository/db"
	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/repository/rest"
	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/service/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token", atHandler.UpdateExpirationTime)

	router.Run(":8082")
}
