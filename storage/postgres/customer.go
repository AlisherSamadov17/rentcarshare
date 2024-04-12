package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	// "database/sql"
	"fmt"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/pkg/logger"

	// "rent-car/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type customerRepo struct {
	db *pgxpool.Pool
	logger logger.ILogger
}

// // GetByLogin implements storage.ICustomerStorage.
// func (c *customerRepo) GetByLogin(context.Context, string) (models.Customer, error) {
// 	panic("unimplemented")
// }

func NewCustomer(db *pgxpool.Pool,log logger.ILogger) customerRepo {
	return customerRepo{
		db: db,
		logger: log,
	}
}

func (c *customerRepo) Create(ctx context.Context, customer models.Customer) (string, error) {
	id := uuid.New()

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return "error while hashing password here", err
	}
	customer.Password = string(hashedpassword)

	query := `insert into customers(
	id,
	first_name,
	last_name,
	gmail,
	phone,
	password,
	is_blocked) 
    values($1,$2,$3,$4,$5,$6,$7)`

	ctx, cancel := context.WithTimeout(ctx, config.TimewithContex)
	defer cancel()

	_, err = c.db.Exec(ctx, query, id.String(), customer.FirstName, customer.LastName, customer.Gmail, customer.Phone, hashedpassword, customer.Is_Blocked)
	if err != nil {
		return "error:", err
	}
	return id.String(), nil
}

func (c *customerRepo) UpdateCustomer(ctx context.Context, customer models.Customer) (string, error) {
	query := `update customers set 
	first_name=$1,
	last_name=$2,
	gmail=$3,
	phone=$4,
	is_blocked=$5,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $6 AND deleted_at = 0
	`
	ctx, cancel := context.WithTimeout(ctx, config.TimewithContex)
	defer cancel()
	_, err := c.db.Exec(ctx, query,
		customer.FirstName,
		customer.LastName,
		customer.Gmail,
		customer.Phone,
		customer.Is_Blocked,
		customer.Id)
	if err != nil {
		return "", err
	}
	return customer.Id, nil
}

// --cu.created_at,
// --cu.updated_at,

func (c *customerRepo) GetAllCustomer(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	var (
		resp   = models.GetAllCustomersResponse{}
		filter = ""
	)

	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(`and first_name ILIKE '%%%v%%'`, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query := `Select 
	cu.id as customer_id,
	cu.first_name as customer_first_name,
	cu.last_name as customer_last_name,
	cu.gmail as customer_gmail,
	cu.phone as customer_phone, 
	cu.password as customer_password,
	o.id,
	--o.from_date,
	--o.to_date,
	o.status,
	o.paid,
	o.amount
	From customers cu 
	JOIN orders o ON  cu.id = o.customer_id
	JOIN  cars ca ON ca.id = o.car_id
	`
	ctx, cancel := context.WithTimeout(ctx, config.TimewithContex)
	defer cancel()

	rows, err := c.db.Query(ctx, query+filter+``)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			customer = models.GetAllCustomer{
				Order: models.Order{},
			}
			// updateAt sql.NullString
		)
		if err := rows.Scan(
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.Gmail,
			&customer.Phone,
			&customer.Password,
			// &customer.CreatedAt,
			// &updateAt,
			&customer.Order.Id,
			// &customer.Order.FromDate,
			// &customer.Order.ToDate,
			&customer.Order.Status,
			&customer.Order.Paid,
			&customer.Order.Amount); err != nil {
			return resp, err
		}
		// customer.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Customers = append(resp.Customers, customer)
	}
	if err = rows.Err(); err != nil {
		return resp, err
	}
	countQuery := `Select count(*) from customers`
	err = c.db.QueryRow(ctx, countQuery).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *customerRepo) GetByID(ctx context.Context, id string) (models.Customer, error) {
	customer := models.Customer{}

	ctx, cancel := context.WithTimeout(ctx, config.TimewithContex)
	defer cancel()
	if err := c.db.QueryRow(ctx, `select id,first_name,last_name,gmail,phone,is_blocked from customers where id = $1`, id).Scan(
		&customer.Id,
		&customer.FirstName,
		&customer.LastName,
		&customer.Gmail,
		&customer.Phone,
		&customer.Password,
		&customer.Is_Blocked); err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepo) Delete(ctx context.Context, id string) error {
	queary := `delete from customers where id = $1`

	ctx, cancel := context.WithTimeout(ctx, config.TimewithContex)
	defer cancel()
	_, err := c.db.Exec(ctx, queary, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepo) UpdateCustomerPassword(ctx context.Context, customer models.PasswordOfCustomer) (string, error) {
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(customer.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error while hashing new password")
	}

	var Pass_cur string
	query := `select password from customers where phone = $1`

	err = c.db.QueryRow(ctx, query, customer.Phone).Scan(&Pass_cur)
	if err != nil {
		return "error:", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(Pass_cur), []byte(hashedNewPassword))
	if err != nil {
		log.Println("error while comparing new and old passwords !!!")
	}

	ctx, cancel := context.WithTimeout(ctx, config.TimewithContex)
	defer cancel()

	_, err = c.db.Exec(ctx, `update customers set password = $1 where phone = $2`, hashedNewPassword, customer.Phone)
	if err != nil {
		return "error:", err
	}

	return "OK", nil
}

func (c *customerRepo) GetPasswordforLogin(ctx context.Context, phone string) (string, error) {
	var hashedPasswordforLogin string

	query := `SELECT password
	FROM customers
	WHERE phone = $1 AND deleted_at = 0`

	err := c.db.QueryRow(ctx, query, phone).Scan(&hashedPasswordforLogin)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("phone is incorrect here")
		} else {
			return "", err
		}
	}

	return hashedPasswordforLogin, nil
}


func (c *customerRepo) GetByLogin(ctx context.Context, login string)(models.GetAllCustomer, error) {
	var (
		firstname sql.NullString
		lastname  sql.NullString
		phone     sql.NullString
		email     sql.NullString
		createdat sql.NullString
		updatedat sql.NullString
	)

	query := `SELECT 
		id, 
		first_name, 
		last_name, 
		phone,
		gmail,
		created_at, 
		updated_at,
		password
		FROM customers WHERE phone = $1 AND deleted_at = 0`

	row := c.db.QueryRow(ctx, query, login)

	customer := models.GetAllCustomer{
		Order: models.Order{},
	}

	err := row.Scan(
		&customer.Id,
		&firstname,
		&lastname,
		&phone,
		&email,
		&createdat,
		&updatedat,
		&customer.Password,
	)

	if err != nil {
		return models.GetAllCustomer{}, err
	}

	customer.FirstName = firstname.String
	customer.LastName = lastname.String
	customer.Phone = phone.String
	customer.Gmail = email.String
	customer.CreatedAt = createdat.String
	customer.UpdatedAt = updatedat.String

	return customer, nil
}
