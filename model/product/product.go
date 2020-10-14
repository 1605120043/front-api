package product

type CategoryList struct {
	Id   uint64 `json:"id"`
	Pid  uint64 `json:"pid"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Sort uint64 `json:"sort"`
}

type TagList struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type ProductList struct {
	ProductId        uint64  `json:"product_id"`
	Name             string  `json:"name"`
	Image            string  `json:"image"`
	Price            float64 `json:"price"`
	ShortDescription string  `json:"short_description"`
}
