package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type PropertyGroup struct {
	XMLName        xml.Name `xml:"PropertyGroup"`
	AssemblyName   string
	Configuration  string
	Platform       string
	PlatformTarget string
	OutputPath     string
}

type xmlProject struct {
	XMLName       xml.Name `xml:"Project"`
	PropertyGroup []PropertyGroup
}

func main1(fname string) error {
	fd, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer fd.Close()
	rawdata, err2 := ioutil.ReadAll(fd)
	if err2 != nil {
		return err2
	}
	var xmldata xmlProject
	err3 := xml.Unmarshal(rawdata, &xmldata)
	if err3 != nil {
		return err3
	}
	for i, p := range xmldata.PropertyGroup {
		if i > 0 {
			fmt.Println()
		}
		fmt.Printf("AssemblyName=%s\n", p.AssemblyName)
		fmt.Printf("Configuration=%s\n", p.Configuration)
		fmt.Printf("Platform=%s\n", p.Platform)
		fmt.Printf("PlatformTarget=%s\n", p.PlatformTarget)
		fmt.Printf("OutputPath=%s\n", p.OutputPath)
	}
	return nil
}

func main() {
	for _, arg1 := range os.Args[1:] {
		if err := main1(arg1); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
