package webframe

import (
	"errors"

	"github.com/crgimenes/glaze"
)

// ErrSnapshotUnsupported is returned when the platform's WebView backend cannot
// rasterize the page (today: everything except macOS).
var ErrSnapshotUnsupported = errors.New("webframe: this backend cannot snapshot the page")

// Snapshot returns the current page as PNG bytes.
//
// The Frame keeps its glaze.WebView unexported, so this is the only way out to
// a capture; it type-asserts the optional Snapshotter interface so backends
// without one keep working. Call it from a goroutine that is not the UI thread:
// the capture completes on the main queue, which Run is pumping.
func (f *Frame) Snapshot() ([]byte, error) {
	if f == nil || f.w == nil {
		return nil, ErrSnapshotUnsupported
	}
	s, ok := f.w.(glaze.Snapshotter)
	if !ok {
		return nil, ErrSnapshotUnsupported
	}
	return s.Snapshot()
}
