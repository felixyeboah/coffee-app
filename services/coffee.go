package services

import (
	"context"
	"time"
)

type Coffee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roast     string    `json:"roast"`
	Image     string    `json:"image"`
	Region    string    `json:"region"`
	Price     float32   `json:"price"`
	GrindUnit int16     `json:"grind_unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Coffee) GetAll() ([]*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, roast, image, region, price, grind_unit, created_at, updated_at FROM coffees`
	rows, err := db.QueryContext(ctx, query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var coffees []*Coffee
	for rows.Next() {
		coffee := &Coffee{}
		err := rows.Scan(&coffee.ID, &coffee.Name, &coffee.Roast, &coffee.Image, &coffee.Region, &coffee.Price, &coffee.GrindUnit, &coffee.CreatedAt, &coffee.UpdatedAt)
		if err != nil {
			return nil, err
		}
		coffees = append(coffees, coffee)
	}

	return coffees, nil
}

func (c *Coffee) GetByID(id string) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, roast, image, region, price, grind_unit, created_at, updated_at FROM coffees WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)

	coffee := &Coffee{}
	err := row.Scan(&coffee.ID, &coffee.Name, &coffee.Roast, &coffee.Image, &coffee.Region, &coffee.Price, &coffee.GrindUnit, &coffee.CreatedAt, &coffee.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return coffee, nil
}

func (c *Coffee) Create(coffee Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO coffees (name, roast, image, region, price, grind_unit, created_at, updated_at) VALUES ($1, $2, $3, 
$4, $5, $6, $7, $8) RETURNING *`
	_, err := db.ExecContext(ctx, query, coffee.Name, coffee.Roast, coffee.Image, coffee.Region, coffee.Price,
		coffee.GrindUnit, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	return &coffee, nil
}

func (c *Coffee) Update(id string, coffee Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE coffees SET name = $1, roast = $2, image = $3, region = $4, price = $5, grind_unit = $6, updated_at = $7 WHERE id = $8 RETURNING *`
	row := db.QueryRowContext(ctx, query, coffee.Name, coffee.Roast, coffee.Image, coffee.Region, coffee.Price, coffee.GrindUnit, time.Now(), id)

	updatedCoffee := &Coffee{}
	err := row.Scan(&updatedCoffee.ID, &updatedCoffee.Name, &updatedCoffee.Roast, &updatedCoffee.Image, &updatedCoffee.Region, &updatedCoffee.Price, &updatedCoffee.GrindUnit, &updatedCoffee.CreatedAt, &updatedCoffee.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return updatedCoffee, nil
}

func (c *Coffee) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM coffees WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
