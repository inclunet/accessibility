package main

import (
	"flag"

	"github.com/inclunet/accessibility/pkg/checker"
)

func main() {
	var url string
	var reportPath string
	var reportFile string
	var lang string
	var checkList string
	flag.StringVar(&url, "url", "", "Url to check accessibility")
	flag.StringVar(&reportPath, "report-path", "reports", "Output file path to generate final report files")
	flag.StringVar(&reportFile, "report", "pagina", "Output file prefix to generate final report files")
	flag.StringVar(&lang, "lang", "pt-BR", "Content language of page")
	flag.StringVar(&checkList, "check-list", "", "List of pages to check")
	flag.Parse()
	evaluator := checker.NewChecker(checkList)
	evaluator.GetDomainName(url)
	evaluator.Lang = lang
	evaluator.ReportPath = reportPath
	evaluator.AddCheckListItem(url, reportFile, lang, reportPath)
	evaluator.CheckAllList()
	evaluator.SaveAllReports()
}
