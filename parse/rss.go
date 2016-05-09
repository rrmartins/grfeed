package parse

import "github.com/mmcdole/gofeed"

// struct of Feed
type Feed struct {
	URL         string
	Title       string
	Description string
	Link        string
	Categories  []string
	Items       []*gofeed.Item
}

func New(url string) *Feed {
	feed := new(Feed)
	feed.URL = url
	return feed
}

func (f *Feed) RunParse() {
	fp := gofeed.NewParser()
	feedparse, _ := fp.ParseURL(f.URL)
	f.Title = feedparse.Title
	f.Description = feedparse.Description
	f.Link = feedparse.Link
	f.Categories = feedparse.Categories

	f.Items = feedparse.Items
}

func (f *Feed) GetTitle() string {
	return f.Title
}

func (f *Feed) GetDescription() string {
	return f.Description
}

func (f *Feed) GetLink() string {
	return f.Link
}

func (f *Feed) GetCategories() []string {
	var str []string
	if len(f.Categories) != 0 {
		str = f.Categories
	}
	str = []string{"No have Categories"}
	return str
}

func (f *Feed) GetTotalItems() int {
	return len(f.Items)
}
