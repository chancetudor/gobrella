package destinations

type Destination struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
	Type        string `json:"type"`
	Comment     string `json:"comment"`
	CreatedAt   string `json:"createdAt"`
}

func NewDestination() *Destination {
	return &Destination{
		ID:          "",
		Destination: "",
		Type:        "",
		Comment:     "",
		CreatedAt:   "",
	}
}
