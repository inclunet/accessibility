package main

import (
	"flag"
	"log"

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
	evaluator, err := checker.NewChecker(checker.AccessibilityChecker{
		FileName:   checkList,
		Lang:       lang,
		ReportPath: reportPath,
	})

	if err != nil {
		log.Println(err)
	}

	evaluator.GetDomainName(url)
	evaluator.AddCheckListItem(url, reportFile)
	evaluator.CheckAllList()
	evaluator.SaveAllReports()
}
