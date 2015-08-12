package main

import (
	"fmt"
	"log"
    "flag"
    "os"
    "net/http"
	"io/ioutil"
    "encoding/xml"

	"code.google.com/p/gcfg"
)

type Config struct {
	Section struct {
		Username string
		Password string
	}
}

type Key struct {
	XMLName xml.Name `xml:"Key"`
	token string `xml:",chardata"`
}


type Session struct {
	XMLName xml.Name `xml:"Session"`
	key Key `xml:"Key"`
}
type QRZDatabase struct {
	XMLName xml.Name `xml:"QRZDatabase"`
	sessions Session `xml:"Session"`
}


var cfg Config

func Readconfig(filename string) {
	err := gcfg.ReadFileInto(&cfg, filename)
	if err != nil {
		log.Fatal(err)
	}
}

func Getdetails(fileName string, callSign string) {
	url := "http://xmldata.qrz.com/xml/current/?"
	Readconfig(fileName)
	u := cfg.Section.Username
	p := cfg.Section.Password
	fmt.Println(u)
	xmlContent, err := http.Get(url + "username=" + u +";password=" + p + ";agent=q5.0")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(xmlContent.Body)
	if err != nil {
		log.Fatal(err)
	}
	xmlData := QRZDatabase{}
	fmt.Println(string(b))
    xml.Unmarshal(b, xmlData)
    fmt.Println(xmlData)

	fmt.Println(xmlData.sessions.key.token)

}

func main() {
    clArg := flag.String("configFile", "", "Location of configuration")
	lookupCallSign := flag.String("callSign", "", "Call Sign to be searched")
	flag.Parse()
	if (*clArg == "") || (*lookupCallSign == "") {
		fmt.Println("Usage: go run GoCQCQCQ.go -configFile </path/of/settings> --callSign <callsign>")
		os.Exit(1)
	}
	Getdetails(*clArg, *lookupCallSign)
}
