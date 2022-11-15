package parser

import (
	"testing"
)

func TestXmlParseBadFile(t *testing.T) {

	cfgInput := "../local/examples/testBadFile.xml"

	fxp := NewFortifyXmlParser()

	_, err := fxp.XmlParse(cfgInput)
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
	fxp := NewFortifyXmlParser()

	reportDef, err := fxp.XmlParse(cfgInput)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if reportDef == nil {
		t.Log("Return should not be nil")
		t.Fail()
	}

	if len(reportDef.ReportSection) != 0 {
		t.Log("ReportSection length should be 0")
		t.Fail()
	}

}

func TestXmlParse(t *testing.T) {

	cfgInput := "../local/examples/test.xml"
	fxp := NewFortifyXmlParser()

	reportDef, err := fxp.XmlParse(cfgInput)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if reportDef == nil {
		t.Log("Return should not be nil")
		t.Fail()
	}

	if len(reportDef.ReportSection) == 0 {
		t.Log("ReportSection length should not be 0")
		t.Fail()
	}
}
