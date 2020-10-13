package product

type Product struct {
	Id           int     `json: "id"`
	ProductCode  string  `json: "productConde"`
	ProductName  string  `json: "productName"`
	Description  string  `json: "description"`
	StandardCost float64 `json: "standardCost"`
	ListPrice    float64 `json: "listPrice"`
	Category     string  `json: "Category"`
}
