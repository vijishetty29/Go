package dtos

import (
	"github.com/vijishetty29/Go/prepim-amap-api/models"
)

type SelectionTopicDTO struct {
	Branch      string `json:"branch"`
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
	Selection   string `json:"selection"`
	Status      string `json:"status,omitempty"`
	TopicNr     string `json:"topicNr,omitempty"`
}

func CreateSelectionTopicDTOResponse(selectionTopic models.SelectionTopic) SelectionTopicDTO {
	return SelectionTopicDTO{
		Branch:    "NON-FOOD",
		Code:      selectionTopic.Code,
		Selection: selectionTopic.Selection.Code,
		Status:    selectionTopic.Status.Code,
		TopicNr:   selectionTopic.TopicNr,
	}
}

type SelectionTopicsResponse []*SelectionTopicDTO

func CreateSelectionTopicsResponse(selectionTopics []*models.SelectionTopic) SelectionTopicsResponse {
	selectionTopicsResp := SelectionTopicsResponse{}
	for _, st := range selectionTopics {
		topic := CreateSelectionTopicDTOResponse(*st)
		selectionTopicsResp = append(selectionTopicsResp, &topic)
	}
	return selectionTopicsResp
}
