package pdfreader

import (
	"bytes"
	libpdf "github.com/ledongthuc/pdf"
)

type Pdf struct {
	Path string
}

func (readPdf *Pdf) Read() (string, error) {
	f, r, err := libpdf.Open(readPdf.Path)

	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)

	return buf.String(), nil
}
