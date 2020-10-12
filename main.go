package main

import (
	"InventarioGO/database"
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"

	"net/http"

	"github.com/go-chi./chi"
)

var databaseConnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_id"`
	Description  string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	const query = `SELECT id, product_code, COALESCE(description, '')
				   FROM products`
	results, err := databaseConnection.Query(query)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}

	respondwithJSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("Insert products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Successfully created"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "applicatio/json")
	w.WriteHeader(code)
	w.Write(response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var producto Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("Update products set product_code =?, description=? where id=?")
	catch(err)

	_, er := query.Exec(producto.Product_Code, producto.Description, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Update successfully"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("delete from products where id=?")
	catch(err)

	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Successfully deleted"})

}

func main() {

	databaseConnection = database.DbConnection()
	defer databaseConnection.Close()

	r := chi.NewRouter()
	r.Get("/products", AllProducts)
	r.Post("/products", CreateProduct)
	r.Put("/products/{id}", UpdateProduct)
	r.Delete("/products/{id}", DeleteProduct)

	http.ListenAndServe(":3000", r)

}
