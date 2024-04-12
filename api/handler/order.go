package handler

import (
	"context"
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/pkg/check"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateOrder godoc
// @Router       /order [POST]
// @Summary      Creates a new orders
// @Description  create a new order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        car body models.Order false "order"
// @Success      201 {object} models.CreateOrder
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) CreateOrder(c *gin.Context) {
	order := models.CreateOrder{}

	data, err := getAuthInfo(c)
	if err != nil {
		handlerResponseLog(c, h.Log, "error while getting auth", http.StatusUnauthorized, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		handlerResponseLog(c,h.Log,"error while reding request", http.StatusBadRequest, err.Error())
		return
	}

    if err := check.ValidateDateOfFormatForOrder(order.FromDate);err != nil{
		handlerResponseLog(c,h.Log,"error in FromDate",http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateDateOfFormatForOrder(order.ToDate);err != nil{
		handlerResponseLog(c,h.Log,"error in ToDate",http.StatusBadRequest, err.Error())
		return
	}
	
	order.Status = config.STATUS_NEW
	order.CustomerId = data.UserID
   
	if err := check.ValidatingOrderStatusForAuth(order.Status);err != nil {
		handlerResponseLog(c,h.Log,"error in order releted with status",http.StatusUnauthorized, err.Error())
		return
	}

	order.CustomerId = c.Param("id")
	err = uuid.Validate(order.CustomerId)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while validating customer id,id: "+order.CustomerId, http.StatusBadRequest, err.Error())
		return
	}

	order.CarId = c.Param("id")
	err = uuid.Validate(order.CarId)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while validating car id,id: "+order.CarId, http.StatusBadRequest, err.Error())
		return
	}

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	id, err := h.Services.Order().Create(ctx,order)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while creating order", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c,h.Log,"ok", http.StatusOK, id)
}

// @Security ApiKeyAuth
// UpdateOrder godoc
// @Router       /order/{id} [PUT]
// @Summary      Update order
// @Description  Update order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id path string true "order_id"
// @Param        order body models.Order true "order"
// @Success      201 {object} models.Order
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) UpdateOrder(c *gin.Context) {
	order := models.UpdateOrder{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handlerResponseLog(c,h.Log,"error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	order.Id = c.Param("id")
	err := uuid.Validate(order.Id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while validating", http.StatusBadRequest, err.Error())
		return
	}

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	id, err := h.Services.Order().Update(ctx,order)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while updating customer,err", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c,h.Log,"ok", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GetOrderList godoc
// @Router       /orders [GET]
// @Summary      Get order list
// @Description  get order list
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201 {object} models.GetAllOrdersResponse
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetAllOrder(c *gin.Context) {
	var (
		request = models.GetAllOrdersRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	orders, err := h.Services.Order().GetOrderAll(ctx,request)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while getting orders", http.StatusInternalServerError, err.Error())
		return
	}

	handlerResponseLog(c, h.Log,"ok", http.StatusOK, orders)
}

// @Security ApiKeyAuth
// GetOrder godoc
// @Router       /order/{id} [GET]
// @Summary      Gets order
// @Description  get order by ID
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id path string true "order"
// @Success      201 {object} models.Order
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDOrder(c *gin.Context) {
	id := c.Param("id")

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	order, err := h.Services.Order().GetByIDOrder(ctx,id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while getting order by id", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c, h.Log,"ok", http.StatusOK, order)
}

// @Security ApiKeyAuth
// DeleteOrder godoc
// @Router       /order/{id} [DELETE]
// @Summary      Delete order
// @Description  Delete order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id path string true "order_id"
// @Success      201 {object} models.Response
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := uuid.Validate(id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	err = h.Services.Customer().Delete(ctx,id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while deleting customer", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c,h.Log,"ok", http.StatusOK, id)
}


// @Security ApiKeyAuth
// UpdateOrderStatus godoc
// @Router 		/order/status/{id} [PATCH]
// @Summary 	update a order
// @Description This api is update a  order status and returns it's id
// @Tags 		order
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Param		status path string true "status"
// @Success		200  {object}  models.GetOrder
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateOrderStatus(c *gin.Context) {
	Order := models.GetOrder{}

	if err := c.ShouldBindJSON(&Order); err != nil {
		handlerResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	Order.Id = c.Param("id")
	Order.Status = c.Param("status")



	if err := check.ValidatingOrderStatusForAuth(Order.Status); err != nil {
		handlerResponseLog(c,  h.Log,"error check order status: "+Order.Status, http.StatusBadRequest,err.Error())
		return
	}

	err := uuid.Validate(Order.Id)
	if err != nil {
		handlerResponseLog(c, h.Log, "error while validating Order id,id: "+Order.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Order().UpdateStatus(context.Background(),Order)
	if err != nil {
		handlerResponseLog(c, h.Log, "error while updating Order status", http.StatusBadRequest, err.Error())
		return
	}

	handlerResponseLog(c, h.Log, "Updated successfully", http.StatusOK, id)
}