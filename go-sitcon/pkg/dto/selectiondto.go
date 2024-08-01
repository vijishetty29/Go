package dto

type SelectionDTO struct {
	Code                 string       `json:"code"`
	DeliveryEndWeek      *YearAndWeek `json:"deliveryEndWeek,omitempty"`
	DeliveryStartWeek    *YearAndWeek `json:"deliveryStartWeek,omitempty"`
	AdvertisingEndWeek   *YearAndWeek `json:"advertisingEndWeek,omitempty"`
	AdvertisingStartWeek *YearAndWeek `json:"advertisingStartWeek,omitempty"`
	Branch               string       `json:"branch"`
	Active               bool         `json:"active"`
}
