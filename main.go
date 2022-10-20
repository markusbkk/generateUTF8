package main

import (
	"fmt"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// main prints all UTF 8 runes starting from 21 to d088
// See: https://www.fileformat.info/info/charset/UTF-8/list.htm
func main() {
	// string buffer
	var sb strings.Builder
	// hex 21
	begin := 33
	// hex d088
	end := 90
	for i := begin; i <= end; i++ {
		sb.WriteString(string(i))
	}
	generatePDF(sb.String())
}

func generatePDF(text string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("arialunicode", "", "ArialUnicodeMS.ttf")
	pdf.SetFont("arialunicode", "", 14)
	pdf.MoveTo(0, 10)
	pdf.Cell(1, 1, text)
	err := pdf.OutputFileAndClose("utf8.pdf")
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println("PDF mit UTF Zeichensatz erfolgreich generiert")
	}
}
