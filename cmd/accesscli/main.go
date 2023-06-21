package main

import (
	"flag"

	"github.com/inclunet/accessibility/pkg/checker"
	"github.com/inclunet/accessibility/pkg/report"
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

	accessibilityReport := report.AccessibilityReport{
		Url:        url,
		ReportFile: reportFile,
	}

	evaluator := checker.AccessibilityChecker{
		FileName:   checkList,
		Lang:       lang,
		ReportPath: reportPath,
	}

	checker.NewChecker(evaluator, accessibilityReport)
}
