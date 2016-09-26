package odd

import (
	"time"
)

// Identification represents common properties for every Odd resource with an ID and Type
type Identification struct {
	ID   []byte
	Type []byte
}

// Base represents common properties for resources with a Title, Description, and Active
type Base struct {
	Title       []byte
	Description []byte
	Active      bool
}

// MediaBase represents common properties for collection and video resources
type MediaBase struct {
	Images      []*Image
	Genres      [][]byte
	ReleaseDate time.Time
}

// Image represents external images
type Image struct {
	URL      []byte
	MimeType []byte
	Label    []byte
	Width    int
	Height   int
}

// Source represents external video sources
type Source struct {
	URL        []byte
	MimeType   []byte
	Label      []byte
	Width      int
	Height     int
	Container  []byte
	MaxBitRate int
}

// Cast represents a person that is related to a video
type Cast struct {
	Name      []byte
	Role      []byte
	Character []byte
}

// Color represents a color in RGBA format
type Color struct {
	Red   int
	Greed int
	Blue  int
	Alpha int
	Label int
}
