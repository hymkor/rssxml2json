package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Item struct {
	XMLName xml.Name `xml:"item" json:"-"`
	Link    string   `xml:"link" json:"link"`
	PubDate string   `xml:"pubDate",json:"pubDate"`
	Desc    string   `xml:"description",json:"description"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss" json:"-"`
	Item    []Item   `xml:"channel>item"`
}

func mains() error {
	var rss Rss

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(in, &rss)
	if err != nil {
		return err
	}
	// fmt.Fprintf(os.Stderr, "%d records\n", len(rss.Item))
	out, err := json.Marshal(&rss.Item)
	if err != nil {
		return err
	}
	os.Stdout.Write(out[:])
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
