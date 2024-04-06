package db

import (
	"context"
)

type Databaser interface {
	AddCar(ctx context.Context, regNum, mark, model string, year, ownerId int) (int64, error)
	AddPeople(ctx context.Context, name, surname, patronymic string) (int64, error)
	GetCar(ctx context.Context, regNum string) (*Car, error)
	DeleteCar(ctx context.Context, regNum string) (int8, error)
	GetCarList(ctx context.Context, limit, offset, year int, regNum, mark, model string) ([]Car, error)
	UpdateCar(ctx context.Context, carId, year, ownerId int, regNum, mark, model string) (int8, error)
}

type Car struct {
	ID     string `db:"id"`
	RegNum string `db:"regNum"`
	Mark   string `db:"mark"`
	Model  string `db:"model"`
	Year   int    `db:"year"`
	Owner  People `db:"owner"`
}

type People struct {
	Name       string `db:"name"`
	Surname    string `db:"surname"`
	Patronymic string `db:"patronymic"`
}
