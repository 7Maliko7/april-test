package postgres

import (
	"context"
	"database/sql"
	
	_ "github.com/cockroachdb/cockroach-go/crdb"

	"github.com/7Maliko7/april-test/pkg/db"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (repo *Repository) Close() error {
	return repo.db.Close()
}

func (repo *Repository) AddPeople(ctx context.Context, name, surname, patronymic string) (int64, error) {
	var id int64
	err := repo.db.QueryRowContext(ctx, AddPeopleQuery, name, surname, patronymic).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}


func (repo *Repository) AddCar(ctx context.Context, regNum, mark, model string, year, ownerId int) (int64, error) {
	var id int64
	err := repo.db.QueryRowContext(ctx, AddCarQuery, regNum, mark, model, year, ownerId).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}





func (repo *Repository) GetCar(ctx context.Context, regNum string) (*db.Car, error) {
	rows, err := repo.db.QueryContext(ctx, GetCarQuery, regNum)
	if err != nil {
		return nil, err
	}

	var s *db.Car
	for rows.Next() {
		s = &db.Car{}
		err = rows.Scan(&s.ID, &s.RegNum, &s.Mark, &s.Model, &s.Owner.Name, &s.Owner.Surname, &s.Owner.Patronymic)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (repo *Repository) DeleteCar(ctx context.Context, regNum string) (int8, error) {
	var count int8
	err := repo.db.QueryRowContext(ctx, DeleteCarQuery, regNum).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (repo *Repository) UpdateCar(ctx context.Context, carId, year, ownerId int, regNum, mark, model string) (int8, error) {
	var count int8
	err := repo.db.QueryRowContext(ctx, UpdateCarQuery, carId, regNum, mark, model, year, ownerId ).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (repo *Repository) GetCarList(ctx context.Context, limit, offset, year int, regNum, mark, model string) ([]db.Car, error) {
	rows, err := repo.db.QueryContext(ctx, GetCarListQuery, limit, offset, regNum, mark, model, year)
	if err != nil {
		return nil, err
	}

	carList := make([]db.Car, 0, 2)
	for rows.Next() {
		var ca db.Car
		err = rows.Scan(&ca.ID, &ca.RegNum, &ca.Mark, &ca.Model, &ca.Year, &ca.Owner)
		if err != nil {
			return nil, err
		}
		carList = append(carList, ca)
	}

	return carList, nil
}
