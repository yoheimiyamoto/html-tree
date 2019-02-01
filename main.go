package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		err := prints(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "prints: %v\n", err)
			continue
		}
	}
}

func prints(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var count int

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if s := print(n); s != "" {
				fmt.Println(s)
			}
		}
		count++
	}

	post := func(n *html.Node) {
		count--
	}

	visit(doc, pre, post)
	return nil
}

func visit(n *html.Node, pre func(n *html.Node), post func(n *html.Node)) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if pre != nil {
			pre(c)
		}
		visit(c, pre, post)
		if post != nil {
			post(c)
		}
	}
}

func print(n *html.Node) string {
	var s string

	// in case of root directory
	if n.Data == "html" {
		return "."
	}

	// element name
	s = n.Data

	// class name
	for _, a := range n.Attr {
		if a.Key == "class" {
			if a.Val != "" {
				s = fmt.Sprintf("%s (%s)", s, a.Val)
			}
		}
	}

	// node
	if hasNextSibling(n) {
		s = fmt.Sprintf("├── %s", s)
	} else {
		s = fmt.Sprintf("└── %s", s)
	}

	// node from beginning to parent element
	for c := n.Parent; c.Parent != nil; c = c.Parent {
		if c.Data == "html" {
			return s
		}
		if hasNextSibling(c) {
			s = fmt.Sprintf("│    %s", s)
		} else {
			s = fmt.Sprintf("     %s", s)
		}
	}

	return s
}

// checking that element node has next sibling element node
func hasNextSibling(n *html.Node) bool {
	for c := n.NextSibling; c != nil; c = c.NextSibling {
		if c != nil {
			if c.Type == html.ElementNode {
				return true
			}
		}
	}
	return false
}
