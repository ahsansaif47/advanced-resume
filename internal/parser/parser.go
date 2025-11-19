package parser

import (
	"bytes"
	"image/png"

	"github.com/gen2brain/go-fitz"
)

type Parser interface {
	ExtractImages() ([]ImageData, error)
	Close() error
}

type ImageData struct {
	page int
	Data []byte
}

type FitzParser struct {
	Doc *fitz.Document
}

func NewFitzParser(path string) (*FitzParser, error) {
	doc, err := fitz.New(path)
	if err != nil {
		return nil, err
	}
	return &FitzParser{Doc: doc}, nil
}

func (p *FitzParser) ExtractImages() ([]ImageData, error) {
	totalPages := p.Doc.NumPage()

	outputImages := make([]ImageData, totalPages)
	for i := 0; i < totalPages; i++ {
		img, err := p.Doc.Image(i)
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		outputImages = append(outputImages, ImageData{
			page: i,
			Data: buf.Bytes(),
		})

	}

	return outputImages, nil
}

func (p *FitzParser) Close() error {
	return p.Doc.Close()
}
