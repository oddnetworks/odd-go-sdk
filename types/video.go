package odd

type Video struct {
	Entity
	Streams []*Stream `json:"streams"`
}
