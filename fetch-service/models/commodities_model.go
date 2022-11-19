package models

type Commodities struct {
	Uuid         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUsd     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}
