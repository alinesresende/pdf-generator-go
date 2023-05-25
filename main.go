package main

import (
	"fmt"
	"os"
	"pdf-generator-go/html"
	"pdf-generator-go/pdf"
)

type Data struct {
	Name string
}

func main() {
	h := html.New("temp")
	wk := pdf.NewWkHtmlToPDF("temp")

	dataHTML := Data{
		Name: "Aline",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(htmlGenerated)
	fmt.Println("HTML gerado", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("PDF gerado", filePDFName)
}
