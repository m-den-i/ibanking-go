package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type stringline []string

func main() {
	f, err := os.Open("report.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rio := bufio.NewReader(f)
	doc, err := goquery.NewDocumentFromReader(rio)
	if err != nil {
		panic(err)
	}
	var x []stringline
	doc.Find(".section_3 tr").Each(func(i int, s *goquery.Selection) {
		var t []string
		for _, l := range strings.Split(s.Text(), "\n") {
			if l != "" {
				t = append(t, l)
			}
		}
		if len(t) > 0 {
			x = append(x, t)
		}
	})
	for _, l := range x[0:20] {
		fmt.Println(len(l), l)
	}
}
