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

func read(r io.Reader) ([]Item, error) {
	var rss Rss

	in, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(in, &rss)
	return rss.Item, err
}

func mains() error {
	var items []Item
	if len(os.Args) > 1 {
		for _, fn := range os.Args[1:] {
			in, err := os.Open(fn)
			if err != nil {
				return err
			}
			_items, err := read(in)
			in.Close()
			if err != nil {
				return err
			}
			items = append(items, _items...)
		}
	} else {
		var err error
		items, err = read(os.Stdin)
		if err != nil {
			return err
		}
	}
	// fmt.Fprintf(os.Stderr, "%d records\n", len(items))
	out, err := json.Marshal(&items)
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
