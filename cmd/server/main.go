package main

import (
	"github.com/ahsansaif47/advanced-resume/internal/parser"
)

func main() {
	// newParser, err := parser.NewFitzParser("/home/ahsansaif/Downloads/AhsanResume202507.pdf")
	// if err != nil {
	// 	return
	// }

	// err = newParser.ExtractAndSaveImages()
	// if err != nil {
	// 	log.Fatalf("Encountered error: %v", err)
	// }

	client := parser.InitClient()
	text, err := parser.GetText("/home/ahsansaif/Downloads/AhsanResume202507.pdf")

}
