package rss2

import "encoding/xml"

type RssFeedTop struct {
	XMLName    xml.Name `xml:"rss"`
	Version    string   `xml:"version,attr"`
	Content    string   `xml:"xmlns:content,attr,omitempty"`
	DublinCore string   `xml:"xmlns:dc,attr,omitempty"`
}
