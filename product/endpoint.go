package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string `json: "category"`
	Description  string `json: "description"`
	ListPrice    string `json: "listPrice"`
	StandardCost string `json: "standardCost"`
	ProductName  string `json: "productName"`
	ProductCode  string `json: "productConde"`
}

type updateProductRequest struct {
	Id           int64   `json: "id"`
	Category     string  `json: "category"`
	Description  string  `json: "description"`
	ListPrice    float32 `json: "listPrice"`
	StandardCost float32 `json: "standardCost"`
	ProductName  string  `json: "productName"`
	ProductCode  string  `json: "productConde"`
}

type deleteProductRequest struct {
	ProductID string
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)

		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIdEndPoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}

func makeAddProductEndPoint(s Service) endpoint.Endpoint {
	addProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)

		if err != nil {
			panic(err)
		}

		return productId, nil
	}

	return addProductEndPoint
}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		r, err := s.UpdateProduct(&req)

		if err != nil {
			panic(err)
		}

		return r, nil
	}

	return updateProductEndPoint
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	deleteProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		r, err := s.DeleteProduct(&req)

		if err != nil {
			panic(err)
		}

		return r, nil
	}

	return deleteProductEndpoint
}
