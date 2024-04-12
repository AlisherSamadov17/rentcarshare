package api

import (
	"errors"
	// "log"
	"net/http"
	"rent-car/api/handler"
	"rent-car/pkg/logger"
	"rent-car/service"


	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services service.IServiceManager,log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(services,log)


	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Use(authMiddleware)
    r.POST("/customer/login",h.CustomerLogin)

	r.POST("/car", h.CreateCar)
	r.GET("/car/:id", h.GetByIDCar)
	r.GET("/cars", h.GetAllCars)
	r.GET("/availablecars", h.GetAvaibleCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	
	r.POST("/customer", h.CreateCustomer)
	r.GET("/customer/:id", h.GetByIDCustomer)
	r.GET("/customers", h.GetAllCustomer)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.PATCH("/customer/password",h.UpdateCustomerPassword)
	r.DELETE("/customer/:id", h.DeleteCustomer)

	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetAllOrder)
	r.GET("/orders", h.GetAllOrder)
	r.PATCH("/order/status/:id",h.UpdateOrderStatus)
	r.PUT("/order/:id", h.UpdateOrder)
	r.DELETE("/order/:id", h.DeleteOrder)

	return r

}


func authMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
	}
	c.Next()
}



