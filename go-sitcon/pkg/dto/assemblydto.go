package dto

type AssemblyDTO struct {
	ContainerSize int    `json:"containerSize"`
	ItemCount     int    `json:"itemCount"`
	Type          string `json:"type,omitempty"`
}
