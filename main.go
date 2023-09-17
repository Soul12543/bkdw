package main

import (
	"fmt"
	//	"io/ioutil"
	"net/http"
	//	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

func main() {
	url := "https://www.52bqg.org/book_128396/"
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("[ERROR]:%s\n", err.Error())
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("[ERROR]:%s\n", err.Error())
	}
	defer resp.Body.Close()

	body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Printf("[ERROR]:%s\n", err.Error())
	}

	doc, err := html.Parse(body)
	if err != nil {
		fmt.Printf("[ERROR]:%s", err.Error())
	}
	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Printf("%s ", attr.Val)
				}
			}
			fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	extractLinks(doc)
}
