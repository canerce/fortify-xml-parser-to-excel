package parser

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/bigbird023/fortify-xml-parser-to-excel/data"
)

type (
	//FortifyXMLParseInterface interface for fortifyxml
	FortifyXMLParseInterface interface {
		XMLParse(inputFile string) (*data.ReportDefinition, error)
	}

	//FortifyXMLParse - object for fortifyxmlparsing
	FortifyXMLParse struct {
	}
)

//NewFortifyXMLParser create new object
func NewFortifyXMLParser() FortifyXMLParseInterface {
	return &FortifyXMLParse{}
}

//XMLParse parsing xml into go struct
func (f *FortifyXMLParse) XMLParse(inputFile string) (*data.ReportDefinition, error) {

	// Open xmlFile
	xmlFile, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xmlFile)

	// we initialize our Users array
	var reportDefinition data.ReportDefinition
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &reportDefinition)

	return &reportDefinition, nil

}
