package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

type xMLSession struct {
	XMLName xml.Name `xml:"Session"`
	Key     string   `xml:"Key"`
	Count   int      `xml:"Count"`
}

type xMLQRZDatabase struct {
	XMLName  xml.Name     `xml:"QRZDatabase"`
	Sessions []xMLSession `xml:"Session"`
}

func readKey(reader io.Reader) ([]xMLSession, error) {
	var xmldata xMLQRZDatabase
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	decoder.Decode(&xmldata)

	return xmldata.Sessions, nil
}

func main() {
	//Read xml file
	xmlFile, err := os.Open("test.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	xmldata, err := readKey(xmlFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	fmt.Println(xmldata[0].Key)
}
