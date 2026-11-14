package webframe

import (
	"errors"

	"github.com/dolfly/glaze"
)

// ErrPrintUnsupported is returned when the platform's WebView backend cannot
// print the page to PDF (today: everything except macOS).
var ErrPrintUnsupported = errors.New("webframe: this backend cannot print the page to PDF")

// PrintPDF returns the page as a vector PDF.
//
// Like Snapshot this exists because the Frame keeps its glaze.WebView
// unexported; it type-asserts glaze's optional PDFPrinter so backends without
// one keep working. The document covers the whole scrollable page, not just the
// visible viewport, and printing takes no pixels, so a hidden page produces one
// just as well as a visible one.
//
// Call it from a goroutine that is not the UI thread: the print completes on the
// main queue, which Run is pumping.
func (f *Frame) PrintPDF() ([]byte, error) {
	if f == nil || f.w == nil {
		return nil, ErrPrintUnsupported
	}
	p, ok := f.w.(glaze.PDFPrinter)
	if !ok {
		return nil, ErrPrintUnsupported
	}
	doc, err := p.PrintPDF()
	if errors.Is(err, glaze.ErrPrintUnsupported) {
		return nil, ErrPrintUnsupported
	}
	return doc, err
}
