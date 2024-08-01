package dtos

import (
	"log"
	"strconv"

	"github.com/vijishetty29/Go/prepim-amap-api/models"
)

type SelectionDTO struct {
	Branch               string         `json:"branch"`
	Code                 string         `json:"code"`
	Active               string         `json:"active"`
	AdvertisingStartWeek YearAndWeekDTO `json:"advertisingStartWeek"`
	AdvertisingEndWeek   YearAndWeekDTO `json:"advertisingEndWeek"`
	DeliveryStartWeek    YearAndWeekDTO `json:"deliveryStartWeek"`
	DeliveryEndWeek      YearAndWeekDTO `json:"deliveryEndWeek"`
}

func CreateSelectionDTOResponse(selection models.Selection) SelectionDTO {
	return SelectionDTO{
		Branch:               "NON-FOOD",
		Code:                 selection.Code,
		Active:               GetActiveStatus(selection.Active),
		AdvertisingStartWeek: GetYearAndWeekDTO(selection.AdvertisingStartWeek, selection.AdvertisingStartYear),
		AdvertisingEndWeek:   GetYearAndWeekDTO(selection.AdvertisingEndWeek, selection.AdvertisingEndYear),
		DeliveryStartWeek:    GetYearAndWeekDTO(selection.DeliveryStartWeek, selection.DeliveryStartYear),
		DeliveryEndWeek:      GetYearAndWeekDTO(selection.DeliveryEndWeek, selection.DeliveryEndYear),
	}
}

type SelectionsResponse []*SelectionDTO

func CreateSelectionsResponse(selections []*models.Selection) SelectionsResponse {
	selectionsResp := SelectionsResponse{}
	for _, s := range selections {
		selection := CreateSelectionDTOResponse(*s)
		selectionsResp = append(selectionsResp, &selection)
	}
	return selectionsResp
}

func GetActiveStatus(active string) string {
	str, err := strconv.ParseBool(active)
	if err != nil {
		log.Fatal("Error while parsing boolean value: ", active)
		return "false"
	}
	return strconv.FormatBool(str)
}
