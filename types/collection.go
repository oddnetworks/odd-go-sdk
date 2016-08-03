package odd

type Collection struct {
	Base
	Entities []*Base `json:"entities"`
}
