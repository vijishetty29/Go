package dtos

type YearAndWeekDTO struct {
	Week string `json:"week"`
	Year string `json:"year"`
}

func GetYearAndWeekDTO(advertisingStartWeek, advertisingStartYear string) YearAndWeekDTO {
	return YearAndWeekDTO{
		Week: advertisingStartWeek,
		Year: advertisingStartYear,
	}
}
