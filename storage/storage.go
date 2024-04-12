package storage

import (
	"context"
	"rent-car/api/models"
)

type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
}

type ICarStorage interface {
	Create(context.Context,models.CreateCar) (string, error)
	GetByID(ctx context.Context,id string) (models.Car, error)
	GetAvaibleCars(ctx context.Context,req models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	GetAll(context.Context,models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(context.Context,models.Car) (string, error)
	Delete(ctx context.Context,id string) error
}

type ICustomerStorage interface {
	Create(context.Context,models.Customer) (string, error)
	GetByID(ctx context.Context,id string) (models.Customer, error)
	GetAllCustomer(ctx context.Context,req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	UpdateCustomer(context.Context,models.Customer) (string, error)
	UpdateCustomerPassword(context.Context,models.PasswordOfCustomer)(string, error)
	GetPasswordforLogin(ctx context.Context, phone string) (string, error)
	Delete(ctx context.Context,id string) error
	GetByLogin(context.Context, string) (models.GetAllCustomer, error)
	
}

type IOrderStorage interface {
	Create(context.Context,models.CreateOrder) (string, error)
	GetByID(ctx context.Context,id string) (models.OrderAll, error)
	GetAll(ctx context.Context,request models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	Update(context.Context,models.UpdateOrder) (string, error)
	Delete(ctx context.Context,id string) error
	UpdateOrderStatus(context.Context,models.GetOrder) (string, error)
}

