package parse

import (
	rss "imartins/parserss/parse"
	"testing"
)

var url = "http://golangweekly.com/rss/1g2bo910"
var parserss = rss.New(url)

func TestGetTitleOK(t *testing.T) {
	parserss.RunParse()
	title := parserss.GetTitle()
	t.Log("Called parse class and GetTitle() method and return is Go Newsletter")
	if title != "Go Newsletter" {
		t.Errorf("Expected title of Go Newsletter, but it was %s instead.", title)
	}
}

func TestGetDescriptionOK(t *testing.T) {
	parserss.RunParse()
	description := parserss.GetDescription()
	t.Log("Called parse class and GetDescription() method and return is A weekly newsletter about the Go programming language")
	if description != "A weekly newsletter about the Go programming language" {
		t.Errorf("Expected description of A weekly newsletter about the Go programming language, but it was %s instead.", description)
	}
}

func TestGetLinkOK(t *testing.T) {
	parserss.RunParse()
	link := parserss.GetLink()
	t.Log("Called parse class and GetLink() method and return is http://golangweekly.com/")
	if link != "http://golangweekly.com/" {
		t.Errorf("Expected link of http://golangweekly.com/, but it was %s instead.", link)
	}
}

func TestGetCategoriesOK(t *testing.T) {
	parserss.RunParse()
	categories := parserss.GetCategories()
	t.Log("Called parse class and GetCategories() method and return is len 0")
	if len(categories) != 0 {
		t.Errorf("Expected categories of 0, but it was %s instead.", categories)
	}
}

func TestGetTotalItemsOK(t *testing.T) {
	parserss.RunParse()
	total_items := parserss.GetTotalItems()
	t.Log("Called parse class and GetTotalItems() method and return is len 4")
	if total_items != 4 {
		t.Errorf("Expected total_items of 4, but it was %d instead.", total_items)
	}
}
