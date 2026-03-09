package animal

type Animal struct {
	ID     string  `json:"id"`
	Type   string  `json:"type"`
	Gender string  `json:"gender"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}
