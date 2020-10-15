package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s), getProductByIDRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByHandler)

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(makeAddProductEndPoint(s), addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(makeUpdateProductEndPoint(s), updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	deleteProductHandler := kithttp.NewServer(makeDeleteProductEndPoint(s), deleteProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	return r
}

func getProductByIDRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productId,
	}, nil
}

func getProductsRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}

func addProductRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}

func updateProductRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}

func deleteProductRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return deleteProductRequest{
		ProductID: chi.URLParam(r, "id"),
	}, nil
}
