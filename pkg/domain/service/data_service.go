package service

import (
	"context"
	"fmt"

	"github.com/lffranca/opentelemetry/internal/gotel"
	"github.com/lffranca/opentelemetry/pkg/domain"
)

func NewDataService[T domain.GenericResource](
	resource T,
	repository domain.DataRepository[T],
	flaskRepository domain.DataRepository[T],
) domain.DataService[T] {
	return &DataServiceImpl[T]{
		resource,
		repository,
		flaskRepository,
	}
}

type DataServiceImpl[T domain.GenericResource] struct {
	resource        T
	repository      domain.DataRepository[T]
	flaskRepository domain.DataRepository[T]
}

func (pkg *DataServiceImpl[T]) List(ctx context.Context, offset, limit int) (data *domain.Pagination[T], err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_service_list", pkg.resource.GetResource()))
	defer span.End()

	var items []T
	var total int

	var itemsFlask []T
	var totalFlask int

	items, total, err = pkg.repository.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	itemsFlask, totalFlask, err = pkg.flaskRepository.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	items = append(items, itemsFlask...)
	total += totalFlask

	pagination := &domain.Pagination[T]{
		Results: items,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return pagination, nil
}

func (pkg *DataServiceImpl[T]) Search(ctx context.Context, text string, offset, limit int) (data *domain.Pagination[T], err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_service_search", pkg.resource.GetResource()))
	defer span.End()

	var items []T
	var total int

	items, total, err = pkg.repository.Search(ctx, text, offset, limit)
	if err != nil {
		return nil, err
	}

	pagination := &domain.Pagination[T]{
		Results: items,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return pagination, nil
}

func (pkg *DataServiceImpl[T]) GetByID(ctx context.Context, id int) (item T, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_service_by_id", pkg.resource.GetResource()))
	defer span.End()

	item, err = pkg.repository.GetByID(ctx, id)
	if err != nil {
		return
	}

	return item, nil
}

func (pkg *DataServiceImpl[T]) Create(ctx context.Context, body T) (item T, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_service_create", pkg.resource.GetResource()))
	defer span.End()

	item, err = pkg.repository.Create(ctx, body)
	if err != nil {
		return
	}

	return item, nil
}

func (pkg *DataServiceImpl[T]) Update(ctx context.Context, body T) (err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_service_update", pkg.resource.GetResource()))
	defer span.End()

	if err = pkg.repository.Update(ctx, body); err != nil {
		return err
	}

	return nil
}

func (pkg *DataServiceImpl[T]) Delete(ctx context.Context, id int) (err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_service_delete", pkg.resource.GetResource()))
	defer span.End()

	if err = pkg.repository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
