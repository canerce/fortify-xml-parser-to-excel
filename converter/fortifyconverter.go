package converter

import (
	"fmt"

	"github.com/bigbird023/fortify-xml-parser-to-excel/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx"
)

type (
	FortifyConverter struct {
		header     []string
		inputFile  string
		outputFile string
		fortifyxml xml.FortifyXMLInterface
	}
)

func NewConverter(input string, output string, fxp xml.FortifyXMLInterface) *FortifyConverter {
	conv := &FortifyConverter{
		header:     []string{"Iid", "RuleId", "Category", "Folder", "Kingdom", "Abstract", "Friority", "Primary.Filename", "Primary.FilePath", "Primary.LineStart", "Primary.Snippet", "Primary.TargetFunction", "Source.Filename", "Source.FilePath", "Source.LineStart", "Source.Snippet", "Source.TargetFunction"},
		inputFile:  input,
		outputFile: output,
		fortifyxml: fxp,
	}
	return conv
}

func (c *FortifyConverter) Convert() error {
	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	reportDefinition, err := c.fortifyxml.XmlParse(c.inputFile)
	if err != nil {
		return err
	}

	c.headerToExcel(sheet)
	if err != nil {
		return err
	}

	for rsloop := 0; rsloop < len(reportDefinition.ReportSection); rsloop++ {
		for ssloop := 0; ssloop < len(reportDefinition.ReportSection[rsloop].SubSection); ssloop++ {
			for gsloop := 0; gsloop < len(reportDefinition.ReportSection[rsloop].SubSection[ssloop].IssueListing.Chart.GroupingSection); gsloop++ {
				for issueloop := 0; issueloop < len(reportDefinition.ReportSection[rsloop].SubSection[ssloop].IssueListing.Chart.GroupingSection[gsloop].Issue); issueloop++ {
					err := c.issueToExcel(&reportDefinition.ReportSection[rsloop].SubSection[ssloop].IssueListing.Chart.GroupingSection[gsloop].Issue[issueloop], sheet)
					if err != nil {
						err = fmt.Errorf("error converting issue to excel")
						return err
					}
				}
			}
		}
	}

	err = c.writeExcelToFile(excelFile)
	if err != nil {
		return err
	}

	return nil
}

func (c *FortifyConverter) headerToExcel(sheet xlsx.Sheet) {

	row := sheet.Row(0)

	headers := [17]string{"Iid", "RuleId", "Category", "Folder", "Kingdom", "Abstract", "Friority", "Primary.Filename", "Primary.FilePath", "Primary.LineStart", "Primary.Snippet", "Primary.TargetFunction", "Source.Filename", "Source.FilePath", "Source.LineStart", "Source.Snippet", "Source.TargetFunction"}

	for p, v := range headers {
		cell := row.Cell(p)
		cell.SetValue(v)
	}
}

func (c *FortifyConverter) issueToExcel(issue *xml.Issue, sheet xlsx.Sheet) error {

	_, totalRows := sheet.Dimension()
	row := sheet.Row(totalRows - 1)

	if row.Cell(0).Value() != "" {
		//if headers are set, move down
		row = sheet.Row(totalRows)
	}

	col := -1
	c.setNextCell(&col, row, issue.Iid)
	c.setNextCell(&col, row, issue.RuleID)
	c.setNextCell(&col, row, issue.Category)
	c.setNextCell(&col, row, issue.Folder)
	c.setNextCell(&col, row, issue.Kingdom)
	c.setNextCell(&col, row, issue.Abstract)
	c.setNextCell(&col, row, issue.Friority)
	c.setNextCell(&col, row, issue.Primary.FileName)
	c.setNextCell(&col, row, issue.Primary.FilePath)
	c.setNextCell(&col, row, issue.Primary.LineStart)
	c.setNextCell(&col, row, issue.Primary.Snippet)
	c.setNextCell(&col, row, issue.Primary.TargetFunction)
	c.setNextCell(&col, row, issue.Source.FileName)
	c.setNextCell(&col, row, issue.Source.FilePath)
	c.setNextCell(&col, row, issue.Source.LineStart)
	c.setNextCell(&col, row, issue.Source.Snippet)
	c.setNextCell(&col, row, issue.Source.TargetFunction)

	return nil
}

func (c *FortifyConverter) setNextCell(colNumber *int, row *xlsx.Row, value string) {
	*colNumber++
	cell := row.Cell(*colNumber)
	cell.SetValue(value)
}

func (c *FortifyConverter) writeExcelToFile(excelFile ooxml.Package) error {

	// Save the XLSX file under different name
	err := excelFile.SaveAs(c.outputFile)
	if err != nil {
		return err
	}

	return nil
}
