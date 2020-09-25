package address

type AddressList struct {
	AddressId uint64 `json:"address_id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	IsDefault bool   `json:"is_default"`
}

type AddressDetail struct {
	AddressId    uint64 `json:"address_id"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	CodeProv     uint64 `json:"code_prov"`
	CodeCity     uint64 `json:"code_city"`
	CodeCoun     uint64 `json:"code_coun"`
	CodeProvName string `json:"code_prov_name"`
	CodeCityName string `json:"code_city_name"`
	CodeCounName string `json:"code_coun_name"`
	Address      string `json:"address"`
	RoomNumber   string `json:"room_number"`
	IsDefault    bool   `json:"is_default"`
}
