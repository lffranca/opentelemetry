package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lffranca/opentelemetry/internal/gotel"
	"github.com/lffranca/opentelemetry/pkg/domain"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

type flaskResponse struct {
	Products []*domain.ProductResource `json:"products"`
}

type ProductRequestImpl struct {
}

func (pkg *ProductRequestImpl) List(ctx context.Context, offset, limit int) (items []*domain.ProductResource, total int, err error) {
	_, span := gotel.Tracer.Start(ctx, "products_repository_flask_list")
	defer span.End()

	var data []byte
	var response flaskResponse

	data, err = pkg.request(ctx, http.MethodGet, "", nil)
	if err != nil {
		return nil, 0, err
	}

	if err = json.Unmarshal(data, &response); err != nil {
		return nil, 0, err
	}

	return response.Products, len(response.Products), nil
}

func (pkg *ProductRequestImpl) Search(ctx context.Context, text string, offset, limit int) (items []*domain.ProductResource, total int, err error) {
	_, span := gotel.Tracer.Start(ctx, "products_repository_flask_search")
	defer span.End()

	var data []byte
	var response flaskResponse

	data, err = pkg.request(ctx, http.MethodGet, fmt.Sprintf("/search/%s", text), nil)
	if err != nil {
		return nil, 0, err
	}

	if err = json.Unmarshal(data, &response); err != nil {
		return nil, 0, err
	}

	return response.Products, len(response.Products), nil
}

func (pkg *ProductRequestImpl) GetByID(ctx context.Context, id int) (item *domain.ProductResource, err error) {
	_, span := gotel.Tracer.Start(ctx, "products_repository_flask_by_id")
	defer span.End()

	var data []byte

	data, err = pkg.request(ctx, http.MethodGet, fmt.Sprintf("/%d", id), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &item); err != nil {
		return
	}

	return item, nil
}

func (pkg *ProductRequestImpl) Create(ctx context.Context, body *domain.ProductResource) (item *domain.ProductResource, err error) {
	_, span := gotel.Tracer.Start(ctx, "products_repository_flask_create")
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

func (pkg *ProductRequestImpl) Update(ctx context.Context, body *domain.ProductResource) (err error) {
	_, span := gotel.Tracer.Start(ctx, "products_repository_flask_update")
	defer span.End()

	if _, err = pkg.request(ctx, http.MethodPut, "", nil); err != nil {
		return err
	}

	return
}

func (pkg *ProductRequestImpl) Delete(ctx context.Context, id int) (err error) {
	_, span := gotel.Tracer.Start(ctx, "products_repository_flask_delete")
	defer span.End()

	if _, err = pkg.request(ctx, http.MethodDelete, "/1", nil); err != nil {
		return err
	}

	return
}

func (pkg *ProductRequestImpl) request(ctx context.Context, method, path string, body interface{}) (responseBody []byte, err error) {
	var span trace.Span

	ctx, span = gotel.Tracer.Start(ctx, "products_repository_flask_request")
	defer span.End()

	var req *http.Request
	var resp *http.Response
	var data []byte

	client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("http://localhost:5001/api/v1/products%s", path)
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
