Param
(
		[Parameter(Mandatory=$true, Position=0)]
		[string] $fprfile
)

$fprfolder = Split-Path $fprfile -Parent
$fprfilename = Split-Path $fprfile -LeafBase
$xmlfilepath = Join-Path $fprfolder ($fprfilename + ".xml")
$xlsxfilepath = Join-Path $fprfolder ($fprfilename + ".xlsx")

$XMLGenerator = "C:\Program Files\Fortify\OpenText_Application_Security_Tools_25.2.0\bin\ReportGenerator.bat"
& $XMLGenerator -source $fprfile `
			   -format xml `
			   -f $xmlfilepath `
			   -template ".\fullreport.xml" `
			   -filterSet "Security Auditor View"

$XLSXGenerator = ".\fortifyxmlparsertoexcel.exe"
& $XLSXGenerator --input $xmlfilepath --output $xlsxfilepath