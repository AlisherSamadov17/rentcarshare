package service

import (
	"context"
	"rent-car/api/models"
	"rent-car/pkg/logger"
	"rent-car/storage"
)


type customerService struct {
	storage storage.IStorage
	logger logger.ILogger
}

func NewCustomerService(storage storage.IStorage,logger logger.ILogger) customerService {
	return customerService{
		storage: storage,
		logger: logger,
	}
}

func (cs customerService) Create(ctx context.Context, customer models.Customer) (string,error) {
	pkey,err := cs.storage.Customer().Create(ctx,customer)
	if err != nil {
		cs.logger.Error("ERROR in service layer while creating customer", logger.Error(err))
		return "", err
	}
	return pkey,nil
}

func (cs customerService) Update(ctx context.Context, customer models.Customer) (string,error) {
	pkey, err := cs.storage.Customer().UpdateCustomer(ctx,customer)
	if err != nil {
		cs.logger.Error("ERROR in service layer while updating customer", logger.Error(err))
		return "",err
	}
	return pkey,nil
}

func (cs customerService) GetByIDCustomer(ctx context.Context, id string) (models.Customer,error) {
	customer, err := cs.storage.Customer().GetByID(ctx,id)
	if err != nil {
		cs.logger.Error("ERROR in service layer while getting by id customer", logger.Error(err))
		return models.Customer{},err
	}
	return customer,nil
}

func (cs customerService) Delete(ctx context.Context, id string) (error) {
	err := cs.storage.Customer().Delete(ctx,id)
	if err != nil {
		cs.logger.Error("ERROR in service layer while deleting customer", logger.Error(err))
		return err
	}
	return nil
}

func (cs customerService) GetCustomerAll(ctx context.Context,customer models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	customers, err := cs.storage.Customer().GetAllCustomer(ctx,customer)
	if err != nil {
		cs.logger.Error("ERROR in service layer while getting all customers", logger.Error(err))
		return customers,err
	}
	return customers,nil
}

func (cs customerService) UpdatePassword(ctx context.Context, customer models.PasswordOfCustomer) (string, error) {

	pKey, err := cs.storage.Customer().UpdateCustomerPassword(ctx, customer)
	if err != nil {
		cs.logger.Error("ERROR in service layer while updating Customer", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u customerService) GetPasswordforLogin(ctx context.Context, phone string) (string, error) {
	pKey, err := u.storage.Customer().GetPasswordforLogin(ctx, phone)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID Customer",logger.Error(err))
		return "Error", err
	}

	return pKey, nil
}