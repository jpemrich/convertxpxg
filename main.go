package main

import (
	"convertxpxg/xpxg116"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"net/url"
)

// This is a copy of the header from the xml package with standalone added
// TODO - is there a better way to do this.
const (
	xmlHeader = `<?xml version="1.0" encoding="utf-8" standalone="no"?>` + "\n"
)

type Config struct {
	Environments []Environment
	Plans []Plan
	Services []Service
}

type Environment struct {
	Name string
	ServiceSuffix string
	UriHost string `json:"uriHost"`
}

type Plan struct {
	Name string
	ServicePrefix string
	UriPath string `json:"uriPath"`
	EnabledServices []string
}

type Service struct {
	Name string
	ProPath string
	GeneralInfo GeneralInfo
	DeploymentInfo DeploymentInfo
	Procedures []string 
}

type GeneralInfo struct {
	SessionFree bool
	Author string
	WorkDir string
	VerboseLogging bool
	LeaveProxyFiles bool
	WebServicesClient bool

}

type DeploymentInfo struct {
	CurrentEncoding int
	WebServiceNamespaceValue string
	uriScheme string
	WebServiceNamespaceUserDefined bool
	ConnectReturnString bool
	GenerateTestWSDL bool
	WSDLRPCEncoded bool
	WSDLRPCLiteral bool
	WSDLDocLiteral bool
	AppserverURL string
	FileName string
	OverWriteFile bool
}

func pprint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {

	xmlFile, err := os.Open("s:/progress_prod/wsa/prod/hpm/Rhapsody.xpxg")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var wsad xpxg116.WSAD
	
	fmt.Println(string(byteValue))
	xml.Unmarshal(byteValue, &wsad)

	fmt.Println(pprint(wsad))

	fmt.Println(wsad.AppObject.DeploymentWizard.SoapEndpointURL)
	
	
	 
	config := new(Config)

	config.Environments = append(config.Environments, Environment{Name : "Development",
		ServiceSuffix : "devl",
		UriHost : "mimcs01-dvl.caidan.local"})

	config.Environments = append(config.Environments, Environment{Name : "QA",
		ServiceSuffix : "qa",
		UriHost : "mimcs01-qvl.caidan.local"})

	config.Environments = append(config.Environments, Environment{Name : "UAT",
		ServiceSuffix : "uat",
		UriHost : "mimcs01-uvl.caidan.local"})
	
	config.Environments = append(config.Environments, Environment{Name : "Production",
		ServiceSuffix : "prod",
		UriHost : "mcs01-ppl.caidan.local"})


	config.Plans = append(config.Plans,Plan{Name : "hpm",
		ServicePrefix : "wsHPM",
		UriPath : "/hpm/wsa1", 
		EnabledServices : []string{"Eligibility","EmployeePortal","ErAlert","GMC","Hpmich","IVR","MCS","Member","Memportaldata","Memportal","Provider","PubWS","Rhapsody","System","TaskTimer","Vendor"}})
	
	config.Plans = append(config.Plans,Plan{Name : "mhpil",
		ServicePrefix : "wsMHPIL",
		UriPath : "/mhpil/wsa1", 
		EnabledServices : []string{"Eligibility","GMC","IVR","MCS","Member","Memportaldata","Memportal","Mhplan","Provider","PubWS","Rhapsody","System","Vendor"}})

	config.Plans = append(config.Plans,Plan{Name : "mhpia",
		ServicePrefix : "wsMHPIA",
		UriPath : "/mhpia/wsa1", 
		EnabledServices : []string{"Eligibility","IVR","MCS","Member","Memportaldata","Memportal","Mhplan","Provider","PubWS","Rhapsody","System","Vendor"}})
	
	config.Plans = append(config.Plans,Plan{Name : "mhpnh",
		ServicePrefix : "wsMHPNH",
		UriPath : "/mhpnh/wsa1", 
		EnabledServices : []string{"Eligibility","IVR","MCS","Member","Memportaldata","Memportal","Mhplan","Provider","PubWS","System","Vendor"}})

	config.Plans = append(config.Plans,Plan{Name : "mhpoh",
		ServicePrefix : "wsMHPOH",
		UriPath : "/mhpoh/wsa1", 
		EnabledServices : []string{"IVR","MCS","PubWS"}})

	config.Plans = append(config.Plans,Plan{Name : "mhpdp",
		ServicePrefix : "wsMHPDP",
		UriPath : "/mhpdp/wsa1", 
		EnabledServices : []string{"GMC","Rhapsody"}})

	var s Service

	s = Service{Name : wsad.AppObject.Name}
	s.ProPath = wsad.AppObject.ProPath
	
	s.GeneralInfo.SessionFree = wsad.AppObject.PGGenInfo.SessionFree
	s.GeneralInfo.Author = wsad.AppObject.PGGenInfo.Author
	s.GeneralInfo.WorkDir = wsad.AppObject.PGGenInfo.WorkDir
	s.GeneralInfo.VerboseLogging = wsad.AppObject.PGGenInfo.VerboseLogging
	s.GeneralInfo.LeaveProxyFiles = wsad.AppObject.PGGenInfo.LeaveProxyFiles
	s.GeneralInfo.WebServicesClient = wsad.AppObject.PGGenInfo.WebServicesClient

	s.DeploymentInfo.CurrentEncoding = wsad.AppObject.DeploymentWizard.CurrentEncoding
	s.DeploymentInfo.WebServiceNamespaceValue = wsad.AppObject.DeploymentWizard.WebServiceNamespace.Value
	u, _ := url.Parse(wsad.AppObject.DeploymentWizard.SoapEndpointURL)
	s.DeploymentInfo.uriScheme = u.Scheme 
	s.DeploymentInfo.WebServiceNamespaceUserDefined = wsad.AppObject.DeploymentWizard.WebServiceNamespace.UserDefined
	s.DeploymentInfo.ConnectReturnString = wsad.AppObject.DeploymentWizard.ConnectReturnString.UserDefined
	s.DeploymentInfo.GenerateTestWSDL = wsad.AppObject.DeploymentWizard.TestWSDL.BGen
	s.DeploymentInfo.WSDLRPCEncoded = wsad.AppObject.DeploymentWizard.TestWSDL.RPCEncoded
	s.DeploymentInfo.WSDLRPCLiteral = wsad.AppObject.DeploymentWizard.TestWSDL.RPCLiteral
	s.DeploymentInfo.WSDLDocLiteral = wsad.AppObject.DeploymentWizard.TestWSDL.DocLiteral
	s.DeploymentInfo.AppserverURL = wsad.AppObject.DeploymentWizard.AppserverURL
	s.DeploymentInfo.FileName = wsad.AppObject.DeploymentWizard.FileName.Value
	s.DeploymentInfo.OverWriteFile = wsad.AppObject.DeploymentWizard.FileName.OverWrite

	config.Services = append(config.Services, s)
	for i, p := range wsad.AppObject.Procedures {
		fmt.Println(i, p.Name, p.ProPath, p.ProcExt)
	}
	
/* 	s = Service {Name : "GMC"}
	s.Name = "GMC"
	s.ProPath = "S:/progress/edi/;S:/progress/mcs/"
	config.Services = append(config.Services, s)
*/

	// fmt.Println(pprint(config))
	
}
