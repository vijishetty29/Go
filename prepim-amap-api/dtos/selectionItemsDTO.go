package dtos

import (
	"github.com/vijishetty29/Go/prepim-amap-api/models"
)

type SelectionItemDTO struct {
	Branch           string `json:"branch"`
	Code             string `json:"code"`
	SelectionItemId  string `json:"selectionItemId"`
	Topic            string `json:"topic"`
	PurchasingRemark string `json:"purchasingRemark"`
}

func CreateSelectionItemDTOResponse(selectionItem models.SelectionItem) SelectionItemDTO {
	return SelectionItemDTO{
		Branch:           "NON-FOOD",
		Code:             selectionItem.Code,
		SelectionItemId:  selectionItem.Code,
		Topic:            selectionItem.TopicCode,
		PurchasingRemark: selectionItem.Note,
	}
}

type SelectionItemsResponse []*SelectionItemDTO

func CreateSelectionItemsResponse(selectionItems []*models.SelectionItem) SelectionItemsResponse {
	selectionItemsResp := SelectionItemsResponse{}
	for _, s := range selectionItems {
		selectionItem := CreateSelectionItemDTOResponse(*s)
		selectionItemsResp = append(selectionItemsResp, &selectionItem)
	}
	return selectionItemsResp
}
