package main

import (
	"flag"

	"github.com/inclunet/accessibility/pkg/accessibility"
)

func main() {
	var url string
	var report string
	var lang string
	flag.StringVar(&url, "url", "https://inclunet.com.br", "Url to check accessibility")
	flag.StringVar(&report, "report", "report.html", "Output filename to generate final html report")
	flag.StringVar(&lang, "lang", "pt-BR", "Content language of page")
	flag.Parse()
	accessibility.EvaluatePage(url, report, lang)
}
