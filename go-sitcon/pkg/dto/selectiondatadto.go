package dto

type SelectionDataDTO struct {
	Selections    []SelectionDTO      `json:"selections"`
	Topics        []SelectionTopicDTO `json:"topics"`
	Products      []SelectionItemDTO  `json:"products"`
	RequestId     string              `json:"requestId"`
	RequestStatus string              `json:"requestStatus"`
}
