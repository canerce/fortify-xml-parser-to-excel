package parser

import (
	"testing"
)

func TestXmlParseBadFile(t *testing.T) {

	cfgInput := "../local/examples/testBadFile.xml"

	fxp := NewFortifyXMLParser()

	_, err := fxp.XMLParse(cfgInput)
	if err != nil {
		if err.Error() != "open ../local/examples/testBadFile.xml: no such file or directory" {
			t.Log(err)
			t.Fail()
		}
	} else {
		t.Log("Error expected")
		t.Fail()
	}

}

func TestXmlParseEmpty(t *testing.T) {

	cfgInput := "../local/examples/testEmpty.xml"
	fxp := NewFortifyXMLParser()

	reportDef, err := fxp.XMLParse(cfgInput)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if reportDef == nil {
		t.Log("Return should not be nil")
		t.Fail()
	}

	if reportDef != nil && len(reportDef.ReportSection) != 0 {
		t.Log("ReportSection length should be 0")
		t.Fail()
	}

}

func TestXmlParse(t *testing.T) {

	cfgInput := "../local/examples/test.xml"
	fxp := NewFortifyXMLParser()

	reportDef, err := fxp.XMLParse(cfgInput)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if reportDef == nil {
		t.Log("Return should not be nil")
		t.Fail()
	}

	if reportDef != nil && len(reportDef.ReportSection) == 0 {
		t.Log("ReportSection length should not be 0")
		t.Fail()
	}
}
