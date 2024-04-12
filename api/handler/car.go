package handler

import (
	"context"
	"fmt"
	"net/http"
	// _ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/pkg/check"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
// @Security ApiKeyAuth
// @Router       /car [POST]
// @Summary      Creates a new cars
// @Description  create a new car
// @Tags       	 car
// @Accept       json
// @Produce      json
// @Param        car body models.Car false "car"
// @Success      201 {object} models.Car
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) CreateCar(c *gin.Context) {
	car := models.CreateCar{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handlerResponseLog(c,h.Log,"error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateCarYear(car.Year); err != nil {
		handlerResponseLog(c,h.Log,"error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())

		return
	}

    ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	id, err := h.Services.Car().Create(ctx,car)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while creating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}
// @Security ApiKeyAuth
// @Router       /car/{id} [PUT]
// @Summary      Update car
// @Description  Update car
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        id path string true "car_id"
// @Param        car body models.Car true "car"
// @Success      201 {object} models.Car
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) UpdateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handlerResponseLog(c,h.Log,"error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateCarYear(car.Year); err != nil {
		handlerResponseLog(c,h.Log,"error while validating car year,year:"+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())
		return
	}

	car.Id = c.Param("id")

	err := uuid.Validate(car.Id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while validating car id,id: "+car.Id, http.StatusBadRequest, err.Error())
		return
	}
	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	id, err := h.Services.Car().Update(ctx,car)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while updating car", http.StatusInternalServerError, err.Error())
		return
	}

	handlerResponseLog(c,h.Log,"Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router       /cars [GET]
// @Summary      Get car list
// @Description  get car list
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201 {object} models.GetAllCarsResponse
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetAllCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"Error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()
	cars, err := h.Services.Car().GetCarAll(ctx,request)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while gettign cars", http.StatusBadRequest, err.Error())
		return
	}

	handlerResponseLog(c, h.Log,"ok", http.StatusOK, cars)
}

// @Security ApiKeyAuth
// @Router       /availablecars [GET]
// @Summary      Get availablecars
// @Description  get availablecars
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201 {object} models.GetAllCarsResponse
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetAvaibleCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"Error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	
	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	cars, err := h.Services.Car().GetAvaibleCars(ctx,request)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while gettign cars", http.StatusBadRequest, err.Error())
		return
	}

	handlerResponseLog(c,h.Log,"", http.StatusOK, cars)
}

// @Security ApiKeyAuth
// @Router       /car/{id} [GET]
// @Summary      Gets car
// @Description  get car by ID
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        id path string true "car"
// @Success      201 {object} models.Car
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDCar(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id:", id)

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()
	car, err := h.Services.Car().GetByIDCar(ctx,id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while getting car by id", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c,h.Log,"", http.StatusOK, car)
}

// @Security ApiKeyAuth
// @Router       /car/{id} [DELETE]
// @Summary      Delete car
// @Description  Delete car
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        id path string true "car_id"
// @Success      201 {object} models.Response
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) DeleteCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id:", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id, err", http.StatusBadRequest, err.Error())
		return
	}

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	err = h.Services.Car().Delete(ctx,id)
	if err != nil {
		handleResponse(c, "error while deleting car", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "ok", http.StatusOK, id)
}
