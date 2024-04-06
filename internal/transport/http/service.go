package http

import (
	"context"
	//"encoding/csv"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/7Maliko7/april-test/internal/config"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/7Maliko7/april-test/internal/transport"
	"github.com/7Maliko7/april-test/internal/transport/structs"
	statusErr "github.com/7Maliko7/april-test/pkg/errors"
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(
	svcEndpoints transport.Endpoints, options []kithttp.ServerOption, logger log.Logger, appConfig *config.Config,
) http.Handler {
	var (
		r            = mux.NewRouter()
		errorLogger  = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)
	options = append(options, errorLogger, errorEncoder)

	carRouter := mux.NewRouter()
	carRoute := carRouter.PathPrefix("/api/v1/car").Subrouter()

	carRoute.Methods(http.MethodPost).Path("/").Handler(kithttp.NewServer(
		svcEndpoints.AddCar,
		decodeAddCarRequest,
		encodeResponse,
		options...,
	))

	carRoute.Methods(http.MethodDelete).Path("/{regNum}").Handler(kithttp.NewServer(
		svcEndpoints.DeleteCar,
		decodeDeleteCarRequest,
		encodeResponse,
		options...,
	))

	carRoute.Methods(http.MethodPost).Path("/{regNum}").Handler(kithttp.NewServer(
		svcEndpoints.UpdateCar,
		decodeUpdateCarRequest,
		encodeResponse,
		options...,
	))

	carRoute.Methods(http.MethodGet).Path("/list").Handler(kithttp.NewServer(
		svcEndpoints.GetCarList,
		decodeGetCarListRequest,
		encodeGetCarListResponse,
		options...,
	))

	initDocs(appConfig, "/api/v1")

	r.Handle("/api/v1/car/{_dummy:.*}", carRouter)
	r.Methods("GET").PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	return r
}

func decodeAddCarRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req structs.AddCarRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, statusErr.InvalidRequest
	}

	if req.RegNum == "" {
		return nil, statusErr.InvalidRequest
	}

	return req, nil
}

func decodeDeleteCarRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	req := structs.DeleteCarRequest{
		RegNum: vars["regNum"],
	}

	return req, nil
}

func decodeUpdateCarRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req structs.UpdateCarRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, statusErr.InvalidRequest
	}

	vars := mux.Vars(r)
	req.Id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, statusErr.InvalidRequest
	}

	return req, nil
}

func decodeGetCarListRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req structs.GetCarListRequest
	vars := mux.Vars(r)
	req.Limit, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, statusErr.InvalidRequest
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)

		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func encodeGetCarListResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)

		return nil
	}

	//resp, _ := response.(*structs.GetCarListResponse)

	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment;filename=report.csv")

	return nil
	//csv.NewWriter(w).WriteAll(resp.CarList)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case statusErr.CarNotFound,
		statusErr.CarNotFound,
		statusErr.DataNotFound:
		return http.StatusNotFound
	case statusErr.InvalidRequest,
		statusErr.CarExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
