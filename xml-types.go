package main

import (
	"encoding/xml"
)

type config struct {
	XMLName           xml.Name          `xml:"configuration"`
	ConfigSections    configSections    `xml:"configSections"`
	AppSettings       appSettings       `xml:"appSettings"`
	ConnectionStrings connectionStrings `xml:"connectionStrings"`
}

type configSections struct {
	XMLName       xml.Name       `xml:"configSections"`
	Sections      []section      `xml:"section"`
	SectionGroups []sectionGroup `xml:"sectionGroup"`
}

type section struct {
	XMLName           xml.Name `xml:"section"`
	Name              string   `xml:"name,attr"`
	Type              string   `xml:"type,attr"`
	RequirePermission bool     `xml:"requirePermission,attr,omitempty"`
}

type sectionGroup struct {
	XMLName           xml.Name  `xml:"sectionGroup"`
	Name              string    `xml:"name,attr"`
	Type              string    `xml:"type,attr,omitempty"`
	RequirePermission bool      `xml:"requirePermission,attr,omitempty"`
	Section           []section `xml:"section"`
}

type appSettings struct {
	XMLName xml.Name `xml:"appSettings"`
	Add     []add    `xml:"add"`
}

type add struct {
	XMLName xml.Name `xml:"add"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

type connectionStrings struct {
	XMLName xml.Name                  `xml:"connectionStrings"`
	Add     []connectionStringsAdd    `xml:"add"`
	Remove  []connectionStringsRemove `xml:"remove"`
}

type connectionStringsAdd struct {
	XMLName          xml.Name `xml:"add"`
	Name             string   `xml:"name,attr"`
	ConnectionString string   `xml:"connectionString,attr"`
	ProviderName     string   `xml:"providerName,attr,omitempty"`
}

type connectionStringsRemove struct {
	XMLName xml.Name `xml:"remove"`
	Name    string   `xml:"name,attr"`
}
