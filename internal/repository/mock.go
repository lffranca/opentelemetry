package repository

import (
	"context"
	"fmt"

	"github.com/lffranca/opentelemetry/internal/gotel"
	"github.com/lffranca/opentelemetry/pkg/domain"
	"github.com/stretchr/testify/mock"
)

type GenericRepositoryImplMock[T domain.GenericResource] struct {
	mock.Mock
	Resource T
}

func (pkg *GenericRepositoryImplMock[T]) List(ctx context.Context, offset, limit int) (items []T, total int, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_list", pkg.Resource.GetResource()))
	defer span.End()

	args := pkg.MethodCalled("List", ctx, offset, limit)

	if len(args) > 0 && args.Get(0) != nil {
		items = args.Get(0).([]T)
	}

	if len(args) > 1 && args.Get(1) != nil {
		total = args.Get(1).(int)
	}

	if len(args) > 2 && args.Get(2) != nil {
		err = args.Error(2).(error)
	}

	return
}

func (pkg *GenericRepositoryImplMock[T]) Search(ctx context.Context, text string, offset, limit int) (items []T, total int, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_search", pkg.Resource.GetResource()))
	defer span.End()

	args := pkg.MethodCalled("Search", ctx, text, offset, limit)

	if len(args) > 0 && args.Get(0) != nil {
		items = args.Get(0).([]T)
	}

	if len(args) > 1 && args.Get(1) != nil {
		total = args.Get(1).(int)
	}

	if len(args) > 2 && args.Get(2) != nil {
		err = args.Error(2).(error)
	}

	return
}

func (pkg *GenericRepositoryImplMock[T]) GetByID(ctx context.Context, id int) (item T, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_by_id", pkg.Resource.GetResource()))
	defer span.End()

	args := pkg.MethodCalled("GetByID", ctx, id)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(T)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1).(error)
	}

	return
}

func (pkg *GenericRepositoryImplMock[T]) Create(ctx context.Context, body T) (item T, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_create", pkg.Resource.GetResource()))
	defer span.End()

	args := pkg.MethodCalled("Create", ctx, body)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(T)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1).(error)
	}

	return
}

func (pkg *GenericRepositoryImplMock[T]) Update(ctx context.Context, body T) (err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_update", pkg.Resource.GetResource()))
	defer span.End()

	args := pkg.MethodCalled("Update", ctx, body)

	if len(args) > 0 && args.Get(0) != nil {
		err = args.Error(0).(error)
	}

	return
}

func (pkg *GenericRepositoryImplMock[T]) Delete(ctx context.Context, id int) (err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_delete", pkg.Resource.GetResource()))
	defer span.End()

	args := pkg.MethodCalled("Delete", ctx, id)

	if len(args) > 0 && args.Get(0) != nil {
		err = args.Error(0).(error)
	}

	return
}
