package dto

type YearAndWeek struct {
	Year int `json:"year"`
	Week int `json:"week"`
	ID   int `json:"id,omitempty"`
}
