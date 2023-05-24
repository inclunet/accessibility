package main

import (
	"flag"

	"github.com/inclunet/accessibility/pkg/checker"
)

func main() {
	var url string
	var reportPath string
	var reportPrefix string
	var lang string
	flag.StringVar(&url, "url", "https://inclunet.com.br", "Url to check accessibility")
	flag.StringVar(&reportPath, "report-path", "reports", "Output file path to generate final report files")
	flag.StringVar(&reportPrefix, "report-prefix", "pagina_", "Output file prefix to generate final report files")
	flag.StringVar(&lang, "lang", "pt-BR", "Content language of page")
	flag.Parse()
	checker.NewPageCheck(url, reportPrefix, reportPath, lang)
}
