package rss2

import "encoding/xml"

type RssFeedTop struct {
	XMLName    xml.Name   `xml:"rss"`
	Version    string     `xml:"version,attr"`
	Content    string     `xml:"xmlns:content,attr,omitempty"`
	DublinCore string     `xml:"xmlns:dc,attr,omitempty"`
	Channel    RssChannel `xml:"channel"`
}

type RssChannel struct {
	// Required fields
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`

	// Optional fields
	Categories []Category `xml:"category,omitempty"`
	Cloud      Cloud      `xml:"cloud,omitempty"`
	Copyright  string     `xml:"copyright,omitempty"`
	Docs       string     `xml:"docs,omitempty"`
	Generator  string     `xml:"generator,omitempty"`
}

type Category struct {
	Domain string `xml:"domain,attr,omitempty"`
	Value  string `xml:",chardata"`
}

type Cloud struct {
	Domain   string `xml:"domain,attr"`
	Port     int    `xml:"port,attr"`
	Path     string `xml:"path,attr"`
	Register string `xml:"registerProcedure,attr"`
	Protocol string `xml:"protocol,attr"`
}
