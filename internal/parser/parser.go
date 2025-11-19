package parser

import (
	"fmt"
	"image/png"
	"os"

	"github.com/gen2brain/go-fitz"
)

type Parser interface {
	ExtractImages() error
	Close() error
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

func (p *FitzParser) ExtractAndSaveImages() error {
	totalPages := p.Doc.NumPage()

	outputDir := "/home/ahsansaif/projects/advanced-resume/resources/images"

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {

	}

	for i := 0; i < totalPages; i++ {
		img, err := p.Doc.Image(i)
		if err != nil {
			return err
		}

		filePath := fmt.Sprintf("%s-page-%d.png", outputDir, i)
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}

		err = png.Encode(f, img)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *FitzParser) Close() error {
	return p.Doc.Close()
}
