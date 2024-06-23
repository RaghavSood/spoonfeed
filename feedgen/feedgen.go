package feedgen

import (
	"encoding/xml"

	"github.com/RaghavSood/spoonfeed/rss2"
)

type FeedType string

const (
	RSS2  FeedType = "rss2"
	ATOM  FeedType = "atom"
	JSON1 FeedType = "json1"
)

type FeedGenerator struct {
	Type FeedType
}

type Feed struct {
	EncodedContent []byte
	CanonicalURL   string
	Type           FeedType
}

func NewFeedGenerator(feedtype FeedType) *FeedGenerator {
	return &FeedGenerator{
		Type: feedtype,
	}
}

func (f *FeedGenerator) Generate() (Feed, error) {
	switch f.Type {
	case RSS2:
		return f.generateRSS2()
	default:
		return Feed{}, nil
	}
}

func (f *FeedGenerator) generateRSS2() (Feed, error) {
	rssFeed := rss2.RssFeedTop{
		Version: "2.0",
	}

	marshalledFeed, err := xml.Marshal(rssFeed)
	if err != nil {
		return Feed{}, err
	}

	feed := Feed{
		EncodedContent: marshalledFeed,
		CanonicalURL:   "http://example.com/rss2",
		Type:           RSS2,
	}

	return feed, nil
}
