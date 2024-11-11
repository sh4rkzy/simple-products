package repository

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	DtCreated string  `json:"dt_created"`
	DtUpdated string  `json:"dt_updated"`
}

func GetProducts() {

}

func GetProductById(id string) {

}

func CreateProduct(product Product) {

}

func UpdateProduct(product Product) {

}

func DeleteProduct(id string) {

}
