/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package parser

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/bigbird023/fortify-xml-parser-to-excel/data"
)

type (
	FortifyXMLInterface interface {
		XmlParse(inputFile string) (*data.ReportDefinition, error)
	}

	FortifyXml struct {
	}
)

func NewFortifyXmlParser() FortifyXMLInterface {
	return &FortifyXml{}
}

func (f *FortifyXml) XmlParse(inputFile string) (*data.ReportDefinition, error) {

	// Open xmlFile
	xmlFile, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var reportDefinition data.ReportDefinition
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &reportDefinition)

	return &reportDefinition, nil

}
