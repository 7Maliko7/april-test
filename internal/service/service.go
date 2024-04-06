package service

import (
	"context"

	"github.com/7Maliko7/april-test/internal/transport/structs"
)

type CarCatalogService interface {
	AddCar(ctx context.Context, req structs.AddCarRequest) (*structs.AddCarResponse, error)
	DeleteCar(ctx context.Context, req structs.DeleteCarRequest) (*structs.DeleteCarResponse, error)
	UpdateCar(ctx context.Context, req structs.UpdateCarRequest) (*structs.UpdateCarResponse, error)
	GetCar(ctx context.Context, req structs.GetCarRequest) (*structs.GetCarResponse, error)
	GetCarList(ctx context.Context, req structs.GetCarListRequest) (*structs.GetCarListResponse, error)
}
