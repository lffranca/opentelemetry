package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lffranca/opentelemetry/internal/gotel"
	"github.com/lffranca/opentelemetry/pkg/domain"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

type GenericRequestImpl[T domain.GenericResource] struct {
	Resource T
}

func (pkg *GenericRequestImpl[T]) List(ctx context.Context, offset, limit int) (items []T, total int, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_list", pkg.Resource.GetResource()))
	defer span.End()

	var data []byte
	var item T
	var pagination domain.Pagination[T]

	data, err = pkg.request(ctx, http.MethodGet, "", item)
	if err != nil {
		return nil, 0, err
	}

	if err = json.Unmarshal(data, &pagination); err != nil {
		return nil, 0, err
	}

	return pagination.Results, pagination.Total, nil
}

func (pkg *GenericRequestImpl[T]) Search(ctx context.Context, text string, offset, limit int) (items []T, total int, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_search", pkg.Resource.GetResource()))
	defer span.End()

	var data []byte
	var item T
	var pagination domain.Pagination[T]

	data, err = pkg.request(ctx, http.MethodGet, fmt.Sprintf("/search/%s", text), item)
	if err != nil {
		return nil, 0, err
	}

	if err = json.Unmarshal(data, &pagination); err != nil {
		return nil, 0, err
	}

	return pagination.Results, pagination.Total, nil
}

func (pkg *GenericRequestImpl[T]) GetByID(ctx context.Context, id int) (item T, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_by_id", pkg.Resource.GetResource()))
	defer span.End()

	var data []byte

	data, err = pkg.request(ctx, http.MethodGet, fmt.Sprintf("/%d", id), item)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &item); err != nil {
		return
	}

	return item, nil
}

func (pkg *GenericRequestImpl[T]) Create(ctx context.Context, body T) (item T, err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_create", pkg.Resource.GetResource()))
	defer span.End()

	var data []byte

	data, err = pkg.request(ctx, http.MethodPost, "", body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &item); err != nil {
		return
	}

	return item, nil
}

func (pkg *GenericRequestImpl[T]) Update(ctx context.Context, body T) (err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_update", pkg.Resource.GetResource()))
	defer span.End()

	var item T

	if _, err = pkg.request(ctx, http.MethodPut, "", item); err != nil {
		return err
	}

	return
}

func (pkg *GenericRequestImpl[T]) Delete(ctx context.Context, id int) (err error) {
	_, span := gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_delete", pkg.Resource.GetResource()))
	defer span.End()

	var item T

	if _, err = pkg.request(ctx, http.MethodDelete, "/123", item); err != nil {
		return err
	}

	return
}

func (pkg *GenericRequestImpl[T]) request(ctx context.Context, method, path string, body T) (responseBody []byte, err error) {
	var span trace.Span

	ctx, span = gotel.Tracer.Start(ctx, fmt.Sprintf("%s_repository_request", pkg.Resource.GetResource()))
	defer span.End()

	var req *http.Request
	var resp *http.Response
	var data []byte

	client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	data, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s%s", os.Getenv("EXTERNAL_API"), pkg.Resource.GetRoutePath(), path)
	req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
