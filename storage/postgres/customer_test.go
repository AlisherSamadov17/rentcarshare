package postgres

import (
	"context"
	"rent-car/api/models"
	// "rent-car/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {

	repo := NewCustomer(db,logg)

	testCustomer := models.Customer{
		FirstName:  "John",
		LastName:   "Doe",
		Gmail:      "johndoe@example.com",
		Phone:      "123456780",
		Is_Blocked: false,
	}

	id, err := repo.Create(context.Background(), testCustomer)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, id, "empty-ID returned from Create")
	}
}

func TestUpdateCustomer(t *testing.T) {

	repo := NewCustomer(db,logg)

	testCustomer := models.Customer{
		Id:         "e9ca5202-3cc7-486f-8290-1144ccfd15c8",
		FirstName:  "UpdatedFirstName",
		LastName:   "UpdatedLastName",
		Gmail:      "updatedemail@example.com",
		Phone:      "9876543210",
		Is_Blocked: false,
	}

	id, err := repo.UpdateCustomer(context.Background(), testCustomer)

	if err != nil {
		t.Errorf("UpdateCustomer  error: %v", err)
	}

	if id != testCustomer.Id {
		t.Errorf("Expected customer ID %s, but got %s", testCustomer.Id, id)
	}
}

func TestGetAllCustomer(t *testing.T) {

	repo := NewCustomer(db,logg)

	testRequest := models.GetAllCustomersRequest{
		Page:   1,
		Limit:  10,
		Search: "Alisher",
	}
	response, err := repo.GetAllCustomer(context.Background(), testRequest)

	if err != nil {
		t.Errorf("GetAllCustomer error: %v", err)
	}

	if len(response.Customers) == 0 {
		t.Errorf("Expected non-empty customer list, but got an empty list")
	}

}


func TestGetByIDCustomer(t *testing.T) {
	repo := NewCustomer(db,logg)

	testID := "4046de26-967f-4073-b8dc-e74d056b405d" 

	customer, err := repo.GetByID(context.Background(), testID)

	if err != nil {
		t.Errorf("GetByIDCustomer error: %v", err)
	}

	if customer.Id != testID {
		t.Errorf("Expected customer ID %s, but got %s", testID, customer.Id)
	}

}



func TestDeleteCustomer(t *testing.T) {
	
	repo := NewCustomer(db,logg) 

	testID := "4046de26-967f-4073-b8dc-e74d056b405d" 

	err := repo.Delete(context.Background(), testID)

	if err != nil {
		t.Errorf("Delete Customer error: %v", err)
	}
	
}
