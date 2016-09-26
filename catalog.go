package odd

// Collection represents the container for holding other collections and vidoes
type Collection struct {
	Identification
	Base
	MediaBase
	Entities []*Identification
	Featured []*Identification
}

// Video represets a video resource
type Video struct {
	Identification
	Base
	MediaBase
	Duration int
	Cast     []*Cast
	Sources  []*Source
	Related  []*Identification
}

// View represets a view resource
type View struct {
	Identification
	Base
	Images []*Image
}

// Promotion represets a promotion resource
type Promotion struct {
	Identification
	Base
	Images []*Image
	URL    []byte
}
