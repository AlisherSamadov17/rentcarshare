package service

import (
	"rent-car/pkg/logger"
	"rent-car/storage"
)



type IServiceManager interface {
	Car() carService
    Customer() customerService
	Order() orderService
	Auth()  authService
}

type Service struct {
	carService carService
	customerService customerService
	orderService orderService
    auth authService

	logger logger.ILogger
}

func New(storage storage.IStorage,log logger.ILogger) Service  {
	services := Service{}
	services.carService = NewCarService(storage,log)
	services.customerService = NewCustomerService(storage,log)
	services.orderService = NewOrderService(storage,log)
	services.auth = NewAuthService(storage,log)
	services.logger=log

	return services
}

func (s Service) Car() carService {
	return s.carService
}

func (s Service) Customer() customerService {
	return s.customerService
}

func (s Service) Order() orderService {
	return s.orderService
}

func (s Service) Auth() authService {
	return s.auth
}