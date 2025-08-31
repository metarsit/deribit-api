package models

type GetDeliveryPricesParams struct {
	IndexName string `json:"index_name"`
	Offset    int    `json:"offset,omitempty"`
	Count     int    `json:"count,omitempty"`
}
