package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type PropertyGroup struct {
	XMLName        xml.Name `xml:"PropertyGroup"`
	AssemblyName   string
	Configuration  string
	Platform       string
	PlatformTarget string
	OutputPath     string
}

func (this PropertyGroup) Get(name string) string {
	switch strings.ToLower(name) {
	case "assemblyname":
		return this.AssemblyName
	case "configuration":
		return this.Configuration
	case "platform":
		return this.Platform
	case "platformtarget":
		return this.PlatformTarget
	case "outputpath":
		return this.OutputPath
	default:
		return ""
	}
}

type xmlProject struct {
	XMLName       xml.Name `xml:"Project"`
	PropertyGroup []PropertyGroup
}

var macro = regexp.MustCompile(`\$\([^\)]+\)`)

func main1(fname string) error {
	rawdata, err2 := ioutil.ReadAll(os.Stdin)
	if err2 != nil {
		return err2
	}
	var xmldata xmlProject
	err3 := xml.Unmarshal(rawdata, &xmldata)
	if err3 != nil {
		return err3
	}
	for _, p := range xmldata.PropertyGroup {
		var buffer bytes.Buffer
		for j, arg1 := range os.Args[1:] {
			if j > 0 {
				buffer.WriteString(" ")
			}
			value := macro.ReplaceAllStringFunc(arg1, func(m string) string {
				return p.Get(m[2 : len(m)-1])
			})
			buffer.WriteString(value)
		}
		if buffer.Len() > 0 {
			fmt.Println(buffer.String())
		}
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
