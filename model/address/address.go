package address

type AddressList struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	IsDefault bool   `json:"is_default"`
}

type AddressDetail struct {
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	IsDefault bool   `json:"is_default"`
}
