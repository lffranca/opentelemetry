package domain

import (
	"context"
)

type DataService[T GenericResource] interface {
	List(ctx context.Context, offset, limit int) (data *Pagination[T], err error)
	Search(ctx context.Context, text string, offset, limit int) (data *Pagination[T], err error)
	GetByID(ctx context.Context, id int) (item T, err error)
	Create(ctx context.Context, body T) (item T, err error)
	Update(ctx context.Context, body T) (err error)
	Delete(ctx context.Context, id int) (err error)
}
