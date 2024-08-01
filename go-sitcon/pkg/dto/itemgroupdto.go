package dto

type ItemGroupDTO struct {
	CombinedGroup string `json:"combinedGroup,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	FamilyName    string `json:"familyName,omitempty"`
	Pillar        string `json:"pillar,omitempty"`
}
