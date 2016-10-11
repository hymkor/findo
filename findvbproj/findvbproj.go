package main

import (
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

func replace1(xmldata xmlProject, lines []string) ([]string, bool) {
	replaced := false
	result := make([]string, 0, len(lines)*2)
	for _, line1 := range lines {
		loc := macro.FindStringIndex(line1)
		if loc == nil {
			result = append(result, line1)
			continue
		}
		name := line1[loc[0]+2 : loc[1]-1]
		found := false
		for _, p := range xmldata.PropertyGroup {
			value := p.Get(name)
			if value != "" {
				newline := line1[0:loc[0]] + value + line1[loc[1]:]
				result = append(result, newline)
				found = true
				replaced = true
			}
		}
		if !found {
			newline := line1[0:loc[0]] + line1[loc[1]:]
			result = append(result, newline)
		}
	}
	return result, replaced
}

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
	lines := []string{strings.Join(os.Args[1:], " ")}
	replaced := true
	for replaced {
		lines, replaced = replace1(xmldata, lines)
	}
	for _, line1 := range lines {
		fmt.Println(line1)
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
