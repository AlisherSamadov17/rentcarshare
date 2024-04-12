package postgres

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"testing"
)

func TestCreateOrder(t *testing.T) {

	repo := NewOrder(db)

	testOrder := models.CreateOrder{
		CarId:      "4df5f517-67a6-4f8d-a35e-d99f3ddffdac",
		CustomerId: "64653dae-7bcb-47d4-9b39-19df6e447bc7",
		FromDate:   "2024-04-05",
		ToDate:     "2024-04-10",
		Status:     "canceled",
		Paid:       true,
		Amount:     100.00,
	}

	id, err := repo.Create(context.Background(), testOrder)

	if err != nil {
		t.Errorf("CreateOrder failed with error: %v", err)
	}

	if id == "" {
		t.Errorf("Expected non-empty ID returned from CreateOrder")
	}

}


func TestUpdateOrder(t *testing.T) {
	repo := NewOrder(db)

	testUpdate := models.UpdateOrder{
		Id:       "3bf82f8a-0138-4f1e-8ec1-2a68ce7978d1",      
		FromDate: "2024-04-06",   
		ToDate:   "2024-04-12",   
		Status:   "in process",    
		Paid:     true,           
		Amount:   100.00,         
	}

	id, err := repo.Update(context.Background(), testUpdate)


	if err != nil {
		t.Errorf("UpdateOrder failed with error: %v", err)
	}

	if id != testUpdate.Id {
		t.Errorf("Expected ID %s, but got %s", testUpdate.Id, id)
	}
}


func TestGetAllOrders(t *testing.T) {

	repo := NewOrder(db) 

	testRequest := models.GetAllOrdersRequest{
		Page:   1,
		Limit:  10,
		Search: "in process",
	}

    response, err := repo.GetAll(context.Background(), testRequest)

	if err != nil {
		t.Errorf("GetAllOrders failed with error: %v", err)
	}

	if len(response.Orders) == 0 {
		t.Errorf("Expected non-empty order list, but got an empty list")
	}

}


func TestGetOrderByID(t *testing.T) {
	repo := NewOrder(db)

	testID := "64c9f418-a0b9-4f18-a8b8-3e0212f0a9e0" 

	order, err := repo.GetByID(context.Background(), testID)
	
	if err != nil {
		t.Errorf("GetOrderByID failed with error: %v", err)
	}
	
	if order.Id != testID {
		t.Errorf("Expected order ID %s, but got %s", testID, order.Id)
	}
}


func TestDeleteOrder(t *testing.T) {
	repo := NewOrder(db)

	orderid := "3bf82f8a-0138-4f1e-8ec1-2a68ce7978d1" 

	err := repo.Delete(context.Background(), orderid)

	if err != nil {
		fmt.Println("Error occurred while deleting order:", err)
		t.Errorf("Failed to delete order with ID %s: %v", orderid, err)
	}
}

