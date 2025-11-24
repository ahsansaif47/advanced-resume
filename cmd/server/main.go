package main

import (
	"encoding/json"
	"fmt"

	"github.com/ahsansaif47/advanced-resume/integrations/gemini"
	"github.com/ahsansaif47/advanced-resume/internal/parser"
)

func main() {
	client, err := gemini.GenAIClient()
	if err != nil {
		return
	}

	res, err := gemini.GetResponse(client, "/home/ahsansaif/projects/advanced-resume/resources/images/AhsanResume202507/page18.png")
	if err != nil {
		return
	}

	fmt.Println(res)

	cleanedData := parser.CleanJSON(res)

	data, err := parser.ParseResume([]byte(cleanedData))
	if err != nil {
		return
	}

	if _, err := json.MarshalIndent(data, "", " "); err != nil {
		return
	}

	fmt.Printf("%v", data)

	// client := parser.InitClient()
	// text, err := parser.GetText(client, "/home/ahsansaif/projects/advanced-resume/resources/images/AhsanResume202507/page-0.png")
	// if err != nil {
	// 	return
	// }

	// fmt.Println(text)

}
