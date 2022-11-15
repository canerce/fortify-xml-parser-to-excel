package data

import "encoding/xml"

type (
	//ReportDefinition main xml tag struct for storing
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

	//Issue struct for storing issues inside xml document
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

	//CodeInfo struct for storing issues inside xml document
	CodeInfo struct {
		Text           string `xml:",chardata"`
		FileName       string `xml:"FileName"`
		FilePath       string `xml:"FilePath"`
		LineStart      string `xml:"LineStart"`
		Snippet        string `xml:"Snippet"`
		TargetFunction string `xml:"TargetFunction"`
	}
	//Source struct for storing issues inside xml document
	Source struct {
		Text           string `xml:",chardata"`
		FileName       string `xml:"FileName"`
		FilePath       string `xml:"FilePath"`
		LineStart      string `xml:"LineStart"`
		Snippet        string `xml:"Snippet"`
		TargetFunction string `xml:"TargetFunction"`
	}
)
