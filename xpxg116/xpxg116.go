package xpxg116

// TODO - Add comments
import "encoding/xml"

// WSAD - Openedge 11.6 WSAD
type WSAD struct {
	// XMLName xml.Name `xml:"WSAD"`
	XMLNS     string    `xml:"xmlns,attr"`
	XSI       string    `xml:"xmlns:xsi,attr"`
	AppObject AppObject // `xml:"AppObject"`
}

// PGVersion - pxpg Structure
type PGVersion struct {
	PXGVersion string // `xml:"PXGVersion"`
}

// InitPGVersion - pxpg Structure
func InitPGVersion() PGVersion {
	return PGVersion{PXGVersion: "03"}
}

// AppObject - pxpg Structure
type AppObject struct {
	// XMLName xml.Name `xml:"AppObject"`
	PGVersion               PGVersion // `xml:"PGVersion"`
	PGGenInfo               PGGenInfo `xml:"PGGenInfo"`
	DeploymentWizard        DeploymentWizard
	Name                    string
	ProPath                 string
	HelpString              string
	ODLHelpString           string
	ObjectType              bool
	AllowUnknown            bool
	IsTTResultSet           bool
	WriteDatasetBeforeImage bool
	Procedures              []Procedure `json:"Procedure" xml:"Procedure"`
}

// PGGenInfo - pxpg Structure
type PGGenInfo struct {
	XMLName                                     xml.Name `json:"-" xml:"PGGenInfo"`
	SessionFree                                 bool     `xml:"isSessionFree,attr"`
	Author                                      string
	VersionString                               string
	VersionNumber                               int
	Package                                     string
	Service                                     string `xml:"Service"`
	WorkDir                                     string
	VerboseLogging                              bool
	LeaveProxyFiles                             bool
	SaveBeforeGen                               bool
	DotNetClient                                bool
	JavaClient                                  bool
	WebServicesClient                           bool
	ESBClient                                   bool
	ESBClient2                                  bool
	SilverlightClient                           bool
	RESTClient                                  bool
	UserDefinedAppService                       bool
	JavaCompilerType                            int
	JavaCompiler                                string
	JavaSwitches                                string
	JavaClasspathSwitch                         string
	JavaClasspath                               string
	DotNetCompilerType                          int
	DotNetCompiler                              string
	DotNetSwitches                              string
	DotNetXSDGenerator                          string
	DotNetNamespace                             string
	DotNetTitle                                 string
	DotNetVersion                               string
	DotNetDesc                                  string
	DotNetCompany                               string
	DotNetProduct                               string
	DotNetPublicKey                             string
	DotNetDelaySign                             bool
	DotNetRuntime                               string
	DotNetDataSetNamespace                      string
	DNUseDefaultDSNamespace                     bool
	DNUseNamespaceAsDirs                        bool
	DotNetUseNullableTypes                      bool
	SilverlightDomainServiceNamespace           string
	SilverlightUseDefaultDomainServiceNamespace bool
	SilverlightEntityNamespace                  string
	SilverlightUseDefaultEntityNamespace        bool
}

// DeploymentWizard - pxpg Structure
type DeploymentWizard struct {
	CurrentEncoding     int
	SoapEndpointURL     string
	WebServiceNamespace WebServiceNamespace
	SoapAction          SoapAction
	ConnectReturnString ConnectReturnString
	TestWSDL            TestWSDL
	ESBEncoding         int
	EntryEndpoint       EntryEndpoint
	FaultEndpoint       FaultEndpoint
	RejectedEndpoint    RejectedEndpoint
	Container           string
	AppserverURL        string
	FileName            FileName
	ResourceDirectory   string
	SonicAppService     SonicAppService
	SonicConnectionInfo SonicConnectionInfo
}

// SonicConnectionInfo - pxpg Structure
type SonicConnectionInfo struct {
	SonicURL    string
	SonicDomain string
	User        string
	Password    string
}

// SonicAppService - pxpg Structure
/*
type SonicAppService struct {
	XMLName           xml.Name `xml:"SonicAppService"`
	ResourceDirectory ResourceDirectory
	XARNName          string
}
*/

// ResourceDirectory - pxpg Structure
type ResourceDirectory struct {
	XMLName   xml.Name `json:"-" xml:"ResourceDirectory"`
	Value     string   `xml:",chardata"`
	OverWrite bool     `xml:"overWrite,attr"`
}

// WebServiceNamespace - pxpg Structure
type WebServiceNamespace struct {
	XMLName     xml.Name `json:"-" xml:"WebServiceNamespace"`
	Value       string   `xml:",chardata"`
	UserDefined bool     `xml:"userDefined,attr"`
}

// SoapAction - pxpg Structure
type SoapAction struct {
	XMLName            xml.Name `json:"-" xml:"SoapAction"`
	Value              string   `xml:",chardata"`
	AppendToSoapAction bool     `xml:"appendToSoapAction,attr"`
}

// FileName - pxpg Structure
type FileName struct {
	XMLName   xml.Name `json:"-" xml:"FileName"`
	Value     string   `xml:",chardata"`
	OverWrite bool     `xml:"overWrite,attr"`
}

// TestWSDL - pxpg Structure
type TestWSDL struct {
	XMLName    xml.Name `json:"-" xml:"TestWSDL"`
	Value      string   `xml:",chardata"`
	BGen       bool     `xml:"bGen,attr"`
	RPCEncoded bool     `xml:",attr"`
	RPCLiteral bool     `xml:",attr"`
	DocLiteral bool     `xml:",attr"`
}

// ConnectReturnString - pxpg Structure
type ConnectReturnString struct {
	XMLName     xml.Name `json:"-" xml:"ConnectReturnString"`
	Value       string   `xml:",chardata"`
	UserDefined bool     `xml:"userDefined,attr"`
}

// EntryEndpoint - pxpg Structure
type EntryEndpoint struct {
	XMLName    xml.Name `json:"-" xml:"EntryEndpoint"`
	Value      string   `xml:",chardata"`
	UseDefault bool     `xml:"useDefault,attr"`
}

// FaultEndpoint - pxpg Structure
type FaultEndpoint struct {
	XMLName    xml.Name `json:"-" xml:"FaultEndpoint"`
	Value      string   `xml:",chardata"`
	UseDefault bool     `xml:"useDefault,attr"`
}

// RejectedEndpoint - pxpg Structure
type RejectedEndpoint struct {
	XMLName    xml.Name `json:"-" xml:"RejectedEndpoint"`
	Value      string   `xml:",chardata"`
	UseDefault bool     `xml:"useDefault,attr"`
}

// SonicAppService - pxpg Structure
type SonicAppService struct {
	UseFileSystem     bool `xml:"useFileSystem,attr"`
	DeployToDS        bool `xml:"deployToDS,attr"`
	CreateXAR         bool `xml:"createXAR,attr"`
	ResourceDirectory ResourceDirectory
	XARName           string
}

// Procedure - pxpg Structure - Used for list of programs
type Procedure struct {
	XMLName      xml.Name `json:"-" xml:"Procedure"`
	IsPersistent bool     `xml:"isPersistent,attr"`
	UseFullName  bool     `xml:"useFullName,attr"`
	Name         string
	ProcPath     string
	ProPath      string
	ProcExt      string
}
