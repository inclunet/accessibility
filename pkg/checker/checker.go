package checker

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/accessibility"
	"github.com/inclunet/accessibility/pkg/report"
	"gopkg.in/yaml.v3"
)

type AccessibilityChecker struct {
	Date       string
	Domain     string
	Lang       string
	ReportPath string
	Reports    []report.AccessibilityReport
}

// this function store a new checklist entry for manual single checks requested be command line parameters.
// It is necessary for some situations like input checklist file is not provided by user and an example is available in main.go file.
func (c *AccessibilityChecker) AddCheckListItem(url string, reportFile string, lang string, reportPath string) {
	if url != "" && reportFile != "" {
		c.Lang = lang
		c.ReportPath = reportPath
		newReport := report.AccessibilityReport{
			Domain:     c.GetDomainName(url),
			Url:        url,
			ReportFile: reportFile,
		}
		c.Reports = append(c.Reports, newReport)
	}
}

func (c *AccessibilityChecker) Check(s *goquery.Selection, accessibilityReport report.AccessibilityReport) report.AccessibilityReport {
	s.Each(func(i int, s *goquery.Selection) {

		accessibilityCheck, err := accessibility.NewElementCheck(s, accessibilityReport.Checks)

		if err == nil {
			accessibilityReport.AddCheck(accessibilityCheck)
		}

		accessibilityReport = c.Check(s.Children(), accessibilityReport)
	})

	return accessibilityReport
}

// this function starts an evaluation process for all pages stored in AccessibilityChecker Struct matrix and store Accessibility Reports in a matrix to future save all process.
// If a single page evaluation fails, this function skips and does not store report results and continue evaluation for the rest of pages.
func (c *AccessibilityChecker) CheckAllList() {
	var err error
	for i, accessibilityReport := range c.Reports {
		accessibilityReport, err = c.CheckPage(accessibilityReport, c.Lang, c.ReportPath)

		if err != nil {
			log.Printf("Is not possible to evaluate page: %s", err)
		}

		c.Reports[i] = accessibilityReport
	}
}

// this function starts a single page evaluation process and returns an AccessibilityReport Struct results.
// if the evaluation process fails, this ffunction skip evaluation and return an error to logger.
func (c *AccessibilityChecker) CheckPage(accessibilityReport report.AccessibilityReport, lang string, reportPath string) (report.AccessibilityReport, error) {
	log.Printf("Starting page evaluation process for %s", accessibilityReport.Url)
	document, html, err := c.GetPage(accessibilityReport.Url)

	if err != nil {
		return accessibilityReport, err
	}

	accessibilityReport.Html = html
	accessibilityReport.Title = document.Find("head title").Text()
	log.Printf("Evaluating page with title: %s", accessibilityReport.Title)
	accessibilityReport.NewSummary()
	accessibilityReport = c.Check(document.Find("html"), accessibilityReport)
	log.Println("evaluation process finished.")

	return accessibilityReport, nil
}

// Get domain name from a given url to use from command line interface
func (c *AccessibilityChecker) GetDomainName(url string) string {
	if url != "" {
		c.Domain = "default"
		if domainName := strings.Split(strings.Split(url, "://")[1], "/")[0]; domainName != "" {
			c.Domain = domainName
			return domainName
		}
	}
	return ""
}

// Get a page for start the accessibility evaluation process.
// this function requires an URL and returns a goquery documment object, a html text and an error if is not possible to get a page.
func (c *AccessibilityChecker) GetPage(url string) (*goquery.Document, string, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, "", errors.New("URL not found")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, "", err
	}

	html := string(body)

	response.Body = io.NopCloser(bytes.NewBuffer(body))

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		return nil, "", err
	}

	return document, html, nil
}

// Create and check if domain report folders exists on disk.
// Returns a string containing a report path to save report files or an error if is not possible to creat directories.
func (c *AccessibilityChecker) getReportPath() (string, error) {
	reportPath := c.ReportPath + "/" + c.Domain

	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		err = os.Mkdir(reportPath, 0755)

		if err != nil {
			return "", err
		}

		err := os.Mkdir(reportPath+"/html", 0755)

		if err != nil {
			return "", err
		}

		err = os.Mkdir(reportPath+"/json", 0755)

		if err != nil {
			return "", err
		}
	}

	return reportPath, nil
}

// Load a checklist file from informed location and populate an AccessibilityChecklist struct matrix to future evaluation process.
func (c *AccessibilityChecker) LoadCheckList(fileName string) error {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &c)

	if err != nil {
		return err
	}

	return nil
}

// Save all AccessibilityReports in HTML and JSON files  in domain reports   files folder.
// the default reports root folder is "reports" and a domain folder is created automatically to store report files.
func (c *AccessibilityChecker) SaveAllReports() {
	log.Println("saving all reports...")
	var err error
	for i, accessibilityReport := range c.Reports {
		log.Printf("saving report file #%d for %s", i, accessibilityReport.ReportFile)
		accessibilityReport.HtmlReportPath, accessibilityReport.JsonReportPath, err = c.Save(accessibilityReport, accessibilityReport.ReportFile, "page-report.html")

		if err != nil {
			log.Printf("is not possible to save report data: %s", err)
		}

		c.Reports[i] = accessibilityReport
	}
}

// Save HTML and JSON report files.
// tis function requires any struct and a template filename to construct reporte files and returns an error if this operation fails.
func (c *AccessibilityChecker) Save(data any, reportFile string, templateFile string) (string, string, error) {
	reportPath, err := c.getReportPath()

	if err != nil {
		return "", "", err
	}

	htmlReportPath := reportPath + "/html/" + reportFile + ".html"

	err = c.SaveHtmlReport(data, htmlReportPath, c.ReportPath+"/templates/"+templateFile)

	if err != nil {
		return "", "", err
	}

	jsonReportPath := reportPath + "/json/" + reportFile + ".json"

	err = c.SaveJsonReport(data, jsonReportPath)

	if err != nil {
		return htmlReportPath, "", err
	}

	return htmlReportPath, jsonReportPath, nil
}

// Save a HTML file from a global or individual page report struct.
// This function Requires any struct, a path address to save html report and a template file to construct the html report files.
// if is not possible to save html report file on disk, this function returns an error to inform.
func (c *AccessibilityChecker) SaveHtmlReport(data any, fileName string, templateFileName string) error {
	File, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer File.Close()

	newTemplate := template.Must(template.ParseFiles(templateFileName))

	if err != nil {
		return err
	}

	err = newTemplate.Execute(File, data)

	if err != nil {
		return err
	}

	return nil
}

// Save a json file from a global report or a page report struct.
// This function requires a struct and a path address to save json report file and returns an error if is not possible to save a json report on disk.
func (c *AccessibilityChecker) SaveJsonReport(data any, fileName string) error {
	File, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer File.Close()

	jsonContent, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = File.Write(jsonContent)

	if err != nil {
		return err
	}

	return nil
}

// Starts a new checker object and initialize checking accessibility if the informed input checklist file is available.
// This function returns an AccessibilityChecker object and espects a yaml input file wit a checklist to a batch evaluation.
// If the imput checklist isn't informed, this functions return an AccessibilityObject but don't check any page.
func NewChecker(fileName string) AccessibilityChecker {
	newChecker := AccessibilityChecker{
		Lang:       "pt-br",
		ReportPath: "reports",
	}

	if fileName != "" {
		err := newChecker.LoadCheckList(fileName)

		if err == nil {
			newChecker.CheckAllList()
			newChecker.SaveAllReports()
		} else {
			log.Printf("skiping load check list input file: %s", err)
		}
	}

	return newChecker
}
