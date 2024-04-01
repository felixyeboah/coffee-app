package services

import (
	"context"
	"time"
)

type Coffee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roast     string    `json:"roaster"`
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

	query := `SELECT id, name, roaster, image, region, price, grind_unit, created_at, updated_at FROM coffee`
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

func (c *Coffee) Create(coffee Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO coffee (name, roaster, image, region, price, grind_unit, created_at, updated_at) VALUES ($1, $2, $3, 
$4, $5, 
$6) RETURNING *`
	_, err := db.ExecContext(ctx, query, coffee.Name, coffee.Roast, coffee.Image, coffee.Region, coffee.Price,
		coffee.GrindUnit, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	return &coffee, nil
}
