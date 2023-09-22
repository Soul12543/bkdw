package bkdw

import (
	"fmt"
	//	"io/ioutil"

	//	"os"

	"golang.org/x/net/html"
)

func main() {
	id := "128396"
	body, _ := GetBookInfo(id)

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
