package http

import (
	"github.com/7Maliko7/april-test/docs"
	"github.com/7Maliko7/april-test/internal/config"
)

func initDocs(appConfig *config.Config, basePath string) {
	docs.SwaggerInfo.Title = "Car Catalog API"
	docs.SwaggerInfo.Description = "This is implementation of car catalog."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = appConfig.ListenAddress
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http"}
}
