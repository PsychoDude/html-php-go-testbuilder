package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// Open 'input.html'
	file, err := os.Open("input.html")
	if err != nil {
		log.Fatalf("Failed to open input.html: %v", err)
	}
	defer file.Close()

	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read input.html: %v", err)
	}

	// Parse the HTML content
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	// Process the HTML nodes
	processNodes(doc)

	// Render the modified HTML back to a string
	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		log.Fatalf("Failed to render HTML: %v", err)
	}

	// After rendering the HTML to buf
	renderedHTML := buf.String()

	// After unescaping PHP tags
	unescapedHTML := strings.ReplaceAll(renderedHTML, "&lt;?php", "<?php")
	unescapedHTML = strings.ReplaceAll(unescapedHTML, "?&gt;", "?>")

	// Replace HTML entities with corresponding characters
	unescapedHTML = strings.ReplaceAll(unescapedHTML, "&#39;", "'")

	// Write the corrected content to 'output.php'
	if err := os.WriteFile("output.php", []byte(unescapedHTML), 0644); err != nil {
		log.Fatalf("Failed to write output.php: %v", err)
	}
}

// processNodes recursively processes each node in the HTML document
func processNodes(n *html.Node) {

	if n.Type == html.ElementNode {
		// Process cms-title attribute
		if hasAttribute(n, "cms-title") {
			if n.FirstChild != nil {
				n.FirstChild.Data = "<?php echo $title; ?>"
			}
		}

		// Process title element directly for placeholder replacement
		if n.Data == "title" && n.FirstChild != nil {
			n.FirstChild.Data = strings.ReplaceAll(n.FirstChild.Data, "{{ TITLE }}", "<?php echo $title; ?>")
		}

		// Process cms-loop
		if hasAttribute(n, "cms-loop") {
			processCmsLoop(n)
		}
	}

	// Recurse for each child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processNodes(c)
	}
}

// processCmsLoop processes nodes with cms-loop attribute
func processCmsLoop(loopNode *html.Node) {
	// Find cms-card elements
	var cards []*html.Node
	for c := loopNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && hasAttribute(c, "cms-card") {
			cards = append(cards, c)
		}
	}

	// Keep only the first cms-card and remove others
	for i, card := range cards {
		if i == 0 {
			// Replace cms-card-title and cms-card-content
			replaceCmsCardContent(card)
		} else {
			// Remove the card from the loop
			loopNode.RemoveChild(card)
		}
	}

	// Wrap the remaining card with PHP loop
	wrapWithPhpLoop(loopNode)
}

// replaceCmsCardContent replaces cms-card-title and cms-card-content with PHP echo statements
func replaceCmsCardContent(cardNode *html.Node) {
	for c := cardNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			if hasAttribute(c, "cms-card-title") {
				c.FirstChild.Data = "<?php echo $card['title']; ?>"
			} else if hasAttribute(c, "cms-card-content") {
				c.FirstChild.Data = "<?php echo $card['content']; ?>"
			}
		}
	}
}

// wrapWithPhpLoop wraps the given node with a PHP foreach loop
func wrapWithPhpLoop(node *html.Node) {
	phpStart := &html.Node{
		Type: html.TextNode,
		Data: "<?php foreach ($cards as $card): ?>",
	}
	phpEnd := &html.Node{
		Type: html.TextNode,
		Data: "<?php endforeach; ?>",
	}

	node.InsertBefore(phpStart, node.FirstChild)
	node.AppendChild(phpEnd)
}

// hasAttribute checks if a node has a specific attribute
func hasAttribute(n *html.Node, attrName string) bool {
	for _, a := range n.Attr {
		if a.Key == attrName {
			return true
		}
	}
	return false
}
