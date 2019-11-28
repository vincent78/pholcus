package clib

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
	"testing"
)

func TestParseHtml(t *testing.T) {
	file, err := os.Open("/Users/vincent/temp/xici.nn.html")
	if err != nil {
		println(err.Error())
		return
	}
	doc, _ := goquery.NewDocumentFromReader(file)
	doc.Find("#ip_list tr").Each(func(i int, tr *goquery.Selection) {
		if i > 0 {
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				println("index:", j, "\ttd:", strings.TrimSpace(td.Text()))

			})
		}
	})
}
