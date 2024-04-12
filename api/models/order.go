package models

type GetOrder struct {
	Id        string   `json:"id"`
	Car       Car      `json:"car"`
	Customer  Customer `json:"customer"`
	FromDate  string   `json:"from_date"`
	ToDate    string   `json:"to_date"`
	Status    string   `json:"status"`
	Paid      bool     `json:"paid"`
	Amount    float32  `json:"amount"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type CreateOrder struct {
	Id         string `json:"id"`
	CarId      string `json:"car_id"`
	CustomerId string `json:"customer_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"paid"`
	Amount     float32  `json:"amount"`
	CreatedAt string   `json:"created_at"`
}
type OrderAll struct {
	Id         string `json:"id"`
	CarId      string `json:"car_id"`
	CustomerId string `json:"customer_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"paid"`
	Amount     float32  `json:"amount"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type UpdateOrder struct {
	Id         string `json:"id"`
	CarId      string `json:"car_id"`
	CustomerId string `json:"customer_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"paid"`
	Amount     float32 `json:"amount"`
	UpdatedAt string  `json:"updated_at"`
}

type GetAllOrdersResponse struct {
	Orders []GetOrder `json:"orders"`
	Count  int        `json:"count"`
}

type GetAllOrdersRequest struct {
    Search string `json:"search"`
	Page uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}


type Order struct {
	Id         string `json:"id"`
	CarId      string `json:"car_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"paid"`
	Amount    float32  `json:"amount"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// type OrderCar struct {
// 	Id         string `json:"id"`
// 	FromDate   string `json:"from_date"`
// 	ToDate     string `json:"to_date"`
// 	CreatedAt string   `json:"created_at"`
// 	Amount     float32 `json:"amount"`
// 	Car        Car_2     `json:"car"`
// }


type SendMessage struct {
	Id        string   `json:"id"`
	CarName   string      `json:"car_name"`
	ClientName string `json:"client_name"`
	FromDate  string   `json:"from_date"`
	ToDate    string   `json:"to_date"`
	Status    string   `json:"status"`
	Paid      bool     `json:"payment_status"`
	Phone     string  `json:"phone"`
}
