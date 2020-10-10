package product

type CategoryList struct {
	Id   uint64 `json:"id"`
	Pid  uint64 `json:"pid"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Sort uint64 `json:"sort"`
}
