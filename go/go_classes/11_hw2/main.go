package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, images int
	visit(doc, &words, &images)
	return words, images
}

func visit(n *html.Node, words, pics *int) {
	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}

var raw = `
<!DOCTYPE HTML>
<html>
    <body>
    <h1>My First Heading</h1>
    <p>My First paragraph</p>
    <p>HTML images are defined with the img tag:</p>
    <img src=xxx.jpg" width="104" height="142">
    </body>
</html> `

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed:%s\n", err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)
	fmt.Printf("%d words and %d images\n", words, pics)
}
