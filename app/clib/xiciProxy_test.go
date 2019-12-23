package clib

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestParseHtml(t *testing.T) {
	file, err := os.Open("/Users/vincent/temp/pholcus/xici.nn.html")
	if err != nil {
		println(err.Error())
		return
	}
	doc, _ := goquery.NewDocumentFromReader(file)
	doc.Find("#ip_list tr").Each(func(i int, tr *goquery.Selection) {
		if i > 0 {
			obj := make(map[string]string)
			//println("No. ", i)
			obj["index"] = strconv.Itoa(i)
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				//println("index:", j, "\ttd:", strings.TrimSpace(td.Text()))
				switch j {
				case 1:
					obj["ip"] = strings.TrimSpace(td.Text())
				case 2:
					obj["port"] = strings.TrimSpace(td.Text())
				case 3:
					obj["region"] = strings.TrimSpace(td.Text())
				case 4:
					obj["category"] = strings.TrimSpace(td.Text())
				case 5:
					obj["type"] = strings.TrimSpace(td.Text())
				}
			})
			println(transfer(obj))
		}
	})
}

func transfer(m map[string]string) string {
	mjson, _ := json.Marshal(m)
	return string(mjson)
}
