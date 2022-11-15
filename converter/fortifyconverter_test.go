package converter

import (
	"os"
	"testing"

	"github.com/bigbird023/fortify-xml-parser-to-excel/data"
	"github.com/bigbird023/fortify-xml-parser-to-excel/mock"
	"github.com/bigbird023/fortify-xml-parser-to-excel/parser"
	"github.com/plandem/xlsx"
)

func compareNextColValue(t *testing.T, colNumber *int, row *xlsx.Row, expectedValue string) {
	*colNumber++
	cell := row.Cell(*colNumber)
	if cell.Value() != expectedValue {
		t.Logf("Expected value %s does not equal actual value %s", expectedValue, cell.Value())
		t.Fail()
	}
}

func TestHeaderToExcel(t *testing.T) {

	c := NewConverter("", "", nil)

	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	c.headerToExcel(sheet)

	cols, rows := sheet.Dimension()
	if cols != len(c.header) {
		t.Log("17 cols are expected")
		t.Fail()
	}

	if rows != 1 {
		t.Log("1 cols are expected")
		t.Fail()
	}

	row := sheet.Row(0)
	for cols := 0; cols < len(c.header); cols++ {
		if row.Cell(cols).Value() != c.header[cols] {
			t.Logf("Header Column %d should be %s", cols, c.header[cols])
			t.Fail()
		}
	}

}

func TestIssueToExcel(t *testing.T) {

	c := NewConverter("", "", nil)

	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	issue := newTestFortifyIssue()

	c.issueToExcel(issue, sheet)

	cols, rows := sheet.Dimension()
	if cols != len(c.header) {
		t.Log("17 cols are expected")
		t.Fail()
	}

	if rows != 1 {
		t.Log("1 rows are expected")
		t.Fail()
	}
	row := sheet.Row(rows - 1)

	assertRowToIssue(t, row, issue)
}

func TestHeaderAndIssueToExcel(t *testing.T) {

	c := NewConverter("", "", nil)

	excelFile := xlsx.New()
	sheet := excelFile.AddSheet("fortifyIssues")

	issue := newTestFortifyIssue()

	c.headerToExcel(sheet)

	err := c.issueToExcel(issue, sheet)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	cols, rows := sheet.Dimension()
	if cols != len(c.header) {
		t.Log("17 cols are expected")
		t.Fail()
	}

	if rows != 2 {
		t.Log("1 rows are expected")
		t.Fail()
	}
	row := sheet.Row(rows - 1)

	assertRowToIssue(t, row, issue)
}

func TestWriteToExcel(t *testing.T) {

	expected := "./local/output/testwritetoexcel.xlsx"
	fxp := parser.NewFortifyXmlParser()

	c := NewConverter("", expected, fxp)

	excelFile := mock.NewMockPackage()

	err := c.writeExcelToFile(excelFile)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if excelFile.SaveAsInterface != expected {
		t.Logf("excel saveas interface was not set correct expected %s, actual %s", expected, excelFile.SaveAsInterface)
		t.Fail()
	}

}

func TestWriteToExcelError(t *testing.T) {

	expected := "./local/output/testwritetoexcel.xlsx"
	fxp := parser.NewFortifyXmlParser()

	c := NewConverter("", expected, fxp)

	excelFile := mock.NewMockPackage()
	excelFile.ForceError = true

	err := c.writeExcelToFile(excelFile)
	if err == nil {
		t.Log("error not thrown")
		t.Fail()
	}

}

func TestConvertInputError(t *testing.T) {

	expectedInput := "./local/test/testmissing.xml"
	expectedOutput := "./local/output/testwritetoexcel.xlsx"
	fxp := parser.NewFortifyXmlParser()

	c := NewConverter(expectedInput, expectedOutput, fxp)

	err := c.Convert()
	if err == nil {
		t.Log("Error expected, missing")
		t.Fail()
	}

}

func TestConvertOutputError(t *testing.T) {
	expectedOutput := "../local/output2/TestConvert.xlsx"

	//verify TestConvert.xlsx doesn't exist, delete if does
	_, err := os.Stat(expectedOutput)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Log(err)
			t.Fail()
		}
	} else {
		err = os.Remove(expectedOutput)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}

	fxp := mock.NewMockFortifyXMLParser()

	c := NewConverter("", expectedOutput, fxp)

	err = c.Convert()
	if err == nil {
		t.Log("Should have generated error")
		t.Fail()
	}

}

func TestConvertIssueToExcelError(t *testing.T) {
	expectedOutput := "../local/output2/TestConvert.xlsx"

	//verify TestConvert.xlsx doesn't exist, delete if does
	_, err := os.Stat(expectedOutput)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Log(err)
			t.Fail()
		}
	} else {
		err = os.Remove(expectedOutput)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}

	fxp := mock.NewMockFortifyXMLParser()
	fxp.EmptyReportDefinition = true

	c := NewConverter("", expectedOutput, fxp)

	err = c.Convert()
	if err == nil {
		t.Log("Should have generated error")
		t.Fail()
	}

}

func newTestFortifyIssue() *data.Issue {
	return &data.Issue{
		Iid:      "iid",
		RuleID:   "ruleid",
		Category: "category",
		Folder:   "folder",
		Kingdom:  "kingdom",
		Abstract: "abstract",
		Friority: "friority",
		Primary: data.CodeInfo{
			FileName:       "filename",
			FilePath:       "filepath",
			LineStart:      "linestart",
			Snippet:        "snippet",
			TargetFunction: "targetfunction",
		},
		Source: data.CodeInfo{
			FileName:       "filename",
			FilePath:       "filepath",
			LineStart:      "linestart",
			Snippet:        "snippet",
			TargetFunction: "targetfunction",
		},
	}
}

func assertRowToIssue(t *testing.T, row *xlsx.Row, issue *data.Issue) {
	col := -1
	compareNextColValue(t, &col, row, issue.Iid)
	compareNextColValue(t, &col, row, issue.RuleID)
	compareNextColValue(t, &col, row, issue.Category)
	compareNextColValue(t, &col, row, issue.Folder)
	compareNextColValue(t, &col, row, issue.Kingdom)
	compareNextColValue(t, &col, row, issue.Abstract)
	compareNextColValue(t, &col, row, issue.Friority)
	compareNextColValue(t, &col, row, issue.Primary.FileName)
	compareNextColValue(t, &col, row, issue.Primary.FilePath)
	compareNextColValue(t, &col, row, issue.Primary.LineStart)
	compareNextColValue(t, &col, row, issue.Primary.Snippet)
	compareNextColValue(t, &col, row, issue.Primary.TargetFunction)
	compareNextColValue(t, &col, row, issue.Source.FileName)
	compareNextColValue(t, &col, row, issue.Source.FilePath)
	compareNextColValue(t, &col, row, issue.Source.LineStart)
	compareNextColValue(t, &col, row, issue.Source.Snippet)
	compareNextColValue(t, &col, row, issue.Source.TargetFunction)
}
