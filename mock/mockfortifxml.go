package mock

import (
	exml "encoding/xml"
	"fmt"

	"github.com/bigbird023/fortify-xml-parser-to-excel/xml"
)

type MockFortifyXml struct {
	EmptyReportDefinition bool
	ForceError            bool
}

func NewMockFortifyXML() *MockFortifyXml {
	return &MockFortifyXml{
		EmptyReportDefinition: false,
		ForceError:            false,
	}
}

func (m *MockFortifyXml) XmlParse(inputFile string) (*xml.ReportDefinition, error) {
	if m.ForceError {
		return nil, fmt.Errorf("forced error")
	}
	var rd *xml.ReportDefinition
	if m.EmptyReportDefinition {
		rd = &xml.ReportDefinition{}
	} else {
		rd = NewReportDefinition()
	}

	return rd, nil
}

func NewReportDefinition() *xml.ReportDefinition {

	xmlDoc := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<ReportDefinition type="standard">
		<TemplateName>Fortify Developer Workbook</TemplateName>
		<TemplatePath></TemplatePath>
		<LogoPath>/MF_logo.jpg</LogoPath>
		<Footnote>Copyright 2018 Micro Focus or one of its affiliates.</Footnote>
		<UserName></UserName>
		<ReportSection enabled="true" optionalSubsections="false">
			<Title>Report Overview</Title>
			<SubSection enabled="true">
				<Title>Report Summary</Title>
				<Description>This provides a high level summary of the findings that the analysis produced.  Also includes basic information on the scope of the scan.</Description>
				<Text>On Sep 13, 2022, a source code review was performed over the UT_JAVA_Adhoc_Sep_13_2022_15_16 code base. 1,697 files, 182,494 LOC (Executable) were scanned. A total of 14119 issues were uncovered during the analysis.  This report provides a comprehensive description of all the types of issues found in this project.  Specific examples and source code are provided for each issue type.</Text>
			</SubSection>
			<SubSection enabled="true">
				<Title>Issue Summary by Fortify Priority Order</Title>
				<Description>A table summarizing the number of issues found and the breakdown of issues in each Fortify Priority Level</Description>
				<IssueListing listing="false" limit="-1">
					<Refinement></Refinement>
					<Chart chartType="table">
						<Axis>Fortify Priority Order</Axis>
						<MajorAttribute>Analysis</MajorAttribute>
						<GroupingSection count="10154">
							<groupTitle>Low</groupTitle>
						</GroupingSection>
						<GroupingSection count="3500">
							<groupTitle>High</groupTitle>
						</GroupingSection>
						<GroupingSection count="354">
							<groupTitle>Medium</groupTitle>
						</GroupingSection>
						<GroupingSection count="111">
							<groupTitle>Critical</groupTitle>
						</GroupingSection>
					</Chart>
				</IssueListing>
			</SubSection>
		</ReportSection>
		<ReportSection enabled="true" optionalSubsections="false">
			<Title>Issue Summary</Title>
			<SubSection enabled="true">
				<Title>Overall number of results</Title>
				<Description>Results count</Description>
				<Text>The scan found 14119 issues.</Text>
			</SubSection>
			<SubSection enabled="true">
				<Title>Issues By Category</Title>
				<IssueListing listing="false" limit="-1">
					<Refinement></Refinement>
					<Chart chartType="table">
						<Axis>Category</Axis>
						<MajorAttribute>Analysis</MajorAttribute>
						<GroupingSection count="3238">
							<groupTitle>Trust Boundary Violation</groupTitle>
						</GroupingSection>
						<GroupingSection count="2241">
							<groupTitle>Access Control: Database</groupTitle>
						</GroupingSection>
						<GroupingSection count="1503">
							<groupTitle>System Information Leak: HTML Comment in JSP</groupTitle>
						</GroupingSection>
						<GroupingSection count="796">
							<groupTitle>Denial of Service: Parse Double</groupTitle>
						</GroupingSection>
					</Chart>
				</IssueListing>
			</SubSection>
		</ReportSection>
		<ReportSection enabled="true" optionalSubsections="true">
			<Title>Results Outline</Title>
			<SubSection enabled="true">
				<Title>Vulnerability Examples by Category</Title>
				<Description>Results summary of all issue categories.  Vulnerability examples are provided by category.</Description>
				<IssueListing listing="true" limit="5">
					<Refinement></Refinement>
					<Chart chartType="list">
						<Axis>Category</Axis>
						<MajorAttribute>Analysis</MajorAttribute>
						<GroupingSection count="3238">
							<groupTitle>Trust Boundary Violation</groupTitle>
							<MajorAttributeSummary>
								<MetaInfo>
	<Name>Abstract</Name>
	<Value>The method setDisplayMsg() in SecurityUtil.java commingles trusted and untrusted data in the same data structure, which encourages programmers to mistakenly trust unvalidated data.</Value>
								</MetaInfo>
								<MetaInfo>
	<Name>Explanation</Name>
	<Value>A trust boundary can be thought of as line drawn through a program. On one side of the line, data is untrusted. On the other side of the line, data is assumed to be trustworthy. The purpose of validation logic is to allow data to safely cross the trust boundary--to move from untrusted to trusted.
	
	A trust boundary violation occurs when a program blurs the line between what is trusted and what is untrusted. The most common way to make this mistake is to allow trusted and untrusted data to commingle in the same data structure.
	
	
	
	Example: The following Java code accepts an HTTP request and stores the usrname parameter in the HTTP session object before checking to ensure that the user has been authenticated.
	
	
	usrname = request.getParameter("usrname");
	if (session.getAttribute(ATTR_USR) != null) {
		session.setAttribute(ATTR_USR, usrname);
	}
	
	
	Without well-established and maintained trust boundaries, programmers will inevitably lose track of which pieces of data have been validated and which have not. This confusion eventually allows some data to be used without first being validated.</Value>
								</MetaInfo>
								<MetaInfo>
	<Name>Recommendations</Name>
	<Value>Define clear trust boundaries in the application. Do not use the same data structure to hold trusted data in some contexts and untrusted data in other contexts. Minimize the number of ways that data can move across a trust boundary.
	
	Trust boundary violations sometimes occur when input needs to be built up over a series of user interactions before being processed. It may not be possible to do complete input validation until all of the data has arrived. In these situations, it is still important to maintain a trust boundary. The untrusted data should be built up in a single untrusted data structure, validated, and then moved into a trusted location.</Value>
								</MetaInfo>
								<MetaInfo>
	<Name>Tips</Name>
	<Value>1. Do not feel that you need to find a "smoking gun" situation in which data that has not been validated is assumed to be trusted. If trust boundaries are not clearly delineated and respected, validation errors are inevitable. Instead of spending time searching for an exploitable scenario, concentrate on teaching programmers to create good trust boundaries.
	
	2. Most programs have trust boundaries that are defined by the semantics of the application. Consider writing custom rules to check for other places where user input crosses a trust boundary.
	
	3. A number of modern web frameworks provide mechanisms to perform user input validation (including Struts and Spring MVC). To highlight the unvalidated sources of input, the Fortify Secure Coding Rulepacks dynamically re-prioritize the issues reported by Fortify Static Code Analyzer by lowering their probability of exploit and providing pointers to the supporting evidence whenever the framework validation mechanism is in use. We refer to this feature as Context-Sensitive Ranking. To further assist the Fortify user with the auditing process, the Fortify Software Security Research group makes available the Data Validation project template that groups the issues into folders based on the validation mechanism applied to their source of input.</Value>
								</MetaInfo>
								<AttributeValue>
	<Name>&lt;Unaudited&gt;</Name>
	<Count>3238</Count>
								</AttributeValue>
								<AttributeValue>
	<Name>Not an Issue</Name>
	<Count>0</Count>
								</AttributeValue>
								<AttributeValue>
	<Name>Reliability Issue</Name>
	<Count>0</Count>
								</AttributeValue>
								<AttributeValue>
	<Name>Bad Practice</Name>
	<Count>0</Count>
								</AttributeValue>
								<AttributeValue>
	<Name>Suspicious</Name>
	<Count>0</Count>
								</AttributeValue>
								<AttributeValue>
	<Name>Exploitable</Name>
	<Count>0</Count>
								</AttributeValue>
							</MajorAttributeSummary>
							<Issue iid="79D46549F7258F4FE0E7545EB991C9B4" ruleID="CBDB6290-DF73-42E1-8D9E-3B5C4B629761">
								<Category>Trust Boundary Violation</Category>
								<Folder>Low</Folder>
								<Kingdom>Encapsulation</Kingdom>
								<Abstract>The method setDisplayMsg() in SecurityUtil.java commingles trusted and untrusted data in the same data structure, which encourages programmers to mistakenly trust unvalidated data.</Abstract>
								<Friority>Low</Friority>
								<Primary>
	<FileName>SecurityUtil.java</FileName>
	<FilePath>CNSI-JAR/src/com/cnsi/security/common/SecurityUtil.java</FilePath>
	<LineStart>411</LineStart>
	<Snippet>        else
				b_objDispMsgBean.addMsgDetail(p_sMsgCode, p_arMsgParams);
			p_objRequest.setAttribute(CommonConstants.ATTR_APP_MSG_BEAN, b_objDispMsgBean);
		}
	</Snippet>
	<TargetFunction>javax.servlet.ServletRequest.setAttribute()</TargetFunction>
								</Primary>
								<Source>
	<FileName>RequestSessionUtil.java</FileName>
	<FilePath>CNSI-JAR/src/com/cnsi/common/util/RequestSessionUtil.java</FilePath>
	<LineStart>249</LineStart>
	<Snippet>        try
			{
				l_sGridName = p_reqRequest.getParameter(GRID_NAME);
				l_enmParameterNames = p_reqRequest.getParameterNames();
				boolean isFld=false;</Snippet>
	<TargetFunction>javax.servlet.ServletRequest.getParameter()</TargetFunction>
								</Source>
							</Issue>
							<Issue iid="50260DAAC3FA7AB42D9C4BD5C6376DF1" ruleID="CBDB6290-DF73-42E1-8D9E-3B5C4B629761">
								<Category>Trust Boundary Violation</Category>
								<Folder>Low</Folder>
								<Kingdom>Encapsulation</Kingdom>
								<Abstract>The method setDisplayMsg() in SecurityUtil.java commingles trusted and untrusted data in the same data structure, which encourages programmers to mistakenly trust unvalidated data.</Abstract>
								<Friority>Low</Friority>
								<Primary>
	<FileName>SecurityUtil.java</FileName>
	<FilePath>CNSI-JAR/src/com/cnsi/security/common/SecurityUtil.java</FilePath>
	<LineStart>411</LineStart>
	<Snippet>        else
				b_objDispMsgBean.addMsgDetail(p_sMsgCode, p_arMsgParams);
			p_objRequest.setAttribute(CommonConstants.ATTR_APP_MSG_BEAN, b_objDispMsgBean);
		}
	</Snippet>
	<TargetFunction>javax.servlet.ServletRequest.setAttribute()</TargetFunction>
								</Primary>
								<Source>
	<FileName>EHRRequestHandler.java</FileName>
	<FilePath>EHR-JAR/src/com/ehr/web/action/EHRRequestHandler.java</FilePath>
	<LineStart>50</LineStart>
	<Snippet>    	
			try {
				String m_sUserLoginID = p_objRequest.getParameter("fhdn&amp;UserId");
				String m_sUserDomainID = p_objRequest.getParameter("fhdn&amp;DomainName");
				String m_sUserProfileId = p_objRequest.getParameter("fhdn&amp;ProfileId");</Snippet>
	<TargetFunction>javax.servlet.ServletRequest.getParameter()</TargetFunction>
								</Source>
							</Issue>
						</GroupingSection>
					</Chart>
				</IssueListing>
			</SubSection>
		</ReportSection>
	</ReportDefinition>
	`

	byteValue := []byte(xmlDoc)

	var reportDefinition *xml.ReportDefinition

	exml.Unmarshal(byteValue, &reportDefinition)

	return reportDefinition
}
