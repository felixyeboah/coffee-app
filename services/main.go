package services

import (
	"database/sql"
	"time"
)

const dbTimeout = 5 * time.Second

var db *sql.DB

type Models struct {
	Coffee       Coffee
	JsonResponse JsonResponse
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}
