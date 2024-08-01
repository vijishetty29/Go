package dto

type SelectionItemDTO struct {
	Code                     string             `json:"code,omitempty"`
	SelectionItemID          string             `json:"selectionItemId,omitempty"`
	Name                     LocalizedStringDTO `json:"name,omitempty"`
	Topic                    string             `json:"topic,omitempty"`
	Status                   ItemStatusDTO      `json:"status,omitempty"`
	PurchasingRemark         string             `json:"purchasingRemark,omitempty"`
	OrderableCountries       []string           `json:"orderableCountries,omitempty"`
	DeliveryDates            []YearAndWeek      `json:"deliveryDates,omitempty"`
	Brand                    BrandDTO           `json:"brand,omitempty"`
	Branch                   string             `json:"branch,omitempty"`
	ProductCategory          string             `json:"productCategory,omitempty"`
	Buyer                    BuyerDTO           `json:"buyer,omitempty"`
	Characteristics          CharacteristicsDTO `json:"characteristics,omitempty"`
	DisplayType              DisplayTypeDTO     `json:"displayType,omitempty"`
	ItemGroup                ItemGroupDTO       `json:"itemGroup,omitempty"`
	Assembly                 AssemblyDTO        `json:"assembly,omitempty"`
	ItemType                 string             `json:"itemType,omitempty"`
	AdditionalDescriptionInt string             `json:"additionalDescriptionInt"`
	LinkedItem               string             `json:"linkedItem,omitempty"`

	//ReferencePredecessor PredecessorDTO   `json:"referencePredecessor"`
	// Predecessors         []PredecessorDTO `json:"predecessors"`
}
