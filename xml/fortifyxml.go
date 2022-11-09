/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package xml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type (
	FortifyXMLInterface interface {
		XmlParse(inputFile string) (*ReportDefinition, error)
	}

	FortifyXml struct {
	}

	ReportDefinition struct {
		XMLName       xml.Name `xml:"ReportDefinition"`
		Text          string   `xml:",chardata"`
		Type          string   `xml:"type,attr"`
		TemplateName  string   `xml:"TemplateName"`
		TemplatePath  string   `xml:"TemplatePath"`
		LogoPath      string   `xml:"LogoPath"`
		Footnote      string   `xml:"Footnote"`
		UserName      string   `xml:"UserName"`
		ReportSection []struct {
			Text                string `xml:",chardata"`
			Enabled             string `xml:"enabled,attr"`
			OptionalSubsections string `xml:"optionalSubsections,attr"`
			Title               string `xml:"Title"`
			SubSection          []struct {
				Chardata     string `xml:",chardata"`
				Enabled      string `xml:"enabled,attr"`
				Title        string `xml:"Title"`
				Description  string `xml:"Description"`
				Text         string `xml:"Text"`
				IssueListing struct {
					Text       string `xml:",chardata"`
					Listing    string `xml:"listing,attr"`
					Limit      string `xml:"limit,attr"`
					Refinement string `xml:"Refinement"`
					Chart      struct {
						Text            string `xml:",chardata"`
						ChartType       string `xml:"chartType,attr"`
						Axis            string `xml:"Axis"`
						MajorAttribute  string `xml:"MajorAttribute"`
						GroupingSection []struct {
							Text                  string `xml:",chardata"`
							Count                 string `xml:"count,attr"`
							GroupTitle            string `xml:"groupTitle"`
							MajorAttributeSummary struct {
								Text     string `xml:",chardata"`
								MetaInfo []struct {
									Text  string `xml:",chardata"`
									Name  string `xml:"Name"`
									Value string `xml:"Value"`
								} `xml:"MetaInfo"`
								AttributeValue []struct {
									Text  string `xml:",chardata"`
									Name  string `xml:"Name"`
									Count string `xml:"Count"`
								} `xml:"AttributeValue"`
							} `xml:"MajorAttributeSummary"`
							Issue []Issue `xml:"Issue"`
						} `xml:"GroupingSection"`
					} `xml:"Chart"`
				} `xml:"IssueListing"`
			} `xml:"SubSection"`
		} `xml:"ReportSection"`
	}

	Issue struct {
		Text     string   `xml:",chardata"`
		Iid      string   `xml:"iid,attr"`
		RuleID   string   `xml:"ruleID,attr"`
		Category string   `xml:"Category"`
		Folder   string   `xml:"Folder"`
		Kingdom  string   `xml:"Kingdom"`
		Abstract string   `xml:"Abstract"`
		Friority string   `xml:"Friority"`
		Primary  CodeInfo `xml:"Primary"`
		Source   CodeInfo `xml:"Source"`
	}

	CodeInfo struct {
		Text           string `xml:",chardata"`
		FileName       string `xml:"FileName"`
		FilePath       string `xml:"FilePath"`
		LineStart      string `xml:"LineStart"`
		Snippet        string `xml:"Snippet"`
		TargetFunction string `xml:"TargetFunction"`
	}
	Source struct {
		Text           string `xml:",chardata"`
		FileName       string `xml:"FileName"`
		FilePath       string `xml:"FilePath"`
		LineStart      string `xml:"LineStart"`
		Snippet        string `xml:"Snippet"`
		TargetFunction string `xml:"TargetFunction"`
	}
)

func NewFortifyXmlParser() FortifyXMLInterface {
	return &FortifyXml{}
}

func (f *FortifyXml) XmlParse(inputFile string) (*ReportDefinition, error) {

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
	var reportDefinition ReportDefinition
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &reportDefinition)

	return &reportDefinition, nil

}
