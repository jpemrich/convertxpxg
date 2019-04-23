package main

import (
	"convertxpxg/xpxg116"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// This is a copy of the header from the xml package with standalone added
// TODO - is there a better way to do this.
const (
	xmlHeader = `<?xml version="1.0" encoding="utf-8" standalone="no"?>` + "\n"
)

func pprint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {

	xmlFile, err := os.Open("data/Rhapsody.xpxg")

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
}
