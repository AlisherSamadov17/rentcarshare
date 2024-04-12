package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type carRepo struct {
	db *pgxpool.Pool	
}

func NewCar(db *pgxpool.Pool) carRepo {
	return carRepo{
		db: db,
	}
}

func (c *carRepo) Create(ctx context.Context,car models.CreateCar) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		model,
		hourse_power,
		colour,
		engine_cap)
		VALUES($1,$2,$3,$4,$5,$6,$7) 
	`

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()

	_,err := c.db.Exec(ctx,query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return id.String(), nil
}

func (c *carRepo) Update(ctx context.Context,car models.Car) (string, error) {

	query := ` UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $7 AND deleted_at = 0
	`
	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()

	_, err := c.db.Exec(ctx,query,
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap, car.Id)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return car.Id, nil
}

func (c *carRepo) GetAll(ctx context.Context,req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf("OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()

	rows, err := c.db.Query(ctx,`select 
				count(id) OVER(),
				id, 
				name,
				brand,
				model,
				hourse_power,
				colour,
				engine_cap,
				--created_at::date,
				updated_at,
				year
	  FROM cars WHERE deleted_at = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			car      = models.Car{}
			updateAt     sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&car.Id,
			&car.Name,
			&car.Brand,
			&car.Model,
			// &car.Year,
			&car.HoursePower,
			&car.Colour,
			&car.EngineCap,
			// &car.CreatedAt,
			&updateAt,
			&car.Year); err != nil {
			return resp, err
		}
         
		car.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Cars = append(resp.Cars, car)
	}
	return resp, nil
}

func (c *carRepo) GetAvaibleCars(ctx context.Context,req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()
	rows,err := c.db.Query(ctx,`
	SELECT name,model 
	FROM cars
	WHERE id NOT IN (SELECT car_id FROM orders WHERE car_id IS NOT NULL)`+ filter + ``)
	if err != nil {
		return resp,err
	}
	for rows.Next() {
	     var (
			car = models.Car{}
		 )

		if err := rows.Scan(

			&car.Name,
			&car.Model,);err != nil{
				return resp,err
			}

		resp.Cars = append(resp.Cars, car)	
	}
	if err = rows.Err();err != nil {
		return resp,err
	   }
	   countQuery := `SELECT COUNT(*) AS count_of_available_cars
	   FROM cars
	   WHERE id NOT IN (SELECT car_id FROM orders WHERE car_id IS NOT NULL)`

	   err = c.db.QueryRow(ctx,countQuery).Scan(&resp.Count)
		 if err != nil{
			return resp,err
		 }
	   return resp,nil
}

func (c *carRepo) GetByID(ctx context.Context,id string) (models.Car, error) {
	car := models.Car{}

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()

	if err := c.db.QueryRow(ctx,`select id,name,brand,model,hourse_power,colour,engine_cap from cars where id = $1`, id).Scan(
		&car.Id,
		&car.Name,
		&car.Brand,
		&car.Model,
		&car.HoursePower,
		&car.Colour,
		&car.EngineCap,
	); err != nil {
		return car, err
	}
	return car, nil
}

func (c *carRepo) Delete(ctx context.Context,id string) error {

	query := `delete from cars WHERE id = $1`

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()

	_, err := c.db.Exec(ctx,query, id)
	if err != nil {
		return err
	}

	return nil
}