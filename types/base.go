package odd

type Base struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Images      []*URI   `json:"images,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}
