package dto

type SelectionTopicDTO struct {
	Code             string              `json:"code"`
	Status           bool                `json:"status"`
	TopicNr          string              `json:"topicNr,omitempty"`
	Branch           string              `json:"branch"`
	Selection        string              `json:"selection"`
	Description      *LocalizedStringDTO `json:"description,omitempty"`
	AdvertisingWeeks []YearAndWeek       `json:"advertisingWeeks,omitempty"`
}
