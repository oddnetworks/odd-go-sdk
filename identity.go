package odd

// Channel represents a channel
type Channel struct {
	Identification
	Base
	display
}

// Platform represents a platform
type Platform struct {
	Identification
	Base
	display
	Category []byte
}

// Viewer represents a viewer
type Viewer struct {
	Identification
	Entitlements [][]byte
}

type display struct {
	Images []*Image
	Colors []*Color
}
