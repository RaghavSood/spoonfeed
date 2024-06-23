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
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}
