package application

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/opentelemetry/internal/ggin/controller"
	"github.com/lffranca/opentelemetry/internal/repository"
	"github.com/lffranca/opentelemetry/pkg/domain"
	"github.com/lffranca/opentelemetry/pkg/domain/service"
	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func CreateExternalRequestApp() *gin.Engine {
	dataDomain := &domain.ProductResource{}

	// repositories
	flaskProductRepository := &repository.ProductRequestImpl{}

	productRepository := &repository.GenericRequestImpl[*domain.ProductResource]{
		Resource: dataDomain,
	}

	// services
	productService := service.NewDataService[*domain.ProductResource](
		dataDomain, productRepository, flaskProductRepository)

	// controllers
	productController := controller.NewGenericController[*domain.ProductResource](
		dataDomain,
		productService,
	)

	// init gin
	router := gin.Default()

	router.Use(otelgin.Middleware(os.Getenv("APPLICATION_NAME")))

	baseGroup := router.Group("/api/v1")

	// routes
	productController.Routes(baseGroup.Group("/products"))

	return router
}

func CreateRepositoryMockApp() *gin.Engine {
	dataDomain := &domain.ProductResource{}

	// repositories
	flaskProductRepository := &repository.ProductRequestImpl{}

	productRepository := &repository.GenericRepositoryImplMock[*domain.ProductResource]{
		Resource: dataDomain,
	}

	productRepository.On("List", mock.Anything, 0, 100).
		Return([]*domain.ProductResource{{ID: 124}, {ID: 125}}, 2)

	productRepository.On("Search", mock.Anything, "", 0, 100).
		Return([]*domain.ProductResource{{ID: 124}, {ID: 125}}, 2)

	productRepository.On("GetByID", mock.Anything, 123).
		Return(&domain.ProductResource{ID: 125})

	productRepository.On("Create", mock.Anything, mock.Anything).
		Return(&domain.ProductResource{ID: 125})

	productRepository.On("Update", mock.Anything, mock.Anything).
		Return()

	productRepository.On("Delete", mock.Anything, 123).
		Return()

	// services
	productService := service.NewDataService[*domain.ProductResource](
		dataDomain, productRepository, flaskProductRepository)

	// controllers
	productController := controller.NewGenericController[*domain.ProductResource](
		dataDomain,
		productService,
	)

	// init gin
	router := gin.Default()

	router.Use(otelgin.Middleware(os.Getenv("APPLICATION_NAME")))

	baseGroup := router.Group("/api/v1")

	// routes
	productController.Routes(baseGroup.Group("/products"))

	return router
}
