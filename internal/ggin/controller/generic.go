package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/opentelemetry/internal/ggin"
	"github.com/lffranca/opentelemetry/internal/gotel"
	"github.com/lffranca/opentelemetry/pkg/domain"
)

func NewGenericController[T domain.GenericResource](
	resource T,
	service domain.DataService[T],
) ggin.GenericController[T] {
	return &GenericControllerImpl[T]{
		resource: resource,
		service:  service,
	}
}

type GenericControllerImpl[T domain.GenericResource] struct {
	resource T
	service  domain.DataService[T]
}

func (pkg *GenericControllerImpl[T]) List(ctx *gin.Context) {
	_, span := gotel.Tracer.Start(ctx.Request.Context(), fmt.Sprintf("%s_controller_list", pkg.resource.GetResource()))
	defer span.End()

	data, err := pkg.service.List(ctx.Request.Context(), 0, 100)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, data)
}

func (pkg *GenericControllerImpl[T]) Search(ctx *gin.Context) {
	_, span := gotel.Tracer.Start(ctx.Request.Context(), fmt.Sprintf("%s_controller_search", pkg.resource.GetResource()))
	defer span.End()

	data, err := pkg.service.Search(ctx.Request.Context(), "", 0, 100)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, data)
}

func (pkg *GenericControllerImpl[T]) GetByID(ctx *gin.Context) {
	_, span := gotel.Tracer.Start(ctx.Request.Context(), fmt.Sprintf("%s_controller_by_id", pkg.resource.GetResource()))
	defer span.End()

	data, err := pkg.service.GetByID(ctx.Request.Context(), 123)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, data)
}

func (pkg *GenericControllerImpl[T]) Create(ctx *gin.Context) {
	_, span := gotel.Tracer.Start(ctx.Request.Context(), fmt.Sprintf("%s_controller_create", pkg.resource.GetResource()))
	defer span.End()

	var body T
	data, err := pkg.service.Create(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusCreated, data)
}

func (pkg *GenericControllerImpl[T]) Update(ctx *gin.Context) {
	_, span := gotel.Tracer.Start(ctx.Request.Context(), fmt.Sprintf("%s_controller_update", pkg.resource.GetResource()))
	defer span.End()

	var body T
	if err := pkg.service.Update(ctx.Request.Context(), body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (pkg *GenericControllerImpl[T]) Delete(ctx *gin.Context) {
	_, span := gotel.Tracer.Start(ctx.Request.Context(), fmt.Sprintf("%s_controller_delete", pkg.resource.GetResource()))
	defer span.End()

	if err := pkg.service.Delete(ctx.Request.Context(), 123); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (pkg *GenericControllerImpl[T]) Routes(router gin.IRoutes) {
	router.GET("", pkg.List)
	router.GET("/search/:query", pkg.Search)
	router.GET("/:id", pkg.GetByID)
	router.POST("", pkg.Create)
	router.PUT("", pkg.Update)
	router.DELETE("/:id", pkg.Delete)
}
