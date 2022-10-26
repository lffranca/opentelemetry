package ggin

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/opentelemetry/pkg/domain"
)

type GenericController[T domain.GenericResource] interface {
	List(ctx *gin.Context)
	Search(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Routes(router gin.IRoutes)
}
