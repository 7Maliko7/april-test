package middleware

import "github.com/7Maliko7/april-test/internal/service"

// Middleware describes a service middleware.
type Middleware func(service service.CarCatalogService) service.CarCatalogService
