package domain

type GenericResource interface {
	DataDomain
}

type DataDomain interface {
	GetResource() string
	GetRoutePath() string
}

type ProductResource Product

func (item *ProductResource) GetResource() string {
	return "product"
}

func (item *ProductResource) GetRoutePath() string {
	return "/products"
}

type Product struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Name        string `json:"name"`
}

type Pagination[T any] struct {
	Results []T `json:"results"`
	Total   int `json:"total"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}
