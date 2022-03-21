package http

import (
	"net/http"

	"github.com/amitabhprasad/bookstore-util-go/rest_errors"

	token "github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/domain/access_token"
	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/service/access_token"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at token.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := rest_errors.NewbadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	access_token, err := h.service.Create(at)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, access_token)
}
func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	var at token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := rest_errors.NewbadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.UpdateExpirationTime(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}
