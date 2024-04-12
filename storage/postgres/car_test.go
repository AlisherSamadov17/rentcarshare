package postgres

import (
	"context"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"rent-car/api/models"
	"testing"
)

func TestCreateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.CreateCar{
		Name:  faker.Name(),
		Year:  2010,
		Brand: faker.Word(),
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	if assert.NoError(t, err) {
		createdCar, err := carRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, createdCar.Name)
			assert.Equal(t, reqCar.Brand, createdCar.Brand)
		} else {
			return
		}
		fmt.Println("Created car", createdCar)
	}
}

func TestUpdate(t *testing.T) {
	repo := NewCar(db)

	testCar := models.Car{
		Id:          "cf180a59-0594-4da3-9d6c-2d79ef3c7eaf",
		Name:        "TestCar",
		Brand:       "TestBrand",
		Model:       "TestModel",
		HoursePower:      200,
		Colour:      "TestColour",
		EngineCap:   2.0,
	}

	result, err := repo.Update(context.Background(), testCar)

	if err != nil {
		t.Errorf("Update failed with error: %v", err)
	}

	if result != testCar.Id {
		t.Errorf("Expected ID %v, but got %v", testCar.Id, result)
	}
}

func TestGetAllCars(t *testing.T) {
	repo := NewCar(db)

	testRequest := models.GetAllCarsRequest{
		Page:   1,
		Limit:  10,
		Search: "BMW",
	}

	response, err := repo.GetAll(context.Background(), testRequest)

	if err != nil {
		t.Errorf("GetAllCars error: %v", err)
	}

	if len(response.Cars) == 0 {
		t.Errorf("Expected not empty car list, but got an empty list")
	}

 fmt.Println("succesfully tested")
}

func TestGetByID(t *testing.T) {

	repo := NewCar(db)

	testID := "cf180a59-0594-4da3-9d6c-2d79ef3c7eaf"

	car, err := repo.GetByID(context.Background(), testID)

	if err != nil {
		t.Errorf("GetByID failed with error: %v", err)
	}

	if car.Id != testID {
		t.Errorf("Expected car ID %s, but got %s", testID, car.Id)
	}

}

func TestDelete(t *testing.T) {
	repo := NewCar(db)

	testID := "861777c8-a985-4747-a57c-3f5ce5675051"

	err := repo.Delete(context.Background(), testID)

	if err != nil {
		t.Errorf("Delete failed with error: %v", err)
	}

}

func TestGetAvailableCars(t *testing.T) {

	carRepository := NewCar(db)

	req := models.GetAllCarsRequest{
		Page:  1,
		Limit: 10,
		Search: "Ms. Darlene Murray",
	}

	resp, err := carRepository.GetAvaibleCars(context.Background(),req)
	if err != nil {
		t.Errorf("Delete failed with error: %v", err)
	}
	fmt.Println(resp)
}




