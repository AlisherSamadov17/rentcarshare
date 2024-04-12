package models

type Customer struct{
	Id          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gmail       string  `json:"gmail"`
	Phone       string  `json:"phone"`
	Password    string  `json:"password"`
 	Is_Blocked    bool  `json:"isblocked"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetAllCustomer struct{
	Id          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gmail       string  `json:"gmail"`
	Phone       string  `json:"phone"`
	Password    string  `json:"password"`
	Is_Blocked  bool    `json:"isblocked"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	Order       Order   `json:"order"`
}

type PasswordOfCustomer struct {
	Phone       string  `json:"phone"`
	NewPassword string `json:"new_password"`
	Password string `json:"password"`
}

type GetAllCustomersResponse struct {
	Customers []GetAllCustomer `json:"customers"`
	Count int16 `json:"count"`
}

type GetAllCustomersRequest struct {
    Search string `json:"search"`
	Page uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}
     
// type GetAllCustomerCars struct{
// 	Id          string  `json:"id"`
// 	Name        string    `json:"name"`
// 	CreatedAt   string    `json:"creatAt"`
// 	Amount      float32    `json:"amount"`
// }

// type GetAllCustomerCarsResponse struct {
// 	Customer []GetAllCustomerCars `json:"orders"`
// 	Count int16 `json:"count"`
// }

// type GetAllCustomerCarsRequest struct {
//     Search string `json:"search"`
// 	Id     string  `json:"id"`
// 	Page uint64 `json:"page"`
// 	Limit uint64 `json:"limit"`
// }

