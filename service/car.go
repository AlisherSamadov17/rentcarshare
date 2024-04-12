package service

import (
	"context"
	"rent-car/api/models"
	"rent-car/pkg/logger"
	"rent-car/storage"
)


type carService struct {
	storage storage.IStorage
	logger logger.ILogger
}

func NewCarService(storage storage.IStorage,logger logger.ILogger) carService {
	return carService{
		storage: storage,
		logger: logger,
	}
}

func (u carService) Create(ctx context.Context, car models.CreateCar) (string,error) {
	pkey,err := u.storage.Car().Create(ctx,car)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating car", logger.Error(err))
		return "", err
	}
	return pkey,nil
}

func (u carService) Update(ctx context.Context, car models.Car) (string,error) {
	pkey, err := u.storage.Car().Update(ctx,car)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating car", logger.Error(err))
		return "",err
	}
	return pkey,nil
}

func (u carService) GetByIDCar(ctx context.Context, id string) (models.Car,error) {
	car, err := u.storage.Car().GetByID(ctx,id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID car", logger.Error(err))
		return models.Car{},err
	}
	return car,nil
}

func (u carService) Delete(ctx context.Context, id string) (error) {
	err := u.storage.Car().Delete(ctx,id)
	if err != nil {
		u.logger.Error("error service delete car", logger.Error(err))
		return err
	}
	return nil
}

func (u carService) GetCarAll(ctx context.Context,car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	cars, err := u.storage.Car().GetAll(ctx,car)
	if err != nil {
		u.logger.Error("error service layer while getting all cars", logger.Error(err))
		return cars,err
	}
	return cars,nil
}

func (u carService) GetAvaibleCars(ctx context.Context,car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	cars, err := u.storage.Car().GetAvaibleCars(ctx,car)
	if err != nil {
		u.logger.Error("error service layer while getting free cars", logger.Error(err))
		return cars,err
	}
	return cars,nil
}

 