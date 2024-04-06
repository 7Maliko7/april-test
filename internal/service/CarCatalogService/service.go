package car_catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/7Maliko7/april-test/internal/config"
	svc "github.com/7Maliko7/april-test/internal/service"
	"github.com/7Maliko7/april-test/internal/transport/structs"
	"github.com/7Maliko7/april-test/pkg/db"
	"github.com/7Maliko7/april-test/pkg/errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
)

const (
	logKeyMethod = "method"
)

type CarCatalog struct {
	Repository db.Databaser
	Logger     log.Logger
	Config     *config.Config
}

func NewService(cfg *config.Config, rep db.Databaser, logger log.Logger) svc.CarCatalogService {
	return &CarCatalog{
		Config:     cfg,
		Repository: rep,
		Logger:     logger,
	}
}

func (s *CarCatalog) AddCar(ctx context.Context, req structs.AddCarRequest) (*structs.AddCarResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Create")

	car, err := s.Repository.GetCar(ctx, req.RegNum)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}

	if car != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.CarExists
	}

	resp, err := http.Get(fmt.Sprintf("%v/info?regNum=%v", s.Config.CarInfoAddress, req.RegNum))
	if err != nil {
		level.Error(logger).Log("client", err.Error())
		return nil, errors.FailedRequest
	}
	if resp == nil {
		level.Error(logger).Log("client", err.Error())
		return nil, errors.DataNotFound
	}
	if resp.StatusCode != 200 {
		level.Error(logger).Log("client", err.Error())
		return nil, errors.DataNotFound
	}
	carInfo := Car{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&carInfo)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}

	ownerId, err := s.Repository.AddPeople(ctx, carInfo.Owner.Name, carInfo.Owner.Surname, carInfo.Owner.Patronymic)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}

	id, err := s.Repository.AddCar(ctx, carInfo.RegNum, carInfo.Mark, carInfo.Model, carInfo.Year, int(ownerId))
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	level.Debug(logger).Log("New Car", id)

	return &structs.AddCarResponse{Id: id}, nil
}

func (s *CarCatalog) DeleteCar(ctx context.Context, req structs.DeleteCarRequest) (*structs.DeleteCarResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Delete Car")

	count, err := s.Repository.DeleteCar(ctx, req.RegNum)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	level.Debug(logger).Log("Deleted Car", count)

	if count == 0 {
		level.Debug(logger).Log("Deleted Car", count)
		return nil, errors.CarNotFound
	}

	return &structs.DeleteCarResponse{Response: structs.Response{Status: "success"}}, nil
}

func (s *CarCatalog) UpdateCar(ctx context.Context, req structs.UpdateCarRequest) (*structs.UpdateCarResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Update Car")
	
	count, err := s.Repository.UpdateCar(ctx, int(req.Id), req.Car.Year, req.OwnerID, req.Car.Mark, req.Car.Mark, req.Car.Model)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	level.Debug(logger).Log("Updated Car", count)

	if count == 0 {
		level.Debug(logger).Log("Updated Car", count)
		return nil, errors.CarNotFound
	}
	
	return &structs.UpdateCarResponse{Response: structs.Response{Status: "success"}}, nil
}

func (s *CarCatalog) GetCar(ctx context.Context, req structs.GetCarRequest) (*structs.GetCarResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Get Car")

	car, err := s.Repository.GetCar(ctx, req.RegNum)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	if car == nil {
		level.Error(logger).Log("repository", "Data not found")
		return nil, errors.DataNotFound
	}

	return &structs.GetCarResponse{
		Car: structs.Car{
			RegNum: car.RegNum,
			Mark:   car.Mark,
			Model:  car.Model,
			Year:   car.Year,
			Owner:  structs.People(car.Owner),
		},
	}, nil
}

func (s *CarCatalog) GetCarList(ctx context.Context, req structs.GetCarListRequest) (*structs.GetCarListResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Get Car List")

	list, err := s.Repository.GetCarList(ctx, int(req.Limit), int(req.Offset), req.Year, req.RegNum, req.Mark, req.Model)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}

	result := make([]structs.Car, 0, len(list))
	for _, v := range list {
		result = append(result, structs.Car{RegNum: v.RegNum, Mark: v.Mark, Model: v.Model, Year: v.Year, Owner: structs.People(v.Owner)})
	}

	return &structs.GetCarListResponse{CarList: result}, nil
}
