package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/pkg/jwt"
	"rent-car/pkg/logger"
	"rent-car/pkg/logger/password"
	"rent-car/storage"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewAuthService(storage storage.IStorage, log logger.ILogger) authService {
	return authService{
		storage: storage,
		log:     log,
	}
}

func (a authService) CustomerLogin(ctx context.Context, loginRequest models.CustomerLoginRequest) (models.CustomerLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Login)
	customer, err := a.storage.Customer().GetByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error while getting customer credentials by login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	if err = password.CompareHashAndPassword(customer.Password, loginRequest.Password); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = customer.Id
	m["user_role"] = config.CUSTOMER_ROLE

    accessToken,refreshToken,err :=jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for customer login",logger.Error(err))
	}

	return models.CustomerLoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	},nil
}
