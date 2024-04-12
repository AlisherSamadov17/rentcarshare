package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/pkg/logger"
	"rent-car/storage"
)


type orderService struct {
	storage storage.IStorage
	logger logger.ILogger
}

func NewOrderService(storage storage.IStorage,logger logger.ILogger) orderService {
	return orderService{
		storage: storage,
		logger: logger,
	}
}

func (os orderService) Create(ctx context.Context, order models.CreateOrder) (string,error) {
	pkey,err := os.storage.Order().Create(ctx,order)
	if err != nil {
		os.logger.Error("ERROR in service layer while creating order", logger.Error(err))
		return "", err
	}
	return pkey,nil
}

func (os orderService) Update(ctx context.Context, order models.UpdateOrder) (string,error) {
	pkey, err := os.storage.Order().Update(ctx,order)
	if err != nil {
		os.logger.Error("ERROR in service layer while updating order", logger.Error(err))
		fmt.Println("",err.Error())
		return "",err
	}
	return pkey,nil
}

func (os orderService) GetByIDOrder(ctx context.Context, id string) (models.OrderAll,error) {
	order, err := os.storage.Order().GetByID(ctx,id)
	if err != nil {
		os.logger.Error("ERROR in service layer while getting by id order", logger.Error(err))
		return models.OrderAll{},err
	}
	return order,nil
}

func (os orderService) Delete(ctx context.Context, id string) (error) {
	err := os.storage.Order().Delete(ctx,id)
	if err != nil {
	os.logger.Error("ERROR in service layer while deleting order", logger.Error(err))
		return err
	}
	return nil
}

func (os orderService) GetOrderAll(ctx context.Context,order models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {
	orders, err := os.storage.Order().GetAll(ctx,order)
	if err != nil {
		os.logger.Error("ERROR in service layer while getting all orders", logger.Error(err))
		return orders,err
	}
	return orders,nil
}

func (os orderService) UpdateStatus(ctx context.Context, Order models.GetOrder) (string, error) {

	pKey, err := os.storage.Order().UpdateOrderStatus(ctx, Order)
	if err != nil {
		os.logger.Error("ERROR in service layer while updating Order", logger.Error(err))
		return "", err
	}

	return pKey, nil
}
