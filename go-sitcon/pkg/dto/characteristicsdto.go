package dto

type CharacteristicsDTO struct {
	RepeatingArticle  bool   `json:"repeatingArticle"`
	DisplayArticle    bool   `json:"displayArticle"`
	InnovationArticle bool   `json:"innovationArticle"`
	Classification    string `json:"classification,omitempty"`
	BriefingHL        bool   `json:"briefingHL"`
	CatalogArticle    bool   `json:"catalogArticle"`
	CountryProposal   string `json:"countryProposal,omitempty"`
}
