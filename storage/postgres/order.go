package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/pkg"
	"rent-car/tgbot"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrder(db *pgxpool.Pool) orderRepo {
	return orderRepo{
		db: db,
	}
}

func (o *orderRepo) Create(ctx context.Context, or models.CreateOrder) (string, error) {
	id := uuid.New()

	query := `insert into orders(
		id,
		car_id,
		customer_id,
		from_date,
		to_date,
		status,
		paid,
		amount
	) values($1,$2,$3,$4,$5,$6,$7,$8)`

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()
	_, err := o.db.Exec(ctx, query, id.String(), or.CarId, or.CustomerId, or.FromDate, or.ToDate, or.Status, or.Paid, or.Amount)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (o *orderRepo) Update(ctx context.Context, or models.UpdateOrder) (string, error) {
	query := `update orders set
	   from_date=$1,
	   to_date=$2,
	   status=$3,
	   paid=$4,
	   amount=$5,
       updated_at=CURRENT_TIMESTAMP
	   WHERE id=$6
	`
	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()
	_, err := o.db.Exec(ctx, query, or.FromDate, or.ToDate, or.Status, or.Paid, or.Amount, or.Id)
	if err != nil {
		return "", err
	}
	return or.Id, nil
}

func (o *orderRepo) GetAll(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {
	var (
		resp   = models.GetAllOrdersResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(`and  status ILIKE '%%%v%%'`, req.Search)
	}
	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query := `Select 
	o.id,
	--o.from_date,
	--o.to_date,
	o.status,
	o.paid,
	o.amount,
	--o.created_at,
	--o.updated_at,
	c.name as car_name,
	c.brand as car_brand,
	c.engine_cap as car_engine_cap,
	cu.id as customer_id,
	cu.first_name as customer_first_name,
	cu.last_name as customer_last_name,
	cu.gmail as customer_gmail,
	cu.phone as customer_phone,
	cu.password as customer_password,
	From orders o JOIN cars c ON o.car_id = c.id
	JOIN customers cu ON o.customer_id = cu.id 	`
	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()
	rows, err := o.db.Query(ctx, query+filter+``)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			order = models.GetOrder{
				Car:      models.Car{},
				Customer: models.Customer{},
			}
			updateAt sql.NullString
		)

		err := rows.Scan(
			&order.Id,
			// &order.FromDate,
			// &order.ToDate,
			&order.Status,
			&order.Paid,
			&order.Amount,
			// &order.CreatedAt,
			// &updateAt,
			&order.Car.Name,
			&order.Car.Brand,
			&order.Car.EngineCap,
			&order.Customer.Id,
			&order.Customer.FirstName,
			&order.Customer.LastName,
			&order.Customer.Gmail, &order.Customer.Phone)
		if err != nil {
			return resp, err
		}
		order.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Orders = append(resp.Orders, order)
	}
	if err = rows.Err(); err != nil {
		return resp, err
	}
	countQuery := `SELECT COUNT(*) FROM orders`

	err = o.db.QueryRow(ctx, countQuery).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (o *orderRepo) GetByID(ctx context.Context, id string) (models.OrderAll, error) {
	order := models.OrderAll{}

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()

	if err := o.db.QueryRow(ctx,`Select 
	 	o.id as order_id,
		o.car_id,
		o.customer_id,
	 	--o.from_date,
	 	--o.to_date,
	 	o.status,
	 	o.paid
	 	--o.created_at,
	 	--o.updated_at 
		from orders o
	 	where id = $1`, id).Scan(
		&order.Id,
		&order.CarId,
		&order.CustomerId,
		// &order.FromDate,
		// &order.ToDate,
		&order.Status,
		&order.Paid,
	// &order.CreatedAt,
	//&order.UpdatedAt
	); err != nil {
		return models.OrderAll{}, err
	}
	return order, nil
}


func (o *orderRepo) Delete(ctx context.Context, id string) error {
    query := `DELETE FROM orders WHERE id = $1`

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()
    _, err := o.db.Exec(ctx, query, id)
    if err != nil {
        return err
    }
    return nil
}

func (o *orderRepo) UpdateOrderStatus(ctx context.Context,order models.GetOrder) (string, error) {
	query := `update orders set 
        status = $1,
        updated_at = CURRENT_TIMESTAMP
        where id = $2`

	_, err := o.db.Exec(ctx,query,
		order.Status, 
		order.Id)

	if err != nil {
		return "", err
	}

   MSGinfo,err:=o.GetMSGINFO(context.Background(),order.Id)
if err != nil {
	return "", err
}

err=tgbot.SendMessageTG(MSGinfo)
if err != nil {
	return "", err
}
	return order.Id, nil
}



func (o *orderRepo) GetMSGINFO(ctx context.Context,orderID string) (models.SendMessage, error) {
	order :=models.SendMessage{}

	query := `SELECT
		o.id,
		c.name AS car_name,
		cu.first_name AS customer_first_name,
		o.from_date,
		o.to_date,
		o.status,
		o.paid,
		cu.phone
		FROM orders o
		JOIN cars c ON o.car_id = c.id
		JOIN customers cu ON o.customer_id = cu.id
		WHERE o.id = $1`

	row := o.db.QueryRow(ctx,query, orderID)
	err := row.Scan(
		&order.Id,
		&order.CarName,
		&order.ClientName,
		&order.FromDate,
		&order.ToDate,
		&order.Status,
		&order.Paid,
		&order.Phone,
	)
	if err != nil {
		return models.SendMessage{}, err
	}

	return order, nil
}
