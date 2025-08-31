package models

type GetDeliveryPricesResponse struct {
	Data []DeliveryPrice `json:"data"`
}

type DeliveryPrice struct {
	Date          string  `json:"date"`
	DeliveryPrice float64 `json:"delivery_price"`
}
