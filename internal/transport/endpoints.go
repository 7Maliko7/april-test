package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/7Maliko7/april-test/internal/service"
	"github.com/7Maliko7/april-test/internal/transport/structs"
)

type Endpoints struct {
	AddCar     endpoint.Endpoint
	DeleteCar  endpoint.Endpoint
	UpdateCar  endpoint.Endpoint
	GetCar     endpoint.Endpoint
	GetCarList endpoint.Endpoint
}

func MakeEndpoints(s service.CarCatalogService) Endpoints {
	return Endpoints{
		AddCar:     makeAddCarEndpoint(s),
		DeleteCar:  makeDeleteCarEndpoint(s),
		UpdateCar:  makeUpdateCarEndpoint(s),
		GetCar:     makeGetCarEndpoint(s),
		GetCarList: makeGetCarListEndpoint(s),
	}
}

// /api/v1/car
//
//	@Summary		Метод добавления авто.
//	@Description	Метод добавления авто.
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	int
//	@Router			/api/v1/car [POST]
func makeAddCarEndpoint(s service.CarCatalogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.AddCarRequest)
		response, err := s.AddCar(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

// /api/v1/car/{regNum}
//
//	@Summary		Метод удаления авто.
//	@Description	Метод удаления авто.
//	@Produce		json
//	@Success		200		{string}	int
//	@Router			/api/v1/car/{regNum} [DELETE]
func makeDeleteCarEndpoint(s service.CarCatalogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.DeleteCarRequest)
		response, err := s.DeleteCar(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

// /api/v1/car/{id}
//
//	@Summary		Метод обновления авто.
//	@Description	Метод обновления авто.
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	string
//	@Router			/api/v1/car/{id} [POST]
func makeUpdateCarEndpoint(s service.CarCatalogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.UpdateCarRequest)
		response, err := s.UpdateCar(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

// /api/v1/car/{regNum}
//
//	@Summary		Метод получения авто.
//	@Description	Метод получения авто.
//	@Produce		json
//	@Success		200		{string}	string
//	@Router			/api/v1/car/{id} [GET]
func makeGetCarEndpoint(s service.CarCatalogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.GetCarRequest)
		response, err := s.GetCar(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

// /api/v1/car/list
//
//	@Summary		Метод получения списка авто.
//	@Description	Метод получения списка авто.
//	@Produce		json
//	@Success		200		{string}	[]string
//	@Router			/api/v1/car/list [GET]
func makeGetCarListEndpoint(s service.CarCatalogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.GetCarListRequest)
		response, err := s.GetCarList(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
