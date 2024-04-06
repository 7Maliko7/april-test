package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/7Maliko7/april-test/internal/service"
	"github.com/7Maliko7/april-test/internal/transport/structs"
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.CarCatalogService) service.CarCatalogService{
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   service.CarCatalogService
	logger log.Logger
}

func (mw loggingMiddleware) AddCar(ctx context.Context, req structs.AddCarRequest) (*structs.AddCarResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "AddCar", "id", "regNum", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.AddCar(ctx, req)
}

func (mw loggingMiddleware) DeleteCar(ctx context.Context, req structs.DeleteCarRequest) (*structs.DeleteCarResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteCar", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.DeleteCar(ctx, req)
}

func (mw loggingMiddleware) UpdateCar(ctx context.Context, req structs.UpdateCarRequest) (*structs.UpdateCarResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "UpdateCar", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.UpdateCar(ctx, req)
}

func (mw loggingMiddleware) GetCar(ctx context.Context, req structs.GetCarRequest) (*structs.GetCarResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetCar", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.GetCar(ctx, req)
}

func (mw loggingMiddleware) GetCarList(ctx context.Context, req structs.GetCarListRequest) (*structs.GetCarListResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetCarList", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.GetCarList(ctx, req)
}
