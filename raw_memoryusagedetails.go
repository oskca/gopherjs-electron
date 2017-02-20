package electron

import "github.com/gopherjs/gopherjs/js"

// MemoryUsageDetails a Structure
type MemoryUsageDetails struct {
	*js.Object
	Count         float64 `js:"count"`
	Size          float64 `js:"size"`
	LiveSize      float64 `js:"liveSize"`
	DecodedSize   float64 `js:"decodedSize"`
	PurgedSize    float64 `js:"purgedSize"`
	PurgeableSize float64 `js:"purgeableSize"`
}
