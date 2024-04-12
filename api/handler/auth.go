package handler

import (
	"fmt"
	"net/http"
	"rent-car/api/models"
	_"rent-car/api/docs"
	"github.com/gin-gonic/gin"
)

// CustomerLogin godoc
// @Router       /customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest true "login"
// @Success      201  {object}  models.CustomerLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) CustomerLogin(c *gin.Context)  {
	loginReq := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq);err != nil {
		handlerResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq:",loginReq)

	loginResp,err := h.Services.Auth().CustomerLogin(c.Request.Context(),loginReq)
	if err != nil {
		handlerResponseLog(c,h.Log,"unauthorized",http.StatusUnauthorized,err)
		return
	}
	handlerResponseLog(c,h.Log,"succes",http.StatusOK,loginResp)
}

